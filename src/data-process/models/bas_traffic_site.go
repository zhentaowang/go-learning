package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type BasTrafficSite struct {
	Id          int       `orm:"column(id);pk"`
	Name        string    `orm:"column(name);size(100)"`
	AsName      string    `orm:"column(as_name);size(30);null"`
	Code        string    `orm:"column(code);size(6);null"`
	Continemt   string    `orm:"column(continemt);size(20)"`
	ContinemtId string    `orm:"column(continemt_id);size(32)"`
	Country     string    `orm:"column(country);size(50)"`
	CountryId   string    `orm:"column(country_id);size(32)"`
	Province    string    `orm:"column(province);size(50);null"`
	ProvinceId  string    `orm:"column(province_id);size(32);null"`
	City        string    `orm:"column(city);size(50)"`
	CityId      string    `orm:"column(city_id);size(32)"`
	Address     string    `orm:"column(address);size(200);null"`
	X           string    `orm:"column(x);size(15)"`
	Y           string    `orm:"column(y);size(15)"`
	Type        int       `orm:"column(type)"`
	Del         int8      `orm:"column(del)"`
	Around      int       `orm:"column(around)"`
	Sort        string    `orm:"column(sort);size(4)"`
	IataCode    string    `orm:"column(iata_code);size(4)"`
	Hot         int       `orm:"column(hot);null"`
	PinYin      string    `orm:"column(pin_yin);size(100);null"`
	Web         string    `orm:"column(web);size(50);null"`
	ImgUrl      string    `orm:"column(img_url);size(150);null"`
	CreateBy    string    `orm:"column(create_by);size(20);null"`
	CreateTime  time.Time `orm:"column(create_time);type(datetime);null"`
	UpdateBy    string    `orm:"column(update_by);size(20);null"`
	UpdateTime  time.Time `orm:"column(update_time);type(datetime);null"`
	Remarks     string    `orm:"column(remarks);size(2000);null"`
}

func (t *BasTrafficSite) TableName() string {
	return "bas_traffic_site"
}

func init() {
	orm.RegisterModel(new(BasTrafficSite))
}

func GetAllBasTrafficSiteNew() ([]BasTrafficSite, error) {

	o := orm.NewOrm()
	orm.Debug = true //true, print database operation log
	var basTrafficSite []BasTrafficSite

	//airline_name like concat('%',?,'%')
	num, err := o.Raw("select id, name from bas_traffic_site").QueryRows(&basTrafficSite)
	fmt.Println(num)
	return basTrafficSite, err

}

// AddBasTrafficSite insert a new BasTrafficSite into database and returns
// last inserted Id on success.
func AddBasTrafficSite(m *BasTrafficSite) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBasTrafficSiteById retrieves BasTrafficSite by Id. Returns error if
// Id doesn't exist
func GetBasTrafficSiteById(id int) (v *BasTrafficSite, err error) {
	o := orm.NewOrm()
	v = &BasTrafficSite{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBasTrafficSite retrieves all BasTrafficSite matches certain condition. Returns empty list if
// no records exist
func GetAllBasTrafficSite(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BasTrafficSite))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []BasTrafficSite
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateBasTrafficSite updates BasTrafficSite by Id and returns error if
// the record to be updated doesn't exist
func UpdateBasTrafficSiteById(m *BasTrafficSite) (err error) {
	o := orm.NewOrm()
	v := BasTrafficSite{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBasTrafficSite deletes BasTrafficSite by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBasTrafficSite(id int) (err error) {
	o := orm.NewOrm()
	v := BasTrafficSite{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&BasTrafficSite{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

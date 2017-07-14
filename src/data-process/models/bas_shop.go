package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type BasShop struct {
	Id                    string       `orm:"column(shop_info_id);pk"`
	ShopCode              string    `orm:"column(shop_code);size(10);null"`
	ShopName              string    `orm:"column(shop_name);size(100);null"`
	ShopType              string    `orm:"column(shop_type);size(20);null"`
	ShopStyle             string    `orm:"column(shop_style);size(20);null"`
	ShopState             string    `orm:"column(shop_state);size(20);null"`
	AirportRailwayId      int64     `orm:"column(airport_railway_id);null"`
	ThreeCode             string    `json:"threeCode" orm:"column(three_code);size(10);null"`
	AirportName           string    `orm:"column(airport_name);size(30);null"`
	TerminalType          string    `orm:"column(terminal_type);size(50);null"`
	TerminalId            int64     `orm:"column(terminal_id);null"`
	TerminalName          string    `orm:"column(terminal_name);size(30);null"`
	Region                string    `orm:"column(region);size(20);null"`
	SecurityBeforeOrAfter string    `orm:"column(security_before_or_after);size(30);null"`
	LocationGuidance      string    `orm:"column(location_guidance);size(120);null"`
	BusinessHours         string    `orm:"column(business_hours);size(30);null"`
	OpeningTime           string    `orm:"column(opening_time);size(20);null"`
	ClosingTime           string    `orm:"column(closing_time);size(20);null"`
	RelevantRule          string    `orm:"column(relevant_rule);size(120);null"`
	BoardGate             string    `orm:"column(board_gate);size(10);null"`
	NearestGate           string    `orm:"column(nearest_gate);size(20);null"`
	Email                 string    `orm:"column(email);size(30);null"`
	Telephone             string    `orm:"column(telephone);size(30);null"`
	Mode                  string    `orm:"column(mode);size(10);null"`
	Kinds                 string    `orm:"column(kinds);size(50);null"`
	ShopDescription       string    `orm:"column(shop_description);size(300);null"`
	Status                int       `orm:"column(status);null"`
	Sort                  int       `orm:"column(sort);null"`
	IsDeleted             int8      `orm:"column(is_deleted);null"`
	ImgUrl                string    `orm:"column(img_url);size(200);null"`
	CreateUser            int64     `orm:"column(create_user);null"`
	CreateTime            time.Time `orm:"column(create_time);type(datetime);null"`
	UpdateUser            int64     `orm:"column(update_user);null"`
	UpdateTime            time.Time `orm:"column(update_time);type(datetime);null"`
	Remarks               string    `orm:"column(remarks);size(2000);null"`
	ClickNum              int       `orm:"column(click_num);null"`
}

func (t *BasShop) TableName() string {
	return "bas_shop"
}

func init() {
	orm.RegisterModel(new(BasShop))
}

func GetBasShopNew() ([]BasShop, error) {

	o := orm.NewOrm()
	orm.Debug = true //true, print database operation log
	var basShopList []BasShop

	//airline_name like concat('%',?,'%')
	num, err := o.Raw("select shop_info_id, three_code from bas_shop").QueryRows(&basShopList)
	fmt.Println(num)
	return basShopList, err

}

func UpdateBasShopByIdNew(m []TerminalInfo) (err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BasShop))
	for _, terminal := range m {
		if _, err := qs.Filter("airport_railway_id", terminal.AirportId).Filter("board_gate", terminal.BoardGate).Update(orm.Params{"terminal_id": terminal.Id}); err != nil {
			return err
		}
	}
	return
}

func GetAllBasShopNew(queryParams map[string]interface{}) (basShop []BasShop, err error) {

	terminalList, _ := GetAllTerminalInfoNew()
	UpdateBasShopByIdNew(terminalList)
	return

}

// UpdateBasShop updates BasShop by Id and returns error if
// the record to be updated doesn't exist
func UpdateBasShopById(m *BasShop) (err error) {

	orm.NewOrm().QueryTable(new(BasShop)).Filter("shop_info_id", m.Id).Update(orm.Params{"business_hours": m.BusinessHours, "opening_time": m.OpeningTime, "closing_time": m.ClosingTime})
	return

}

// AddBasShop insert a new BasShop into database and returns
// last inserted Id on success.
func AddBasShop(m *BasShop) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBasShopById retrieves BasShop by Id. Returns error if
// Id doesn't exist
func GetBasShopById(id int) (v *BasShop, err error) {
	//o := orm.NewOrm()
	//v = &BasShop{Id: id}
	//if err = o.Read(v); err == nil {
	//	return v, nil
	//}
	return nil, err
}

// GetAllBasShop retrieves all BasShop matches certain condition. Returns empty list if
// no records exist
func GetAllBasShop(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BasShop))
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

	var l []BasShop
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

// DeleteBasShop deletes BasShop by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBasShop(id int) (err error) {
	//o := orm.NewOrm()
	//v := BasShop{Id: id}
	//// ascertain id exists in the database
	//if err = o.Read(&v); err == nil {
	//	var num int64
	//	if num, err = o.Delete(&BasShop{Id: id}); err == nil {
	//		fmt.Println("Number of records deleted in database:", num)
	//	}
	//}
	return
}

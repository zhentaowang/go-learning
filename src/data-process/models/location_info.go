package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"strconv"
	"airport-cloud/base-info/utils"
)

type LocationInfo struct {
	Id                 int       `json:"id" orm:"column(location_info_id);auto"`
	AreaName           string    `json:"areaName" orm:"column(area_name);size(20);null"`
	AreaAlias          string    `json:"areaAlias" orm:"column(area_alias);size(20);null"`
	Pinyin             string    `json:"pinyin" orm:"column(pinyin);size(20);null"`
	Level              string    `json:"level" orm:"column(level);size(20);null"`
	Type               int16     `json:"type" orm:"column(type);null"`
	UpRegion           string    `json:"upRegion" orm:"column(up_region);size(20);null"`
	Sort               int       `json:"sort" orm:"column(sort);null"`
	TelephonePrefix    string    `json:"telephonePrefix" orm:"column(telephone_prefix);size(20);null"`
	LicensePlatePrefix string    `json:"licensePlatePrefix" orm:"column(license_plate_prefix);size(20);null"`
	TimeZone           string    `json:"timeZone" orm:"column(time_zone);size(20);null"`
	Longitude          string    `json:"longitude" orm:"column(longitude);size(20);null"`
	Latitude           string    `json:"latitude" orm:"column(latitude);size(20);null"`
	Ios                string    `json:"ios" orm:"column(ios);size(20);null"`
	Ios3               string    `json:"ios3" orm:"column(ios3);size(30);null"`
	DigitalCode        string    `json:"digitalCode" orm:"column(digital_code);size(20);null"`
	Remark             string    `json:"remark" orm:"column(remark);size(200);null"`
	IsDeleted          int8      `json:"isDeleted,omitempty" orm:"column(is_deleted);null"`
	CreateUser         int64     `json:"createUser,omitempty" orm:"column(create_user);null"`
	UpdateUser         int64     `json:"updateUser,omitempty" orm:"column(update_user);null"`
	CreateTime         time.Time `json:"-" orm:"column(create_time);type(datetime);null;auto_now_add"`
	UpdateTime         time.Time `json:"-" orm:"column(update_time);type(datetime);null"`
}

func (t *LocationInfo) TableName() string {
	return "location_info"
}

func init() {
	orm.RegisterModel(new(LocationInfo))
}

// AddLocationInfo insert a new LocationInfo into database and returns
// last inserted Id on success.
func AddLocationInfo(m *LocationInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return id, nil
}

// UpdateLocationInfo updates LocationInfo by Id and returns error if
// the record to be updated doesn't exist
// UpdateAirlineInfo updates AirlineInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateLocationInfoById(m *LocationInfo) (err error) {

	o := orm.NewOrm()
	v := LocationInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return err

}

// GetLocationInfoById retrieves LocationInfo by Id. Returns error if
// Id doesn't exist
func GetLocationInfoById(id int) (v *LocationInfo, err error) {
	o := orm.NewOrm()
	v = &LocationInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateLocationInfo updates LocationInfo by Id and returns error if
// the record to be updated doesn't exist
func DeleteLocationInfoInBatch(m []interface{}) (err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LocationInfo))
	for _,id := range m {
		if _, err := qs.Filter("location_info_id", id).Update(orm.Params{"is_deleted": 1}); err != nil {
			return err
		}
	}
	return err
}

// GetAllLocationInfoNew retrieves all LocationInfo matches certain condition. Returns empty list if
// no records exist
func GetAllLocationInfoNew(queryParams map[string]interface{}) ([]LocationInfo, int64, error) {

	o := orm.NewOrm()
	orm.Debug = true //true, print database operation log
	qs := o.QueryTable(new(LocationInfo))

	var sqlCondition = "is_deleted = 0"
	for k, v := range queryParams {
		// assemble sql
		if k != "page" && k != "rows" && k != "operation" && v != "" && v != nil {
			k = utils.FormatSwitch(k) // switch format from hump to snake
			if k == "area_name" { // fuzzy query
				sqlCondition += " and area_name like " +  "'%" + v.(string) + "%'" + " or pinyin like " +  "'%" + v.(string) + "%'"
			} else { // =
				sqlCondition += " and " + k + " = " + v.(string)
			}
		}
	}

	var locationInfo []LocationInfo
	fmt.Println(sqlCondition)
	page, _ := strconv.Atoi(queryParams["page"].(string))
	rows, _ := strconv.Atoi(queryParams["rows"].(string))
	//airline_name like concat('%',?,'%')
	total, _ := qs.Filter("is_deleted", 0).Filter("area_name__contains", queryParams["areaName"]).Count()
	if total == 0 {
		total, _ = qs.Filter("is_deleted", 0).Filter("pinyin__contains", queryParams["areaName"]).Count()
	}
	num, err := o.Raw("select * from location_info where " + sqlCondition + " order by location_info_id desc limit ?,?", (page-1)*10, rows).QueryRows(&locationInfo)
	fmt.Println(num)
	return locationInfo, total, err

}

// GetAllLocationInfo retrieves all LocationInfo matches certain condition. Returns empty list if
// no records exist
func GetAllLocationInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LocationInfo))
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

	var l []LocationInfo
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

// DeleteLocationInfo deletes LocationInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLocationInfo(id int) (err error) {
	o := orm.NewOrm()
	v := LocationInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LocationInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

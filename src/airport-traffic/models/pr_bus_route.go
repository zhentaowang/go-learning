package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type PrBusRoute struct {
	Id               int       `orm:"column(route_id);auto"`
	RouteName        string    `orm:"column(route_name);size(20);null"`
	RouteType        int8      `orm:"column(route_type);null"`
	ParentId         int64     `orm:"column(parent_id);null"`
	TrafficId        int64     `orm:"column(traffic_id);null"`
	Hotline          string    `orm:"column(hotline);size(255);null"`
	FirstBus         string    `orm:"column(first_bus);size(255);null"`
	LastBus          string    `orm:"column(last_bus);size(255);null"`
	RouteDescription string    `orm:"column(route_description);size(200);null"`
	IsDeleted        int8      `orm:"column(is_deleted);null"`
	CreateUser       int64     `orm:"column(create_user);null"`
	UpdateUser       int64     `orm:"column(update_user);null"`
	CreateTime       time.Time `orm:"column(create_time);type(datetime);null;auto_now_add"`
	UpdateTime       time.Time `orm:"column(update_time);type(datetime);null"`
}

func (t *PrBusRoute) TableName() string {
	return "pr_bus_route"
}

func init() {
	orm.RegisterModel(new(PrBusRoute))
}

// AddPrBusRoute insert a new PrBusRoute into database and returns
// last inserted Id on success.
func AddPrBusRoute(m *PrBusRoute) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPrBusRouteById retrieves PrBusRoute by Id. Returns error if
// Id doesn't exist
func GetPrBusRouteById(id int) (v *PrBusRoute, err error) {
	o := orm.NewOrm()
	v = &PrBusRoute{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPrBusRoute retrieves all PrBusRoute matches certain condition. Returns empty list if
// no records exist
func GetAllPrBusRoute(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PrBusRoute))
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

	var l []PrBusRoute
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

// UpdatePrBusRoute updates PrBusRoute by Id and returns error if
// the record to be updated doesn't exist
func UpdatePrBusRouteById(m *PrBusRoute) (err error) {
	o := orm.NewOrm()
	v := PrBusRoute{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePrBusRoute deletes PrBusRoute by Id and returns error if
// the record to be deleted doesn't exist
func DeletePrBusRoute(id int) (err error) {
	o := orm.NewOrm()
	v := PrBusRoute{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PrBusRoute{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

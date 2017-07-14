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

type ShopInfo struct {
	Id                    int       `json:"id" orm:"column(shop_info_id);auto"`
	TerminalId            int64     `json:"terminalId" orm:"column(terminal_id);null"`
	AirportName           string	`json:"airportName" orm:"column(airport_name);size(30);null"`
	TerminalName          string	`json:"terminalName" orm:"column(terminal_name);size(20);null"`
	TerminalType          string	`json:"terminalType" orm:"column(terminal_type);size(20);null"`
	ShopCode              string    `json:"shopCode" orm:"column(shop_code);size(20);null"`
	ShopName              string    `json:"shopName" orm:"column(shop_name);size(20);null"`
	OpeningTime           string    `json:"openingTime" orm:"column(opening_time);size(20);null"`
	ClosingTime           string    `json:"closingTime" orm:"column(closing_time);size(20);null"`
	Telephone             string    `json:"telephone" orm:"column(telephone);size(20);null"`
	ShopType              string    `json:"shopType" orm:"column(shop_type);size(20);null"`
	ShopStyle             string    `json:"shopStyle" orm:"column(shop_style);size(20);null"`
	ShopDescription       string    `json:"shopDescription" orm:"column(shop_description);size(200);null"`
	ShopState             string    `json:"shopState" orm:"column(shop_state);size(20);null"`
	ImgUrl                string    `json:"imgUrl" orm:"column(img_url);size(100);null"`
	BoardGate             string    `json:"boardGate" orm:"column(board_gate);size(20);null"`
	SecurityBeforeOrAfter string    `json:"securityBeforeOrAfter" orm:"column(security_before_or_after);size(20);null"`
	LocationGuidance      string    `json:"locationGuidance" orm:"column(location_guidance);size(100);null"`
	RelevantRule          string    `json:"relevantRule" orm:"column(relevant_rule);size(200);null"`
	IsDeleted             int8      `json:"isDeleted,omitempty" orm:"column(is_deleted);null"`
	CreateUser            int64     `json:"createUser,omitempty" orm:"column(create_user);null"`
	UpdateUser            int64     `json:"updateUser,omitempty" orm:"column(update_user);null"`
	CreateTime            time.Time `json:"-" orm:"column(create_time);type(datetime);null;auto_now_add"`
	UpdateTime            time.Time `json:"-" orm:"column(update_time);type(datetime);null"`
}

func (t *ShopInfo) TableName() string {
	return "shop_info"
}

func init() {
	orm.RegisterModel(new(ShopInfo))
}

// AddShopInfo insert a new ShopInfo into database and returns
// last inserted Id on success.
func AddShopInfo(m *ShopInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// UpdateShopInfo updates ShopInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateShopInfoById(m *ShopInfo) (err error) {
	o := orm.NewOrm()
	v := ShopInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// GetShopInfoById retrieves ShopInfo by Id. Returns error if
// Id doesn't exist
func GetShopInfoById(id int) (v *ShopInfo, err error) {
	o := orm.NewOrm()
	v = &ShopInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllShopInfoNew retrieves all ShopInfo matches certain condition. Returns empty list if
// no records exist
func GetAllShopInfoNew(queryParams map[string]interface{}) ([]ShopInfo, int64, error) {

	o := orm.NewOrm()
	orm.Debug = true //true, print database operation log
	qs := o.QueryTable(new(ShopInfo))

	var sqlCondition = "is_deleted = 0"
	for k, v := range queryParams {
		// assemble sql
		if k != "page" && k != "rows" && k != "operation" && v != "" && v != nil {
			k = utils.FormatSwitch(k) // switch format from hump to snake
			if k == "shop_name" { // fuzzy query
				sqlCondition += " and shop_name like " +  "'%" + v.(string) + "%'"
			} else { // =
				sqlCondition += " and " + k + " = " + v.(string)
			}
		}
	}
	var shopInfo []ShopInfo
	fmt.Println(sqlCondition)
	page, _ := strconv.Atoi(queryParams["page"].(string))
	rows, _ := strconv.Atoi(queryParams["rows"].(string))
	//airline_name like concat('%',?,'%')
	total, err := qs.Filter("is_deleted", 0).Filter("shop_name__contains", queryParams["shopName"]).Count()
	num, err := o.Raw("select * from shop_info where " + sqlCondition + " order by shop_info_id desc limit ?,?", (page-1)*10, rows).QueryRows(&shopInfo)
	fmt.Println(num)
	return shopInfo, total, err

}

// GetAllShopInfo retrieves all ShopInfo matches certain condition. Returns empty list if
// no records exist
func GetAllShopInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ShopInfo))
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

	var l []ShopInfo
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

// DeleteShopInfoInBatch updates ShopInfo by Id and returns error if
// the record to be updated doesn't exist
func DeleteShopInfoInBatch(m []interface{}) (err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ShopInfo))
	for _,id := range m {
		if _, err := qs.Filter("shop_info_id", id).Update(orm.Params{"is_deleted": 1}); err != nil {
			return err
		}
	}
	return err
}

// DeleteShopInfo deletes ShopInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteShopInfo(id int) (err error) {
	o := orm.NewOrm()
	v := ShopInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ShopInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

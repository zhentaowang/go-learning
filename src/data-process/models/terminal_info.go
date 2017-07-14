package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"time"
	"strconv"
)

type TerminalInfo struct {
	Id                   int       `json:"terminalId" orm:"column(terminal_info_id)";pk`
	AirportId            int64     `json:"airportId" orm:"column(airport_id);null"`
	ThreeCode            string    `json:"threeCode" orm:"column(three_code);size(10);null"`
	AirportName          string    `json:"airportName" orm:"column(airport_name);size(32);null"`
	TerminalName         string    `json:"terminalName" orm:"column(terminal_name);size(32);null"`
	BoardGate            string    `json:"boardGate" orm:"column(board_gate);size(10);null"`
	TerminalType         string    `json:"terminalType" orm:"column(terminal_type);size(50);null"`
	Remark               string    `json:"remark" orm:"column(remark);size(200);null"`
	IsDeleted             int8      `json:"isDeleted,omitempty" orm:"column(is_deleted);null"`
	CreateUser            int64     `json:"createUser,omitempty" orm:"column(create_user);null"`
	UpdateUser            int64     `json:"updateUser,omitempty" orm:"column(update_user);null"`
	CreateTime            time.Time `json:"-" orm:"column(create_time);type(datetime);null;auto_now_add"`
	UpdateTime            time.Time `json:"-" orm:"column(update_time);type(datetime);null"`
	//AirportRailwayInfo        *AirportRailwayInfo `json:"airportRailwayInfo,omitempty" orm:"null;rel(fk)"` // set one to many relation
}

func (t *TerminalInfo) TableName() string {
	return "terminal_info"
}

func init() {
	orm.RegisterModel(new(TerminalInfo))
}

// DeleteTerminalInfo updates TerminalInfo by Id and returns error if
// the record to be updated doesn't exist
func DeleteTerminalInfoInBatch(id interface{}) (err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TerminalInfo))
	_, err = qs.Filter("airport_id", id).Update(orm.Params{"is_deleted": 1})
	return err
}

// Delete TerminalInfo by Ids and returns error if
// the record to be updated doesn't exist
func DeleteBest(idMap []int, airportId int) (err error) {

	o := orm.NewOrm()
	orm.Debug = true
	qs := o.QueryTable(new(TerminalInfo))

	if len(idMap) == 0 {
		_, err = qs.Filter("airport_id", airportId).Update(orm.Params{"is_deleted": 1})
	}else {
		_, err = qs.Exclude("terminal_info_id__in", idMap).Filter("airport_id", airportId).Update(orm.Params{"is_deleted": 1})
	}
	if err != nil {
		fmt.Println("删除terminal_info信息出错")
	}
	return err

}

// Delete TerminalInfo by Ids and returns error if
// the record to be updated doesn't exist
func Delete(ids string, airportId int) (err error) {

	o := orm.NewOrm()
	orm.Debug = true
	var sqlCondition string
	idMap := []int{}

	if ids != "" {
		sqlCondition = "and terminal_info_id not in ("
		idArray := strings.Split(ids, ",")
		for _, id := range idArray {
			tem, _:= strconv.Atoi(id)
			idMap = append(idMap, tem)
			sqlCondition += "?,"
		}
		sqlCondition = sqlCondition[0:len(sqlCondition)-1]
		sqlCondition += ")"
	}

	_, err = o.Raw("update terminal_info set is_deleted = 1 where airport_id = ? " + sqlCondition, airportId, idMap).Exec()
	if err != nil {
		fmt.Println("删除terminal_info信息出错")
	}
	return err

}

// AddTerminalInfo insert a new TerminalInfo into database and returns
// last inserted Id on success.
func AddTerminalInfo(m *TerminalInfo) (id int64, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	id, err = o.Insert(m)

	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}

	return
}

// GetTerminalInfoById retrieves TerminalInfo by Id. Returns error if
// Id doesn't exist
func GetTerminalInfoById(id int) (v *TerminalInfo, err error) {
	o := orm.NewOrm()
	v = &TerminalInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTerminalInfoNew retrieves TerminalInfo by airportId. Returns error if
// airportId doesn't exist
func GetAllTerminalInfoNew() (terminalList []TerminalInfo, err error) {
	o := orm.NewOrm()
	total, err := o.Raw("select terminal_info_id, airport_id, board_gate from terminal_info").QueryRows(&terminalList)
	fmt.Println(total)
	return terminalList, err
}

// GetAllTerminalInfo retrieves all TerminalInfo matches certain condition. Returns empty list if
// no records exist
func GetAllTerminalInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TerminalInfo))
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

	var l []TerminalInfo
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

// UpdateTerminalInfo updates TerminalInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateTerminalInfoById(m *TerminalInfo) (err error) {
	o := orm.NewOrm()
	v := TerminalInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTerminalInfo deletes TerminalInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTerminalInfo(id int) (err error) {
	o := orm.NewOrm()
	v := TerminalInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TerminalInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

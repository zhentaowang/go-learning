package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"encoding/json"
	"strconv"
)

type AirportRailwayInfo struct {
	Id              int       `json:"id" orm:"column(airport_railway_info_id);pk"`
	AirportId       string    `json:"airportId" orm:"column(airport_id)"`
	Name            string    `json:"name" orm:"column(name);size(32);null"`
	Type            string    `json:"type" orm:"column(type);size(50);null"`
	ThreeCode       string    `json:"threeCode" orm:"column(three_code);size(10);null"`
	Pinyin          string    `json:"pinyin" orm:"column(pinyin);size(32);null"`
	ShortTitle      string    `json:"shortTitle" orm:"column(short_title);size(20);null"`
	HotNum          int64     `json:"hotNum,omitempty" orm:"column(hot_num);null"`
	City            string    `json:"city" orm:"column(city);size(20);null"`
	Region          int8      `json:"region" orm:"column(region);null"`
	Position        string    `json:"position" orm:"column(position);size(30);null"`
	OfficialWebsite string    `json:"officialWebsite" orm:"column(official_website);size(100);null"`
	ImgUrl          string    `json:"imgUrl" orm:"column(img_url);size(200);null"`
	Remark          string    `json:"remark" orm:"column(remark);size(200);null"`
	IsDeleted       int8      `json:"isDeleted,omitempty" orm:"column(is_deleted);null"`
	CreateUser      int64     `json:"createUser,omitempty" orm:"column(create_user);null"`
	UpdateUser      int64     `json:"updateUser,omitempty" orm:"column(update_user);null"`
	CreateTime      time.Time `json:"-" orm:"column(create_time);type(datetime);null;auto_now_add"`
	UpdateTime      time.Time `json:"-" orm:"column(update_time);type(datetime);null"`
	//Terminal        []*TerminalInfo `json:"terminal" orm:"reverse(many)"` // set one to many reverse relation
}

func (t *AirportRailwayInfo) TableName() string {
	return "airport_railway_info"
}

func init() {
	orm.RegisterModel(new(AirportRailwayInfo))
}

// AddAirportRailwayInfoNew insert a new AirportRailwayInfo into database and returns
// last inserted Id on success.
func AddAirportRailwayInfoNew(m *AirportRailwayInfo, terminal []interface{}) (id int64, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	id, err = o.Insert(m)

	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}

	if m.Type == "机场" {
		for  _,t := range terminal {
			var tt TerminalInfo
			dd, _ := json.Marshal(t)
			json.Unmarshal(dd, &tt)
			tt.AirportId = id
			_, _ = AddTerminalInfo(&tt)
		}
	}
	return
}

// AddAirportRailwayInfo insert a new AirportRailwayInfo into database and returns
// last inserted Id on success.
func AddAirportRailwayInfo(m *AirportRailwayInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// UpdateAirportRailwayInfo updates AirportRailwayInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateAirportRailwayInfoById(m *AirportRailwayInfo) (err error) {
	o := orm.NewOrm()
	v := AirportRailwayInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// GetAirportRailwayInfoById retrieves AirportRailwayInfo by Id. Returns error if
// Id doesn't exist
func GetAirportRailwayInfoById(id int) (v map[string]interface{}, err error) {
	o := orm.NewOrm()
	var airportRailwayInfo AirportRailwayInfo
	airportRailwayInfoMap := make(map[string]interface{})
	err = o.Raw("select * from airport_railway_info where is_deleted = 0 and airport_railway_info_id = ?", id).QueryRow(&airportRailwayInfo)
	return airportRailwayInfoMap, err
}

// for flight-center
func GetHotAirportRailwayInfoByCondition(count int,region int) ([]AirportRailwayInfo, error) {
	o := orm.NewOrm()
	orm.Debug = true
	var airportRailwayInfos []AirportRailwayInfo
	num,err := o.Raw("SELECT * FROM airport_railway_info WHERE region = ? AND is_deleted = 0 ORDER BY hot_num DESC LIMIT ?", region,count).QueryRows(&airportRailwayInfos)
	fmt.Println(num)
	return airportRailwayInfos, err
}

func GetAllAirportRailwayInfoByCondition(queryName string,region int,trafficType string) ([]AirportRailwayInfo,error) {
	o := orm.NewOrm()
	orm.Debug = true
	var airportRailwayInfo []AirportRailwayInfo
	var err error
	var num int64
	
	if len(strings.TrimSpace(trafficType)) !=0 {
		if len(strings.TrimSpace(queryName)) != 0{
			num,err = o.Raw("SELECT * FROM( " +
				"SELECT * FROM airport_railway_info WHERE name like CONCAT('%',?,'%')" +
				"UNION SELECT * FROM airport_railway_info WHERE three_code like CONCAT('%',?,'%')" +
				"UNION SELECT * FROM airport_railway_info WHERE short_title like CONCAT('%',?,'%')" +
				") ari WHERE ari.type = ? ORDER BY ari.three_code DESC", queryName,queryName,queryName,trafficType).QueryRows(&airportRailwayInfo)
		}else {
			num,err = o.Raw("SELECT * FROM airport_railway_info WHERE region = ? AND type = ? AND is_deleted = 0 ORDER BY three_code DESC", region,trafficType).QueryRows(&airportRailwayInfo)
		}
	}else {
		if len(strings.TrimSpace(queryName)) != 0{
			num,err = o.Raw("SELECT * FROM( " +
				"SELECT * FROM airport_railway_info WHERE name like CONCAT('%',?,'%')" +
				"UNION SELECT * FROM airport_railway_info WHERE three_code like CONCAT('%',?,'%')" +
				"UNION SELECT * FROM airport_railway_info WHERE short_title like CONCAT('%',?,'%')" +
				") ari ORDER BY ari.three_code DESC", queryName,queryName,queryName).QueryRows(&airportRailwayInfo)
		}else {
			num,err = o.Raw("SELECT * FROM airport_railway_info WHERE region = ? AND is_deleted = 0 ORDER BY three_code DESC", region).QueryRows(&airportRailwayInfo)
		}
	}
	
	fmt.Println(num)
	return airportRailwayInfo, err
}

func GetAirportRailwayInfoByThreeCode(airportCode string) (AirportRailwayInfo,error) {
	o := orm.NewOrm()
	var airportRailwayInfo AirportRailwayInfo
	err := o.Raw("SELECT * FROM airport_railway_info WHERE three_code = ? AND is_deleted =0 ", airportCode).QueryRow(&airportRailwayInfo)
	return airportRailwayInfo, err
}

func UpdateAirportRailwayInfoNew(m *AirportRailwayInfo, terminal []interface{}) (err error) {
	o := orm.NewOrm()
	orm.Debug = true
	v := AirportRailwayInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			if m.Type == "机场" {
				idMap := []int{}
				for  _,t := range terminal {
					var tt TerminalInfo
					dd, _ := json.Marshal(t)
					json.Unmarshal(dd, &tt)
					if tt.Id != 0 {
						UpdateTerminalInfoById(&tt)
						idMap = append(idMap, tt.Id)
					}else {
						id, _ :=AddTerminalInfo(&tt)
						id0, _ := strconv.Atoi(strconv.FormatInt(id, 10))
						idMap = append(idMap, id0)
					}
				}
				DeleteBest(idMap, m.Id)
			}
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func UpdateHotNumByThreeCode(airportCode string) (int int64, err error) {
	o := orm.NewOrm()
	res,err :=o.Raw("UPDATE airport_railway_info SET hot_num = hot_num +1 WHERE three_code = ? AND is_deleted =0",airportCode).Exec()
	if err == nil{
		num,_ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	return
}

// GetAllAirlineInfoNew retrieves all AirlineInfo matches certain condition. Returns empty list if
// no records exist
func GetAllAirportRailwayInfoNew(queryParams map[string]interface{}) ([]AirportRailwayInfo, error) {

	//basTrafficSites, err := GetAllBasTrafficSiteNew()
	//UpdateAirportRailwayInfoByIdNew(basTrafficSites)
	//
	//
	//return basTrafficSites, err

	o := orm.NewOrm()
	orm.Debug = true //true, print database operation log
	var airportRailwayInfoList []AirportRailwayInfo

	//airline_name like concat('%',?,'%')
	num, err := o.Raw("select airport_railway_info_id, three_code from airport_railway_info").QueryRows(&airportRailwayInfoList)
	fmt.Println(num)

	//UpdateBasShopByIdNew(airportRailwayInfoList)

	return airportRailwayInfoList, err

}

func UpdateAirportRailwayInfoByIdNew(m []BasTrafficSite) (err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AirportRailwayInfo))
	for _, basTrafficSite := range m {
		if _, err := qs.Filter("name", basTrafficSite.Name).Update(orm.Params{"airport_id": basTrafficSite.Id}); err != nil {
			return err
		}
	}
	return
}

// DeleteAirportRailwayInfo updates AirlineInfo by Id and returns error if
// the record to be updated doesn't exist
func DeleteAirportRailwayInfoInBatch(m []interface{}, category string) (err1 error, err2 error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AirportRailwayInfo))
	for _,id := range m {
		if _, err1 := qs.Filter("airport_railway_info_id", id).Update(orm.Params{"is_deleted": 1}); err1 == nil {
			if category == "机场" {
				err2 := DeleteTerminalInfoInBatch(id)
				return err1, err2
			}
		}
	}
	return err1, err2
}

// GetAllAirportRailwayInfo retrieves all AirportRailwayInfo matches certain condition. Returns empty list if
// no records exist
func GetAllAirportRailwayInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AirportRailwayInfo))
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

	var l []AirportRailwayInfo
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

// DeleteAirportRailwayInfo deletes AirportRailwayInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAirportRailwayInfo(id int) (err error) {
	o := orm.NewOrm()
	v := AirportRailwayInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AirportRailwayInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

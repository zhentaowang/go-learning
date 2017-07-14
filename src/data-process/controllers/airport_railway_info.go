package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"log"
	"code.aliyun.com/wyunshare/thrift-server/gen-go/server"
	"data-process/models"
)

// AirportRailwayInfoController operations for AirportRailwayInfo
type AirportRailwayInfoController struct {
	beego.Controller
}

// URLMapping ...
func (c *AirportRailwayInfoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *AirportRailwayInfoController) GetAirportRailwayInfoById(param map[string]interface{}) (r *server.Response, err error) {

	//id, _ := strconv.Atoi(param["id"].(string))
	//category := param["type"].(string)
	//v, err := models.GetAirportRailwayInfoByIdNew(id, category)
	//
	//r = server.NewResponse()
	//paramOutput := make(map[string]interface{})
	//if err != nil {
	//	paramOutput["msg"] = "操作失败"
	//	paramOutput["status"] = server.RESCODE__500.Int()
	//	paramOutput["data"] = v
	//	result := []byte{}
	//	result, err = json.Marshal(paramOutput)
	//	r.ResponseJSON = result
	//	r.ResponeCode = server.RESCODE__500
	//}else {
	//	paramOutput["msg"] = "操作成功"
	//	paramOutput["status"] = server.RESCODE__200.Int()
	//	paramOutput["data"] = v
	//	result := []byte{}
	//	result, err = json.Marshal(paramOutput)
	//	r.ResponseJSON = result
	//	r.ResponeCode = server.RESCODE__200
	//}
	return r, nil

}

func (c *AirportRailwayInfoController) GetAirportRailwayInfoAll(queryParams map[string]interface{}) (r *server.Response, err error) {

	l, err := models.GetAllAirportRailwayInfoNew(queryParams)
	data := make(map[string]interface{})
	data["rows"] = l
	r = server.NewResponse()
	paramOutput := make(map[string]interface{})
	if err != nil {
		paramOutput["msg"] = "操作失败"
		paramOutput["status"] = server.RESCODE__500.Int()
		result := []byte{}
		result, err = json.Marshal(paramOutput)
		r.ResponseJSON = result
		r.ResponeCode = server.RESCODE__500
	}else {
		paramOutput["msg"] = "操作成功"
		paramOutput["status"] = server.RESCODE__200.Int()
		paramOutput["data"] = data
		result := []byte{}
		result, err = json.Marshal(paramOutput)
		r.ResponseJSON = result
		r.ResponeCode = server.RESCODE__200
	}
	return r, nil

}

func (c *AirportRailwayInfoController) AddAirportRailwayInfo(queryParams []byte) (r *server.Response, err error) {

	var v models.AirportRailwayInfo
	var terminal []interface{}
	r = server.NewResponse()
	paramInput := make(map[string]interface{})
	paramOutput := make(map[string]interface{})
	json.Unmarshal(queryParams, &paramInput)

	if paramInput["type"] == "机场" {
		terminal = paramInput["terminal"].([]interface{})
	}

	json.Unmarshal(queryParams, &v)

	if _, err := models.AddAirportRailwayInfoNew(&v, terminal); err == nil {
		paramOutput["msg"] = "操作成功"
		paramOutput["status"] = server.RESCODE__200.Int()
		result := []byte{}
		result, err = json.Marshal(paramOutput)
		r.ResponseJSON = result
		r.ResponeCode = server.RESCODE__200
	} else {
		paramOutput["msg"] = "操作失败"
		paramOutput["status"] = server.RESCODE__500.Int()
		result := []byte{}
		result, err = json.Marshal(paramOutput)
		r.ResponseJSON = result
		r.ResponeCode = server.RESCODE__500
	}
	return r, err

}

func (c *AirportRailwayInfoController) UpdateAirportRailwayInfoById(queryParams []byte) (r *server.Response, err error) {

	var v models.AirportRailwayInfo
	var terminal []interface{}
	r = server.NewResponse()
	paramInput := make(map[string]interface{})
	paramOutput := make(map[string]interface{})
	json.Unmarshal(queryParams, &paramInput)

	if paramInput["type"] == "机场" {
		terminal = paramInput["terminal"].([]interface{})
	}

	json.Unmarshal(queryParams, &v)

	if err := models.UpdateAirportRailwayInfoNew(&v, terminal); err == nil {
		paramOutput["msg"] = "操作成功"
		paramOutput["status"] = server.RESCODE__200.Int()
		result := []byte{}
		result, err = json.Marshal(paramOutput)
		r.ResponseJSON = result
		r.ResponeCode = server.RESCODE__200
	} else {
		paramOutput["msg"] = "操作失败"
		paramOutput["status"] = server.RESCODE__500.Int()
		result := []byte{}
		result, err = json.Marshal(paramOutput)
		r.ResponseJSON = result
		r.ResponeCode = server.RESCODE__500
	}
	return r, err

}

func (c *AirportRailwayInfoController) DeleteAirportRailwayInfoInBatch(queryParams []byte) (r *server.Response, err1 error) {

	paramInput := make(map[string]interface{})
	json.Unmarshal(queryParams, &paramInput)
	ids := paramInput["data"].([]interface{})
	category := paramInput["type"].(string)
	log.Println(ids)

	var v models.AirportRailwayInfo
	r = server.NewResponse()
	paramOutput := make(map[string]interface{})
	json.Unmarshal(queryParams, &v)

	if err1, err2 := models.DeleteAirportRailwayInfoInBatch(ids, category); err1 == nil && err2 == nil {
		paramOutput["msg"] = "操作成功"
		paramOutput["status"] = server.RESCODE__200.Int()
		result := []byte{}
		result, _ = json.Marshal(paramOutput)
		r.ResponseJSON = result
		r.ResponeCode = server.RESCODE__200
	} else {
		paramOutput["msg"] = "操作失败"
		paramOutput["status"] = server.RESCODE__500.Int()
		result := []byte{}
		result, _ = json.Marshal(paramOutput)
		r.ResponseJSON = result
		r.ResponeCode = server.RESCODE__500
	}
	return r, err1

}

// for flight-center by ironc 2017-06-22
func (c *AirportRailwayInfoController) QueryHotAirport(params map[string]interface{})(r *server.Response, err error) {
	
	var count int
	var region int
	val_count := params["count"]
	val_region := params["region"]
	
	if val_count == nil{
		count = 6
	}else {
		count,err = strconv.Atoi(params["count"].(string))
		if count <0{
			count = 6
		}
	}
	if val_region == nil{
		region = 0
	}else {
		region,err = strconv.Atoi(val_region.(string))
 		if region !=0 && region !=1{
			region = 0
		}
	}
	
	r = server.NewResponse()
	v,err := models.GetHotAirportRailwayInfoByCondition(count,region)
	paramOutput := make(map[string]interface{})
	if err != nil {
		paramOutput["msg"] = "操作失败"
		paramOutput["status"] = server.RESCODE__500.Int()
		result := []byte{}
		result, err = json.Marshal(paramOutput)
		r.ResponseJSON = result
		r.ResponeCode = server.RESCODE__500
	}else {
		paramOutput["msg"] = "操作成功"
		paramOutput["status"] = server.RESCODE__200.Int()
		paramOutput["data"] = v
		result := []byte{}
		result, err = json.Marshal(paramOutput)
		r.ResponseJSON = result
		r.ResponeCode = server.RESCODE__200
	}
  	return r, nil
}

func (c *AirportRailwayInfoController) QueryAirports(params map[string]interface{}) (r *server.Response, err error) {
	
	var queryName string
	var region int
	var trafficType string
	val_queryName := params["queryName"]
	val_region := params["region"]
	val_type := params["type"]
	
	if val_queryName == nil{
		queryName = ""
	}else {
		queryName = val_queryName.(string)
	}
	if val_region == nil{
		region = 0
	}else {
		region,err = strconv.Atoi(val_region.(string))
		if region !=0 && region !=1{
			region = 0
		}
	}
	if val_type == nil{
		trafficType = ""
	}else {
		trafficType = val_type.(string)
	}
	r = server.NewResponse()
	v,err := models.GetAllAirportRailwayInfoByCondition(queryName,region,trafficType)
	paramOutput := make(map[string]interface{})
	if err != nil {
		paramOutput["msg"] = "操作失败"
		paramOutput["status"] = server.RESCODE__500.Int()
		result := []byte{}
		result, err = json.Marshal(paramOutput)
		r.ResponseJSON = result
		r.ResponeCode = server.RESCODE__500
	}else {
		paramOutput["msg"] = "操作成功"
		paramOutput["status"] = server.RESCODE__200.Int()
		paramOutput["data"] = v
		result := []byte{}
		result, err = json.Marshal(paramOutput)
		r.ResponseJSON = result
		r.ResponeCode = server.RESCODE__200
	}
	return r, nil
}

func (c *AirportRailwayInfoController) AddHotNum(params map[string]interface{}) (r *server.Response, err error) {
	
	val_airportCode := params["airportCode"]
	
	paramOutput := make(map[string]interface{})
	
	if val_airportCode == nil{
		paramOutput["msg"] = "参数不能为空"
		paramOutput["status"] = server.RESCODE__500.Int()
		result := []byte{}
		result, err = json.Marshal(paramOutput)
		r.ResponseJSON = result
		r.ResponeCode = server.RESCODE__500
	}else {
		r = server.NewResponse()
		_,err := models.UpdateHotNumByThreeCode(val_airportCode.(string))
		if err != nil {
			paramOutput["msg"] = "操作失败"
			paramOutput["status"] = server.RESCODE__500.Int()
			result := []byte{}
			result, err = json.Marshal(paramOutput)
			r.ResponseJSON = result
			r.ResponeCode = server.RESCODE__500
		}else {
			paramOutput["msg"] = "操作成功"
			paramOutput["status"] = server.RESCODE__200.Int()
			result := []byte{}
			result, err = json.Marshal(paramOutput)
			r.ResponseJSON = result
			r.ResponeCode = server.RESCODE__200
		}
	}
	return r, nil
}

func (c *AirportRailwayInfoController) QueryAirportByAirportCode(params map[string]interface{}) (r *server.Response, err error) {
	
	val_airportCode := params["airportCode"]
	
	paramOutput := make(map[string]interface{})
	
	if val_airportCode == nil{
		paramOutput["msg"] = "参数不能为空"
		paramOutput["status"] = server.RESCODE__500.Int()
		result := []byte{}
		result, err = json.Marshal(paramOutput)
		r.ResponseJSON = result
		r.ResponeCode = server.RESCODE__500
	}else {
		r = server.NewResponse()
		v,err := models.GetAirportRailwayInfoByThreeCode(val_airportCode.(string))
		if err != nil {
			paramOutput["msg"] = "操作失败"
			paramOutput["status"] = server.RESCODE__500.Int()
			result := []byte{}
			result, err = json.Marshal(paramOutput)
			r.ResponseJSON = result
			r.ResponeCode = server.RESCODE__500
		}else {
			paramOutput["msg"] = "操作成功"
			paramOutput["status"] = server.RESCODE__200.Int()
			paramOutput["data"] = v
			result := []byte{}
			result, err = json.Marshal(paramOutput)
			r.ResponseJSON = result
			r.ResponeCode = server.RESCODE__200
		}
	}
	return r, nil
	
}

// Post ...
// @Title Post
// @Description create AirportRailwayInfo
// @Param	body		body 	models.AirportRailwayInfo	true		"body for AirportRailwayInfo content"
// @Success 201 {int} models.AirportRailwayInfo
// @Failure 403 body is empty
// @router / [post]
func (c *AirportRailwayInfoController) Post() {
	var v models.AirportRailwayInfo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddAirportRailwayInfo(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get AirportRailwayInfo by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.AirportRailwayInfo
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AirportRailwayInfoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetAirportRailwayInfoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get AirportRailwayInfo
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.AirportRailwayInfo
// @Failure 403
// @router / [get]
func (c *AirportRailwayInfoController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllAirportRailwayInfo(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the AirportRailwayInfo
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.AirportRailwayInfo	true		"body for AirportRailwayInfo content"
// @Success 200 {object} models.AirportRailwayInfo
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AirportRailwayInfoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.AirportRailwayInfo{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateAirportRailwayInfoById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the AirportRailwayInfo
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AirportRailwayInfoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteAirportRailwayInfo(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

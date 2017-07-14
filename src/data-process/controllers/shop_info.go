package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"airport-cloud/base-info/models"

	"github.com/astaxie/beego"
	"log"
	"code.aliyun.com/wyunshare/thrift-server/gen-go/server"
)

// ShopInfoController operations for ShopInfo
type ShopInfoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ShopInfoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *ShopInfoController) GetShopInfoById(param map[string]interface{}) (r *server.Response, err error) {

	id, _ := strconv.Atoi(param["id"].(string))
	v, err := models.GetShopInfoById(id)

	r = server.NewResponse()
	paramOutput := make(map[string]interface{})
	if err != nil {
		paramOutput["msg"] = "操作失败"
		paramOutput["status"] = server.RESCODE__500.Int()
		paramOutput["data"] = v
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

func (c *ShopInfoController) GetShopInfoAll(queryParams map[string]interface{}) (r *server.Response, err error) {

	l, total, err := models.GetAllShopInfoNew(queryParams)
	data := make(map[string]interface{})
	data["rows"] = l
	data["total"] = total
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

func (c *ShopInfoController) AddShopInfo(queryParams []byte) (r *server.Response, err error) {

	var v models.ShopInfo
	r = server.NewResponse()
	paramOutput := make(map[string]interface{})
	json.Unmarshal(queryParams, &v)

	if _, err := models.AddShopInfo(&v); err == nil {
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

func (c *ShopInfoController) UpdateShopInfoById(queryParams []byte) (r *server.Response, err error) {

	var v models.ShopInfo
	r = server.NewResponse()
	paramOutput := make(map[string]interface{})
	json.Unmarshal(queryParams, &v)

	if err := models.UpdateShopInfoById(&v); err == nil {
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

func (c *ShopInfoController) DeleteShopInfoInBatch(queryParams []byte) (r *server.Response, err error) {

	paramInput := make(map[string]interface{})
	json.Unmarshal(queryParams, &paramInput)
	ids := paramInput["data"].([]interface{})
	log.Println(ids)

	var v models.ShopInfo
	r = server.NewResponse()
	paramOutput := make(map[string]interface{})
	json.Unmarshal(queryParams, &v)

	if err := models.DeleteShopInfoInBatch(ids); err == nil {
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

// Post ...
// @Title Post
// @Description create ShopInfo
// @Param	body		body 	models.ShopInfo	true		"body for ShopInfo content"
// @Success 201 {int} models.ShopInfo
// @Failure 403 body is empty
// @router / [post]
func (c *ShopInfoController) Post() {
	var v models.ShopInfo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddShopInfo(&v); err == nil {
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
// @Description get ShopInfo by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ShopInfo
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ShopInfoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetShopInfoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get ShopInfo
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ShopInfo
// @Failure 403
// @router / [get]
func (c *ShopInfoController) GetAll() {
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

	l, err := models.GetAllShopInfo(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ShopInfo
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ShopInfo	true		"body for ShopInfo content"
// @Success 200 {object} models.ShopInfo
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ShopInfoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.ShopInfo{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateShopInfoById(&v); err == nil {
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
// @Description delete the ShopInfo
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ShopInfoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteShopInfo(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

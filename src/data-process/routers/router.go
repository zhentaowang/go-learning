// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"data-process/controllers"

	"github.com/astaxie/beego"
	"encoding/json"
	"code.aliyun.com/wyunshare/thrift-server/gen-go/server"
	"code.aliyun.com/wyunshare/thrift-server/business"
	"fmt"
)

type BusinessServiceImpl struct {
}

func (msi *BusinessServiceImpl) Handle(operation string, paramJSON []byte) (r *server.Response, err error) {

	paramInput := make(map[string]interface{})
	json.Unmarshal(paramJSON, &paramInput)
	fmt.Println("-->Send:",paramJSON)

	c1 := controllers.AirportRailwayInfoController{}
	c2 := controllers.BasShopController{}
	c3 := controllers.LocationInfoController{}
	c4 := controllers.ShopInfoController{}

	switch operation {

	case "getAirportRailwayInfoAll":
		return c1.GetAirportRailwayInfoAll(paramInput)
	case "getBasShopAll":
		return c2.GetBasShopAll(paramInput)
	case "getLocationInfoById":
		return c3.GetLocationInfoById(paramInput)
	case "getLocationInfoAll":
		return c3.GetLocationInfoAll(paramInput)
	case "addLocationInfo":
		return c3.AddLocationInfo(paramJSON)
	case "updateLocationInfoById":
		return c3.UpdateLocationInfoById(paramJSON)
	case "deleteLocationInfoInBatch":
		return c3.DeleteLocationInfoInBatch(paramJSON)
	case "getShopInfoById":
		return c4.GetShopInfoById(paramInput)
	case "getShopInfoAll":
		return c4.GetShopInfoAll(paramInput)
	case "addShopInfo":
		return c4.AddShopInfo(paramJSON)
	case "updateShopInfoById":
		return c4.UpdateShopInfoById(paramJSON)
	case "deleteShopInfoInBatch":
		return c4.DeleteShopInfoInBatch(paramJSON)
	default:
		paramOutput := make(map[string]interface{})
		paramOutput["msg"] = "not found this method"
		paramOutput["status"] = server.RESCODE__404.Int()
		result := []byte{}
		result, err = json.Marshal(paramOutput)
		r.ResponseJSON = result
		r.ResponeCode = server.RESCODE__404
		return r,nil
	}
}

func GetHandler() (handler business.IBusinessService) {
	handler = &BusinessServiceImpl{}
	return handler
}

func init() {
	ns := beego.NewNamespace("/data-process",

		beego.NSNamespace("/bas_area",
			beego.NSInclude(
				&controllers.BasAreaController{},
			),
		),

		beego.NSNamespace("/bas_shop",
			beego.NSInclude(
				&controllers.BasShopController{},
			),
		),

		beego.NSNamespace("/location_info",
			beego.NSInclude(
				&controllers.LocationInfoController{},
			),
		),

		beego.NSNamespace("/shop_info",
			beego.NSInclude(
				&controllers.ShopInfoController{},
			),
		),

		beego.NSNamespace("/bas_traffic_site",
			beego.NSInclude(
				&controllers.BasTrafficSiteController{},
			),
		),

		beego.NSNamespace("/airport_railway_info",
			beego.NSInclude(
				&controllers.AirportRailwayInfoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

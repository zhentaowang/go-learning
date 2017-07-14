package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["data-process/controllers:AirportRailwayInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:AirportRailwayInfoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:AirportRailwayInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:AirportRailwayInfoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:AirportRailwayInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:AirportRailwayInfoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:AirportRailwayInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:AirportRailwayInfoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:AirportRailwayInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:AirportRailwayInfoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:BasAreaController"] = append(beego.GlobalControllerRouter["data-process/controllers:BasAreaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:BasAreaController"] = append(beego.GlobalControllerRouter["data-process/controllers:BasAreaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:BasAreaController"] = append(beego.GlobalControllerRouter["data-process/controllers:BasAreaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:BasAreaController"] = append(beego.GlobalControllerRouter["data-process/controllers:BasAreaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:BasAreaController"] = append(beego.GlobalControllerRouter["data-process/controllers:BasAreaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:BasShopController"] = append(beego.GlobalControllerRouter["data-process/controllers:BasShopController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:BasShopController"] = append(beego.GlobalControllerRouter["data-process/controllers:BasShopController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:BasShopController"] = append(beego.GlobalControllerRouter["data-process/controllers:BasShopController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:BasShopController"] = append(beego.GlobalControllerRouter["data-process/controllers:BasShopController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:BasShopController"] = append(beego.GlobalControllerRouter["data-process/controllers:BasShopController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:BasTrafficSiteController"] = append(beego.GlobalControllerRouter["data-process/controllers:BasTrafficSiteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:BasTrafficSiteController"] = append(beego.GlobalControllerRouter["data-process/controllers:BasTrafficSiteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:BasTrafficSiteController"] = append(beego.GlobalControllerRouter["data-process/controllers:BasTrafficSiteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:BasTrafficSiteController"] = append(beego.GlobalControllerRouter["data-process/controllers:BasTrafficSiteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:BasTrafficSiteController"] = append(beego.GlobalControllerRouter["data-process/controllers:BasTrafficSiteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:LocationInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:LocationInfoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:LocationInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:LocationInfoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:LocationInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:LocationInfoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:LocationInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:LocationInfoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:LocationInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:LocationInfoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:ShopInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:ShopInfoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:ShopInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:ShopInfoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:ShopInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:ShopInfoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:ShopInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:ShopInfoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:ShopInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:ShopInfoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:TerminalInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:TerminalInfoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:TerminalInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:TerminalInfoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:TerminalInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:TerminalInfoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:TerminalInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:TerminalInfoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["data-process/controllers:TerminalInfoController"] = append(beego.GlobalControllerRouter["data-process/controllers:TerminalInfoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}

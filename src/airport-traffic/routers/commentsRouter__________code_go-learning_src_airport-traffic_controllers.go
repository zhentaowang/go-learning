package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["airport-traffic/controllers:PrBusRouteController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrBusRouteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrBusRouteController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrBusRouteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrBusRouteController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrBusRouteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrBusRouteController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrBusRouteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrBusRouteController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrBusRouteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrBusStationController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrBusStationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrBusStationController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrBusStationController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrBusStationController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrBusStationController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrBusStationController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrBusStationController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrBusStationController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrBusStationController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingLotController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingLotController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingLotController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingLotController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingLotController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingLotController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingLotController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingLotController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingLotController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingLotController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingMeterController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingMeterController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingMeterController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingMeterController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingMeterController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingMeterController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingMeterController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingMeterController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingMeterController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingMeterController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingServiceController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingServiceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingServiceController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingServiceController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingServiceController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingServiceController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingServiceController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingServiceController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingServiceController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrParkingServiceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrTerminalPointController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrTerminalPointController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrTerminalPointController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrTerminalPointController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrTerminalPointController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrTerminalPointController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrTerminalPointController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrTerminalPointController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airport-traffic/controllers:PrTerminalPointController"] = append(beego.GlobalControllerRouter["airport-traffic/controllers:PrTerminalPointController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}

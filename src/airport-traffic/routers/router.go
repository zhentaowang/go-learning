// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"airport-traffic/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/airport-traffic",

		beego.NSNamespace("/pr_bus_route",
			beego.NSInclude(
				&controllers.PrBusRouteController{},
			),
		),

		beego.NSNamespace("/pr_bus_station",
			beego.NSInclude(
				&controllers.PrBusStationController{},
			),
		),

		beego.NSNamespace("/pr_parking_lot",
			beego.NSInclude(
				&controllers.PrParkingLotController{},
			),
		),

		beego.NSNamespace("/pr_parking_meter",
			beego.NSInclude(
				&controllers.PrParkingMeterController{},
			),
		),

		beego.NSNamespace("/pr_parking_service",
			beego.NSInclude(
				&controllers.PrParkingServiceController{},
			),
		),

		beego.NSNamespace("/pr_terminal_point",
			beego.NSInclude(
				&controllers.PrTerminalPointController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

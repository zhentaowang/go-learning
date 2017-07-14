package models

import "time"

type PrAssRouteStation struct {
	StationId      int64     `orm:"column(station_id)"`
	RouteId        int64     `orm:"column(route_id)"`
	StationOrdinal int       `orm:"column(station_ordinal);null"`
	ArrivedTime    string    `orm:"column(arrived_time);size(20);null"`
	LeaveTime      string    `orm:"column(leave_time);size(20);null"`
	Stay           string    `orm:"column(stay);size(10);null"`
	Mileage        string    `orm:"column(mileage);size(10);null"`
	GetOnOff       int8      `orm:"column(get_on_off);null"`
	Description    string    `orm:"column(description);size(200);null"`
	TicketPrice    float64   `orm:"column(ticket_price);null;digits(8);decimals(2)"`
	Currency       string    `orm:"column(currency);size(10);null"`
	IsDeleted      int8      `orm:"column(is_deleted);null"`
	CreateUser     int64     `orm:"column(create_user);null"`
	UpdateUser     int64     `orm:"column(update_user);null"`
	CreateTime     time.Time `orm:"column(create_time);type(datetime);null;auto_now_add"`
	UpdateTime     time.Time `orm:"column(update_time);type(datetime);null"`
}

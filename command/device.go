package command

import (
	_ "embed"
)

var (
	Devices = make(map[string][]interface{})
	//go:embed sqls/deviceCount.sql
	devicesCount string
	//go:embed sqls/deviceActive.sql
	deviceActive string

	deviceMysql = mysql{
		stat: devicesCount,
		data: Devices,
		callback: func(run string) []interface{} {
			return unmarshal(run, OrderMap{})
		},
	}
)

const (
	LastWeekSale   = "lastWeekSale"
	ThisWeekSale   = "thisWeekSale"
	GrowSale       = "growSale"
	ThisWeekActive = "thisWeekActive"
)

func init() {
	lastWeekSale()
	thisWeekSale()
	growSale()
	thisWeekActive()

	deviceMysql.sum()
}

func thisWeekActive() {
	deviceMysql.stat = deviceActive
	deviceMysql.load(ThisWeekActive, oneWeek, timestamp)
}

func growSale() {
	deviceMysql.grow(ThisWeekSale, LastWeekSale, GrowSale)
}

func thisWeekSale() {
	deviceMysql.load(ThisWeekSale, timestamp)
}

func lastWeekSale() {
	deviceMysql.load(LastWeekSale, oneWeek)
}

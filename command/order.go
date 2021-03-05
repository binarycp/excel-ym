package command

import (
	_ "embed"
)

var (
	Order = make(map[string][]interface{})
	//go:embed sqls/orderCount.sql
	orderCount string
	orderMysql = &mysql{
		stat: orderCount,
		data: Order,
		callback: func(run string) []interface{} {
			return unmarshal(run, OrderMap{})
		},
	}
)

const (
	TotalOrder               = "totalOrder"
	DoThisWeekFinishedOrder  = "doThisWeekFinishedOrder"
	DoLastWeekNewOrder       = "doLastWeekNewOrder"
	DoThisWeekNewOrder       = "doThisWeekNewOrder"
	DoLastWeekRepairingOrder = "doLastWeekRepairingOrder"
	DoThisWeekRepairingOrder = "doThisWeekRepairingOrder"
	DoLastWeekFinishedOrder  = "doLastWeekFinishedOrder"

	DoGrowFinishedOrder  = "doGrowFinishedOrder"
	DoGrowNewOrder       = "doGrowNewOrder"
	DoGrowRepairingOrder = "doGrowRepairingOrder"
)

type OrderMap struct {
	Province string
	Count    string
}

func init() {
	// 总数
	totalOrder()
	// 新开单上周
	doLastWeekNewOrder()
	// 新开单本周
	doThisWeekNewOrder()
	// 维修中上周
	doLastWeekRepairingOrder()
	// 维修中本周
	doThisWeekRepairingOrder()
	// 已完成上周
	doLastWeekFinishedOrder()
	// 已完成本周
	doThisWeekFinishedOrder()

	// 已完成增长
	doGrowFinishedOrder()
	// 新开单增长
	doGrowNewOrder()
	// 维修中增长
	doGrowRepairingOrder()

	// 总和
	orderMysql.sum()
}

func doGrowFinishedOrder() {
	orderMysql.grow(DoThisWeekFinishedOrder, DoLastWeekFinishedOrder, DoGrowFinishedOrder)
}

func doGrowNewOrder() {
	orderMysql.grow(DoThisWeekNewOrder, DoLastWeekNewOrder, DoGrowNewOrder)
}

func doGrowRepairingOrder() {
	orderMysql.grow(DoThisWeekRepairingOrder, DoLastWeekRepairingOrder, DoGrowRepairingOrder)
}

func totalOrder() {
	orderMysql.load(TotalOrder, "1,2", 0, timestamp)
}

func doThisWeekFinishedOrder() {
	orderMysql.load(DoThisWeekFinishedOrder, "2", 0, timestamp)
}

func doLastWeekNewOrder() {
	// 14日前到七日前(要算执行那天)
	orderMysql.load(DoLastWeekNewOrder, "1,2", twoWeek, oneWeek)
}

func doThisWeekNewOrder() {
	orderMysql.load(DoThisWeekNewOrder, "1,2", oneWeek, timestamp)
}

func doLastWeekRepairingOrder() {
	orderMysql.load(DoLastWeekRepairingOrder, "1", 0, oneWeek)
}

func doThisWeekRepairingOrder() {
	orderMysql.load(DoThisWeekRepairingOrder, "1", 0, timestamp)
}

func doLastWeekFinishedOrder() {
	orderMysql.load(DoLastWeekFinishedOrder, "2", 0, oneWeek)
}

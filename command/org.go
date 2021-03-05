package command

import (
	_ "embed"
)

var (
	Org = make(map[string][]interface{})
	//go:embed sqls/orgCount.sql
	orgCount string

	orgMysql = &mysql{
		stat: orgCount,
		data: Org,
		callback: func(run string) []interface{} {
			return unmarshal(run, OrderMap{})
		},
	}
)

const (
	LastWeekOrg = "LastWeekOrg"
	ThisWeekOrg = "thisWeekOrg"
	GrowOrg     = "growOrg"
)

func init() {
	lastWeekOrg()
	thisWeekOrg()
	growOrg()

	orgMysql.sum()
}

func growOrg() {
	orgMysql.grow(ThisWeekOrg, LastWeekOrg, GrowOrg)
}

func thisWeekOrg() {
	orgMysql.load(ThisWeekOrg, timestamp)
}

func lastWeekOrg() {
	orgMysql.load(LastWeekOrg, oneWeek)
}

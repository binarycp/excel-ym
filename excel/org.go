package excel

import (
	"excel-ym/command"
	"excel-ym/utils"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"strconv"
)

var (
	org = New("新开M站列表{$date}.xlsx", "template/org.xlsx")
	// 字段对应哪一列
	orgExcelMap = map[string]string{
		command.LastWeekOrg: "B",
		command.GrowOrg:     "D",
	}
)

func init() {
	orgExcel(org)
}

func orgExcel(info *Info) {
	for k, v := range orgConvert(info.F) {
		utils.Err(info.F.SetCellValue(SHEET, k, v))
	}
	utils.Err(info.F.SaveAs(info.Path))
}

func orgConvert(f *excelize.File) map[string]interface{} {
	c := make([]convert, 0)
	index := make(map[string]int)

	// 优先处理包含全部列的内容
	for k, v := range command.Org[command.ThisWeekOrg] {
		column := k + 3
		orderMap := v.(command.OrderMap)
		index[orderMap.Province] = column
		atoi, err := strconv.Atoi(orderMap.Count)
		utils.Err(err)
		c = append(c, convert{
			row:    "A",
			column: column,
			value:  orderMap.Province,
		}, convert{
			row:    "C",
			column: column,
			value:  atoi,
		})
		utils.Err(f.SetRowHeight(SHEET, column, 30))
	}

	// 处理其他的数据
	for k, v := range command.Org {
		if k == command.ThisWeekOrg {
			continue
		}

		exist := make(map[string]int)
		for _, vv := range v {
			orderMap := vv.(command.OrderMap)
			atoi, err := strconv.Atoi(orderMap.Count)
			utils.Err(err)
			exist[orderMap.Province] = 0
			c = append(c, convert{
				row:    orgExcelMap[k],
				column: index[orderMap.Province],
				value:  atoi,
			})
		}

		for kk, vv := range index {
			if _, ok := exist[kk]; !ok {
				c = append(c, convert{
					row:    orgExcelMap[k],
					column: vv,
					value:  0,
				})
			}
		}
	}

	return ToMap(c)
}

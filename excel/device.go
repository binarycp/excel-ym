package excel

import (
	"excel-ym/command"
	"excel-ym/utils"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"strconv"
)

var (
	device = New("新增设备列表{$date}.xlsx", "template/device.xlsx")
	// 字段对应哪一列
	deviceExcelMap = map[string]string{
		command.LastWeekSale:   "B",
		command.GrowSale:       "D",
		command.ThisWeekActive: "E",
	}
)

func init() {
	deviceExcel(device)
}

func deviceExcel(info *Info) {
	for k, v := range deviceConvert(info.F) {
		utils.Err(info.F.SetCellValue(SHEET, k, v))
	}
	utils.Err(info.F.SaveAs(info.Path))
}

func deviceConvert(f *excelize.File) map[string]interface{} {
	c := make([]convert, 0)
	index := make(map[string]int)

	// 优先处理包含全部列的内容
	for k, v := range command.Devices[command.ThisWeekSale] {
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
	for k, v := range command.Devices {
		if k == command.ThisWeekSale {
			continue
		}

		exist := make(map[string]int)
		for _, vv := range v {
			orderMap := vv.(command.OrderMap)
			atoi, err := strconv.Atoi(orderMap.Count)
			utils.Err(err)
			exist[orderMap.Province] = 0
			c = append(c, convert{
				row:    deviceExcelMap[k],
				column: index[orderMap.Province],
				value:  atoi,
			})
		}

		for kk, vv := range index {
			if _, ok := exist[kk]; !ok {
				c = append(c, convert{
					row:    deviceExcelMap[k],
					column: vv,
					value:  0,
				})
			}
		}
	}

	return ToMap(c)
}

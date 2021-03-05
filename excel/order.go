package excel

import (
	"excel-ym/command"
	"excel-ym/utils"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"strconv"
)

var (
	Order = New("工单开单详情列表{$date}.xlsx", "template/order.xlsx")
	// 字段对应哪一列
	orderExcelMap = map[string]string{
		command.DoLastWeekNewOrder:       "C",
		command.DoThisWeekNewOrder:       "D",
		command.DoGrowNewOrder:           "E",
		command.DoLastWeekRepairingOrder: "F",
		command.DoThisWeekRepairingOrder: "G",
		command.DoGrowRepairingOrder:     "H",
		command.DoLastWeekFinishedOrder:  "I",
		command.DoThisWeekFinishedOrder:  "J",
		command.DoGrowFinishedOrder:      "K",
	}
)

func init() {
	orderExcel(Order)
}

// 生成excel文件
func orderExcel(info *Info) {
	for k, v := range orderConvert(info.F) {
		utils.Err(info.F.SetCellValue(SHEET, k, v))
	}

	// 字体居左没有生效
	//style, err := info.F.NewStyle(&excelize.Style{
	//	Alignment: &excelize.Alignment{
	//		Horizontal: "left",
	//	},
	//})
	//
	//utils.Err(err)
	//utils.Err(info.F.SetCellStyle("Sheet1", "A3", "D12", style))
	utils.Err(info.F.SaveAs(info.Path))
}

/**
关联转换
哪一列
哪一行
A3:L24
province:行数
字段:列数
*/
func orderConvert(f *excelize.File) map[string]interface{} {
	c := make([]convert, 0)
	index := make(map[string]int)

	// 优先处理包含全部列的内容
	for k, v := range command.Order[command.TotalOrder] {
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
			row:    "B",
			column: column,
			value:  atoi,
		})
		utils.Err(f.SetRowHeight(SHEET, column, 30))
	}

	// 处理其他的数据
	for k, v := range command.Order {
		if k == command.TotalOrder {
			continue
		}

		exist := make(map[string]int)
		for _, vv := range v {
			orderMap := vv.(command.OrderMap)
			atoi, err := strconv.Atoi(orderMap.Count)
			utils.Err(err)
			exist[orderMap.Province] = 0
			c = append(c, convert{
				row:    orderExcelMap[k],
				column: index[orderMap.Province],
				value:  atoi,
			})
		}

		for kk, vv := range index {
			if _, ok := exist[kk]; !ok {
				c = append(c, convert{
					row:    orderExcelMap[k],
					column: vv,
					value:  0,
				})
			}
		}
	}

	return ToMap(c)
}

// 设置头部（不需要了，直接打开模板文件，写入数据后另存为即可）
func orderHeader(f *excelize.File) {
	utils.Err(f.SetCellValue(SHEET, "A1", "区域（省）"))
	utils.Err(f.SetCellValue(SHEET, "B1", "开展M站治理业务总数"))
	utils.Err(f.SetCellValue(SHEET, "C1", "新开单"))
	utils.Err(f.SetCellValue(SHEET, "C2", "上周"))
	utils.Err(f.SetCellValue(SHEET, "D2", "本周"))
	utils.Err(f.SetCellValue(SHEET, "E2", "增长"))
	utils.Err(f.SetCellValue(SHEET, "F1", "维修中"))
	utils.Err(f.SetCellValue(SHEET, "F2", "上周"))
	utils.Err(f.SetCellValue(SHEET, "G2", "本周"))
	utils.Err(f.SetCellValue(SHEET, "H2", "增长"))
	utils.Err(f.SetCellValue(SHEET, "I1", "已完成"))
	utils.Err(f.SetCellValue(SHEET, "I2", "上周"))
	utils.Err(f.SetCellValue(SHEET, "J2", "本周"))
	utils.Err(f.SetCellValue(SHEET, "K2", "增长"))
	utils.Err(f.SetCellValue(SHEET, "L1", "分析/措施"))

	// 合并单元格
	utils.Err(f.MergeCell(SHEET, "A1", "A2"))
	utils.Err(f.MergeCell(SHEET, "B1", "B2"))
	utils.Err(f.MergeCell(SHEET, "L1", "L2"))
	utils.Err(f.MergeCell(SHEET, "C1", "E1"))
	utils.Err(f.MergeCell(SHEET, "F1", "H1"))
	utils.Err(f.MergeCell(SHEET, "I1", "K1"))

	// 设置样式
	style, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D0CECE"],"pattern":0},"alignment":{"horizontal":"center","vertical":"center"}}`)
	utils.Err(err)
	utils.Err(f.SetCellStyle(SHEET, "A1", "L2", style))

	// 设置宽高
	utils.Err(f.SetSheetFormatPr(SHEET, excelize.DefaultColWidth(13), excelize.DefaultRowHeight(30)))
	utils.Err(f.SetSheetPrOptions(SHEET, excelize.AutoPageBreaks(true)))

	utils.Err(f.SetColWidth(SHEET, "A", "L", 13))
	utils.Err(f.SetRowHeight(SHEET, 1, 30))
	utils.Err(f.SetRowHeight(SHEET, 2, 30))
}

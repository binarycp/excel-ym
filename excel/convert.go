package excel

import "strconv"

/**
关联转换
哪一列
哪一行
A3:L24
province:行数
字段:列数
*/
type convert struct {
	row    string
	column int
	value  interface{}
}

// 转为map
func ToMap(converts []convert) map[string]interface{} {
	m := make(map[string]interface{})
	for _, v := range converts {
		m[v.row+strconv.Itoa(v.column)] = v.value
	}
	return m
}

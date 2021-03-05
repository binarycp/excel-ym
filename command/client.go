package command

import (
	"excel-ym/app"
	"excel-ym/utils"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	cli       *app.Cli
	timestamp int64
	now       time.Time
	oneWeek   int64
	twoWeek   int64
)

type mysql struct {
	stat     string
	data     map[string][]interface{}
	callback func(run string) []interface{}
}

func (m *mysql) load(key string, a ...interface{}) {
	sql := fmt.Sprintf(m.stat, a...)
	//fmt.Println(sql)
	run, err := cli.Run(command(sql))
	utils.Err(err)
	m.data[key] = m.callback(run)
	//if key == ThisWeekOrg {
	//	fmt.Println(sql)
	//	fmt.Println(run)
	//}
}

func (m *mysql) grow(thisWeek string, lastWeek string, key string) {
	result := make(map[string]int)
	for _, v := range m.data[thisWeek] {
		orderMap := v.(OrderMap)
		atoi, err := strconv.Atoi(orderMap.Count)
		utils.Err(err)
		result[orderMap.Province] = atoi
	}

	for _, v := range m.data[lastWeek] {
		orderMap := v.(OrderMap)
		atoi, err := strconv.Atoi(orderMap.Count)
		utils.Err(err)
		if _, ok := result[orderMap.Province]; ok {
			result[orderMap.Province] -= atoi
		} else {
			result[orderMap.Province] = 0 - atoi
		}
	}
	maps := make([]interface{}, 0)
	for k, v := range result {
		maps = append(maps, OrderMap{
			Province: k,
			Count:    strconv.Itoa(v),
		})
	}

	m.data[key] = maps
}

func (m *mysql) sum() {
	for k, r := range m.data {
		total := 0
		for _, v := range r {
			orderMap := v.(OrderMap)
			atoi, err := strconv.Atoi(orderMap.Count)
			utils.Err(err)
			total += atoi
		}
		m.data[k] = append(m.data[k], OrderMap{
			Province: "总计",
			Count:    strconv.Itoa(total),
		})
	}
}

func init() {
	cli = app.New(app.Conf.Host, app.Conf.UserName, app.Conf.Password, app.Conf.Port)
	now = time.Now()
	timestamp = now.Unix()
	oneWeek = now.AddDate(0, 0, -6).Unix()
	twoWeek = now.AddDate(0, 0, -13).Unix()
}

func command(sql string) string {
	return fmt.Sprintf(app.Connect, app.Conf.MysqlHost, app.Conf.MysqlPort, app.Conf.MysqlUserName, app.Conf.MysqlPassword, app.Conf.MYSQLDB, sql)
}

//解析mysql的结果，要求i是struct,并且字段类型是string
func unmarshal(data string, i interface{}) []interface{} {
	ret := make([]interface{}, 0)

	// 移除结尾的换行符，避免以换行分隔时出现空数据
	p := []byte(data)
	if p[len(p)-1] == '\n' {
		data = string(p[:len(p)-1])
	}

	// mysql的数据以换行分隔行，制表符分隔列
	split := strings.Split(data, "\n")
	if len(split) < 2 {
		ret = append(ret, i)
		return ret
	}

	typeOf := reflect.TypeOf(i)
	fields := make(map[string]int)
	for k, v := range split {
		result := strings.Split(v, "\t")
		if k == 0 {
			for kk, vv := range result {
				fields[vv] = kk
			}
		} else {
			e := reflect.New(typeOf).Elem()
			for kk, vv := range fields {
				fieldByName := caseInsensitiveFieldByName(e, kk)
				// 判断字段是否在结构体中定义
				if fieldByName.Kind() != reflect.Invalid {
					fieldByName.Set(reflect.ValueOf(result[vv]))
				}
			}
			ret = append(ret, e.Interface())
		}
	}

	return ret
}

// 首字母大写
func UcWords(s string) string {
	p := []byte(s)
	if len(p) > 0 && p[0] >= 97 && p[0] <= 122 {
		p[0] = p[0] - 32
	}
	return string(p)
}

// 反射，不区分大小写
func caseInsensitiveFieldByName(v reflect.Value, name string) reflect.Value {
	name = strings.ToLower(name)
	return v.FieldByNameFunc(func(n string) bool { return strings.ToLower(n) == name })
}

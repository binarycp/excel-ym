package app

import (
	_ "embed"
	"encoding/json"
	"os"
)

type Config struct {
	Ssh   `json:"ssh"`
	Mysql `json:"mysql"`
}

type Ssh struct {
	Host     string `json:"host"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

type Mysql struct {
	MysqlHost     string `json:"mysql_host"`
	MysqlUserName string `json:"mysql_user_name"`
	MysqlPassword string `json:"mysql_password"`
	MysqlPort     int    `json:"mysql_port"`
	MYSQLDB       string `json:"mysqldb"`
}

var (
	//go:embed config.json
	c    []byte
	Conf = Config{}
)

func init() {
	err := json.Unmarshal(c, &Conf)
	if err != nil {
		println("没有读取到配置")
		os.Exit(1)
	}
}

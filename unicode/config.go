package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/stackcats/gosh"
)

// MysqlHelper ...
type MysqlConfig struct {
	Host            string
	Net             string
	UserName        string
	Password        string
	Database        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	Parameters      string
}

// Config 发入consul的配置...
type Config struct {
	MysqlURL               string `consul:"db/mysql/mysqlurl"`
	Haproxy                string `consul:"haproxy"`
	ERC20Port              string `consul:"service/erc20/port"`
	Payee                  string `consul:"service/erc20/payee"`
	KeystorePort           string `consul:"service/user-keystore/port"`
	Geth                   string `consul:"service/geth/address"`
	RedisURI               string `consul:"redis/cache/host"`
	RedisPassword          string `consul:"redis/cache/password"`
	ExpireOrderCallBackURL string `consul:"service/points-mall/order/expire"`
	TokenFlowPort          string `consul:"service/token-flow/port"`
}

var config *Config
var mysqlConfig *MysqlConfig

func init() {

	config = &Config{
		MysqlURL: "192.168.3.18:3306",
		Haproxy:  "192.168.3.18",
	}

	consul := strings.Trim(os.Getenv("consul"), " ")
	if consul != "" {
		c := gosh.DefaultConfig()
		c.Address = consul + ":8500"
		g, err := gosh.NewClient(c)
		if err != nil {
			log.Fatal(err)
		}

		err = g.Unmarshal(config)
		if err != nil {
			log.Fatal(err)
		}
	}

	mysqlConfig = &MysqlConfig{
		Host:            config.MysqlURL,
		Net:             "tcp",
		UserName:        "golo",
		Password:        "launch888",
		Database:        "golo_eva",
		MaxOpenConns:    100,
		MaxIdleConns:    10,
		ConnMaxLifetime: time.Second * 10,
		Parameters:      "charset=utf8&parseTime=True&loc=Local",
	}
}

// NewConfig ...
func NewConfig() *Config {
	return config
}

// NewMysqlConfig ...
func NewMysqlConfig() *MysqlConfig {
	return mysqlConfig
}

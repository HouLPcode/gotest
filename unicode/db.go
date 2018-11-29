package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shopspring/decimal"
	"github.com/stackcats/iris/utils"
)

var log = utils.NewLogger()

const (
	//=======mysql表名
	CommodityTableName     = "commodity"
	OrderTableName         = "order"
	RecipientTableName     = "recipient_info"
	ViolationTableName     = "eva_violation"
	ViolationCityTableName = "eva_violation_city"
)

type Report struct {
	ReportID   int    `gorm:"column:report_id" json:"report_id"`
	UserId     int    `gorm:"column:user_id" json:"user_id"`
	ReportNo   string `gorm:"column:report_no" json:"report_no"`
	Vin        string `gorm:"column:vin" json:"vin"`
	Mileage    int    `gorm:"column:mileage" json:"mileage"`
	Content    string `gorm:"column:content" json:"content"`
	ReportType int    `gorm:"report_type" json:"report_type"`
	OrderId    int    `gorm:"column:order_id" json:"order_id"`
	ParentID   int    `gorm:"column:parent_id" json:"parent_id"`
}


// C ...
type C interface {
	CName() string
}

func GetConn() (*gorm.DB, error) {
	conf := config.NewMysqlConfig()
	var err error
	username := conf.UserName
	password := conf.Password
	net := conf.Net
	host := conf.Host
	database := conf.Database
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	//parse := "charset=utf8&parseTime=True&loc=Local"
	parameters := conf.Parameters
	dburl := fmt.Sprintf("%s:%s@%s(%s)/%s?%s", username, password, net, host, database, parameters)
	//log.Debug("url ", dburl)
	db, err := gorm.Open("mysql", dburl)
	if err != nil {
		log.Fatalf("数据库连接异常：%v", err)
		panic(err)
	}
	db.LogMode(true)
	return db, nil
}

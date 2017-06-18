package mysql

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql" // Connector
)

var mysqlDb = "coloso"
var mysqlUser = "root"
var mysqlPassword = "123456"

// Mysql client
var Client *gorm.DB


func init() {
  connectedDb, err := gorm.Open("mysql", mysqlUser + ":" + mysqlPassword + "@/" + mysqlDb + "?charset=utf8&parseTime=True&loc=Local")

  if err != nil {
    panic(err)
  }

  Client = connectedDb
}

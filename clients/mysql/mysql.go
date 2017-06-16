package mysql

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql" // Connector
)

var mysqlDb = "coloso"
var mysqlUser = "root"
var mysqlPassword = "123456"
var client gorm.DB

// GetClient - Singleton
func GetClient() gorm.DB {
  if client != nil {
    return client
  }

  db, err := gorm.Open("mysql", mysqlUser + ":" + mysqlPassword + "@/" + mysqlDb + "?charset=utf8&parseTime=True&loc=Local")

  return db
}

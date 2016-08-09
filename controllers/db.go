package controllers

import (

  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
  "log"

  "github.com/shenxl/mockdata/models"
)

type DBController struct {
	DB *gorm.DB
}

func (dc *DBController) InitDB() {
	var err error

  dc.DB, err = gorm.Open("mysql","root:kingsoft@/mockdata?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatalf("Error when connect database, the error is '%v'", err)
	}
	dc.DB.LogMode(true)
}

func (dc *DBController) GetDB() *gorm.DB {
  return dc.DB
}

func (dc *DBController) InitSchema(){
  dc.DB.AutoMigrate(&models.Company{})
}

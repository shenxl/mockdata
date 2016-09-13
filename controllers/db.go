package controllers

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/shenxl/mockdata/models"
)

type DBController struct {
	DB *gorm.DB
}

func (dc *DBController) InitDB(connstr string, show_log bool) {
	var err error

	// dc.DB, err = gorm.Open("mysql", "root:kingsoft@tcp(192.168.132.105:3306)/mockdata?charset=utf8&parseTime=True")
	// dc.DB, err = gorm.Open("mysql", "wpsstat:stat+0756@tcp(127.0.0.1:3306)/wpsupdate?loc=Local&parseTime=True&charset=utf8")
	dc.DB, err = gorm.Open("mysql", connstr)
	if err != nil {
		log.Fatalf("Error when connect database, the error is '%v'", err)
	}
	dc.DB.LogMode(show_log)
}

func (dc *DBController) GetDB() *gorm.DB {
	return dc.DB
}

func (dc *DBController) InitSchema() {
	dc.DB.AutoMigrate(&models.Company{}, &models.CompanySn{}, &models.Order{}, &models.Group{})
}

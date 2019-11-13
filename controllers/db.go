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

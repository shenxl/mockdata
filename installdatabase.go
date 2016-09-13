package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shenxl/mockdata/controllers"
	"github.com/shenxl/mockdata/models"
)

func main() {
	dc := controllers.DBController{}
	dc.InitDB(
		
	dc.InitSchema()
	results := []models.Company{}
	err := dc.GetDB().Find(&results).Error

	if err != nil {
		return
	}

}

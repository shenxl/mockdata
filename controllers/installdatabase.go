package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"

	"time"
)

var loginstall = logging.MustGetLogger("Install")

type InstallController struct {
	DB *gorm.DB
}

func (ac *InstallController) SetDB(d *gorm.DB) {
	ac.DB = d
	ac.DB.LogMode(true)
}

func (ac *InstallController) Install(c *gin.Context) {

	// results := []models.Company{}
	// err := ac.DB.Find(&results).Error
	//
	// if err != nil {
	// 	logcompany.Debugf("Error when looking up companyList, the error is '%v'", err)
	// 	res := gin.H{
	// 		"status": "404",
	// 		"error":  "No company found",
	// 	}
	// 	c.JSON(404, res)
	// 	return
	//
	type sourceData struct {
		Sales        string
		Sn           string
		SnNumber     int
		CompanyGroup string
		Company      string
		Region       string
		Province     string
		City         string
		Address      string
		Industry     string
		ServiceYear  int
		ServiceDate  time.Time
	}
	data := []sourceData{}
	err := ac.DB.Find(&data).Error

	if err != nil {
		logcompany.Debugf("Error when looking up companyList, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "No company found",
		}
		c.JSON(404, res)
		return
	}
	content := gin.H{
		"status":   "200",
		"success":  true,
		"length":   len(data),
		"companys": data,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

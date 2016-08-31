package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/shenxl/mockdata/models"
)

var logcompanymonthly = logging.MustGetLogger("CompanyMonthly")

type CompanyMonthlyController struct {
	DB *gorm.DB
}

func (ac *CompanyMonthlyController) SetDB(d *gorm.DB) {
	ac.DB = d
	ac.DB.LogMode(true)
}

func (ac *CompanyMonthlyController) AllByHere(c *gin.Context) {

	month := c.Params.ByName("month")
	year := c.Params.ByName("year")
	type ResultData struct {
		CompanyId   int64  `json:"companyid"`
		CompanyName string `json:"name"`
		Region      string `json:"region"`
		Industry    string `json:"industry"`
		ActivityMax int64  `json:"activity_Max"`
		ActivityAvg int64  `json:"activity_avg"`
		Install     int64  `json:"install"`
	}
	results := []ResultData{}

	err := ac.DB.Table("company_monthly as m").
		Select("c.id as company_id,c.company as company_name,c.region,c.industry,m.year,m.month,m.activity_max,m.activity_avg").
		Joins("left join company as c on m.company = c.id ").
		Where("year=?", year).
		Where("month=? ", month).
		Scan(&results).Error
	// err := ac.DB.Table("server").
	// 	Select("server.* , company.*").
	// 	Joins("left join company on server.company = company.id ").
	// 	Scan(&results).Error
	//
	// err := ac.DB.Find(&results).Error

	if err != nil {
		logcompanymonthly.Debugf("Error when looking up company_monthly List, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "No company_monthly list found",
		}
		c.JSON(404, res)
		return
	}
	content := gin.H{
		"status":             "200",
		"success":            true,
		"count":              len(results),
		"companymonthlyList": results,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

func (ac *CompanyMonthlyController) List(c *gin.Context) {

	results := []models.CompanyMonthly{}
	err := ac.DB.Find(&results).Error

	if err != nil {
		logcompanymonthly.Debugf("Error when looking up company_monthly List, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "No company_monthly list found",
		}
		c.JSON(404, res)
		return
	}
	content := gin.H{
		"status":             "200",
		"success":            true,
		"companymonthlyList": results,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

func (ac *CompanyMonthlyController) GetCompanyMonthly(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
	entity := models.CompanyMonthly{}
	err := ac.DB.Where("id=?", id).Find(&entity).Error

	if err != nil {
		logcompanymonthly.Debugf("Error when looking up company_monthly, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "company_monthly not found",
		}
		c.JSON(404, res)
		return
	}

	content := gin.H{
		"status":         "201",
		"success":        true,
		"companymonthly": entity,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

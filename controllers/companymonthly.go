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

// Get a company
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

package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/shenxl/mockdata/models"
)

var logcompanydaily = logging.MustGetLogger("CompanyDaily")

type CompanyDailyController struct {
	DB *gorm.DB
}

func (ac *CompanyDailyController) SetDB(d *gorm.DB) {
	ac.DB = d
	ac.DB.LogMode(true)
}

func (ac *CompanyDailyController) List(c *gin.Context) {

	results := []models.CompanyDaily{}
	err := ac.DB.Find(&results).Error

	if err != nil {
		logcompanydaily.Debugf("Error when looking up company_daily List, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "No company_daily list found",
		}
		c.JSON(404, res)
		return
	}
	content := gin.H{
		"status":           "200",
		"result":           "Success",
		"companydailyList": results,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

// Get a company
func (ac *CompanyDailyController) GetCompanyDaily(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
	entity := models.CompanyDaily{}
	err := ac.DB.Where("id=?", id).Find(&entity).Error

	if err != nil {
		logcompanydaily.Debugf("Error when looking up company_daily, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "User not found",
		}
		c.JSON(404, res)
		return
	}

	content := gin.H{
		"status":       "201",
		"result":       "Success",
		"companydaily": entity,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

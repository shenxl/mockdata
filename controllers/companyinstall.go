package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/shenxl/mockdata/models"
)

var logcompanyinstall = logging.MustGetLogger("CompanyInstall")

type CompanyInstallController struct {
	DB *gorm.DB
}

func (ac *CompanyInstallController) SetDB(d *gorm.DB) {
	ac.DB = d
	ac.DB.LogMode(true)
}

func (ac *CompanyInstallController) List(c *gin.Context) {

	results := []models.CompanyInstall{}
	err := ac.DB.Find(&results).Error

	if err != nil {
		logcompanyinstall.Debugf("Error when looking up company_install List, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "No company_install list found",
		}
		c.JSON(404, res)
		return
	}
	content := gin.H{
		"status":             "200",
		"success":            true,
		"companyinstallList": results,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

// Get a company
func (ac *CompanyInstallController) GetCompanyInstall(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
	entity := models.CompanyInstall{}
	err := ac.DB.Where("id=?", id).Find(&entity).Error

	if err != nil {
		logcompanyinstall.Debugf("Error when looking up company_install, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "company_install not found",
		}
		c.JSON(404, res)
		return
	}

	content := gin.H{
		"status":         "201",
		"success":        true,
		"companyinstall": entity,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/shenxl/mockdata/models"
)

var logcompanySn = logging.MustGetLogger("Company_Sn")

type CompanySnController struct {
	DB *gorm.DB
}

func (ac *CompanySnController) SetDB(d *gorm.DB) {
	ac.DB = d
	ac.DB.LogMode(true)
}

func (ac *CompanySnController) List(c *gin.Context) {

	results := []models.CompanySn{}
	err := ac.DB.Find(&results).Error

	if err != nil {
		logcompanySn.Debugf("Error when looking up companySnList, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "No company_sn list found",
		}
		c.JSON(404, res)
		return
	}
	content := gin.H{
		"status":     "200",
		"success":    true,
		"companysns": results,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

// Get a company
func (ac *CompanySnController) GetCompanySn(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
	entity := models.CompanySn{}
	err := ac.DB.Where("id=?", id).Find(&entity).Error

	if err != nil {
		logcompanySn.Debugf("Error when looking up conpany_sn, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "conpany_sn not found",
		}
		c.JSON(404, res)
		return
	}

	content := gin.H{
		"status":    "201",
		"success":   true,
		"CompanySn": entity,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

// Create a company
func (ac *CompanySnController) Create(c *gin.Context) {
	var entity models.CompanySn
	c.BindJSON(&entity)
	err := ac.DB.Save(&entity).Error
	if err != nil {
		logcompanySn.Debugf("Error while creating a company_sn, the error is '%v'", err)
		res := gin.H{
			"status": "403",
			"error":  "Unable to create company_sn",
		}
		c.JSON(404, res)
		return
	}

	content := gin.H{
		"status":      "201",
		"success":     true,
		"CompanySnId": entity.Id,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)
}

func (ac *CompanySnController) Update(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
	var entity models.CompanySn

	c.BindJSON(&entity)
	err := ac.DB.Table("company_sn").Where("id = ?", id).Updates(&entity).Error
	if err != nil {
		logcompanySn.Debugf("Error while updating a company_sn, the error is '%v'", err)
		res := gin.H{
			"status": "403",
			"error":  "Unable to update company_sn",
		}
		c.JSON(403, res)
		return
	}

	content := gin.H{
		"status":      "201",
		"success":     true,
		"CompanySnID": id,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)
}

func (ac *CompanySnController) Delete(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
	var entity models.CompanySn

	c.BindJSON(&entity)

	// Update Timestamps
	//user.UpdateDate = time.Now().String()

	err := ac.DB.Where("id = ?", id).Delete(&entity).Error
	if err != nil {
		logcompanySn.Debugf("Error while deleting a company_sn, the error is '%v'", err)
		res := gin.H{
			"status": "403",
			"error":  "Unable to delete company_sn",
		}
		c.JSON(403, res)
		return
	}

	content := gin.H{
		"success":     true,
		"CompanySnID": id,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)
}

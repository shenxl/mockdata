package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/shenxl/mockdata/models"
)

var logcompany = logging.MustGetLogger("Companys")

type CompanyController struct {
	DB *gorm.DB
}

func (ac *CompanyController) SetDB(d *gorm.DB) {
	ac.DB = d
	ac.DB.LogMode(true)
}

func (ac *CompanyController) List(c *gin.Context) {

	results := []models.Company{}
	err := ac.DB.Find(&results).Error

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
		"length":   len(results),
		"companys": results,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

func (ac *CompanyController) GetListByIndustry(c *gin.Context) {

	results := []models.Company{}
	name := c.Params.ByName("name")
	err := ac.DB.Where("industry = ?", name).Find(&results).Error
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
		"length":   len(results),
		"companys": results,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

// Get a company
func (ac *CompanyController) GetCompany(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
	company := models.Company{}
	err := ac.DB.Where("id=?", id).Find(&company).Error

	if err != nil {
		logcompany.Debugf("Error when looking up company, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "User not found",
		}
		c.JSON(404, res)
		return
	}

	content := gin.H{
		"status":  "201",
		"success": true,
		"company": company,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

// Create a company
func (ac *CompanyController) Create(c *gin.Context) {
	var company models.Company
	c.BindJSON(&company)

	err := ac.DB.Save(&company).Error
	if err != nil {
		logcompany.Debugf("Error while creating a company, the error is '%v'", err)
		res := gin.H{
			"status": "403",
			"error":  "Unable to create company",
		}
		c.JSON(404, res)
		return
	}

	content := gin.H{
		"status":  "201",
		"success": true,
		// "CompanyId": company.id,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)
}

func (ac *CompanyController) GetIndustry(c *gin.Context) {
	type IndustrysData struct {
		Industry string `json:"industry"`
	}

	results := []IndustrysData{}
	err := ac.DB.Table("company").
		Select("industry as industry").
		Group("industry").
		Scan(&results).Error

	if err != nil {
		logcompany.Debugf("Error when looking up industry, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "No industry found",
		}
		c.JSON(404, res)
		return
	}
	content := gin.H{
		"status":     "200",
		"success":    true,
		"industries": results,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

func (ac *CompanyController) Update(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
	var entity models.Company

	c.BindJSON(&entity)
	err := ac.DB.Table("company").Where("id = ?", id).Updates(&entity).Error
	if err != nil {
		logcompany.Debugf("Error while updating a company, the error is '%v'", err)
		res := gin.H{
			"status": "403",
			"error":  "Unable to update company",
		}
		c.JSON(403, res)
		return
	}

	content := gin.H{
		"status":    "201",
		"success":   true,
		"CompanyID": id,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)
}

func (ac *CompanyController) Delete(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
	var company models.Company

	c.BindJSON(&company)

	// Update Timestamps
	//user.UpdateDate = time.Now().String()

	err := ac.DB.Where("id = ?", id).Delete(&company).Error
	if err != nil {
		logcompany.Debugf("Error while deleting a company, the error is '%v'", err)
		res := gin.H{
			"status": "403",
			"error":  "Unable to delete company",
		}
		c.JSON(403, res)
		return
	}

	content := gin.H{
		"success":   true,
		"CompanyID": id,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)
}

package controllers

import (

  "github.com/op/go-logging"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"

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


func (ac *CompanyController) ListCompanys(c *gin.Context) {

    results := []models.Company{}
    err := ac.DB.Find(&results).Error

    if err != nil {
        logcompany.Debugf("Error when looking up companyList, the error is '%v'", err)
        res := gin.H{
                "status": "404",
                "error": "No company found",
        }
        c.JSON(404, res)
        return
    }
    content := gin.H {
            "status": "200",
            "result": "Success",
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
                "error": "User not found",
        }
        c.JSON(404, res)
        return
    }

    content := gin.H{
            "status": "201",
            "result": "Success",
            "company": company,
        }

    c.Writer.Header().Set("Content-Type", "application/json")
    c.JSON(200, content)
}

// Create a company
func (ac *CompanyController) CreateCompany(c *gin.Context) {
  var company models.Company
  c.BindJSON(&company)

  err := ac.DB.Save(&company).Error
  if err != nil {
    logcompany.Debugf("Error while creating a company, the error is '%v'", err)
    res := gin.H{
      "status": "403",
      "error": "Unable to create company",
    }
    c.JSON(404, res)
    return
  }

  content := gin.H{
    "status": "201",
    "result": "Success",
    // "CompanyId": company.id,
  }

  c.Writer.Header().Set("Content-Type", "application/json")
  c.JSON(201, content)
}


func (ac *CompanyController) UpdateCompany(c *gin.Context) {
  // Grab id
  id := c.Params.ByName("id")
  var company models.Company

  c.BindJSON(&company)
    err := ac.DB.Where("id = ?", id).Updates(&company).Error
    if err != nil {
        logcompany.Debugf("Error while updating a company, the error is '%v'", err)
        res := gin.H{
                "status": "403",
                "error": "Unable to update company",
        }
        c.JSON(403, res)
        return
    }

    content := gin.H{
            "status": "201",
            "result": "Success",
            // "CompanyID": company.id,
        }

    c.Writer.Header().Set("Content-Type", "application/json")
    c.JSON(201, content)
}

func (ac *CompanyController) DeleteCompany(c *gin.Context) {
    // Grab id
  id := c.Params.ByName("id")
  var company models.Company

  c.BindJSON(&company)

    // Update Timestamps
    //user.UpdateDate = time.Now().String()

    err := ac.DB.Where("id = ?", id).Delete(&company).Error
    if err != nil {
        logcompany.Debugf("Error while deleting a user, the error is '%v'", err)
        res := gin.H{
                "status": "403",
                "error": "Unable to delete user",
        }
        c.JSON(403, res)
        return
    }

    content := gin.H {
            "result": "Success",
            // "CompanyID": company.id,
        }

  c.Writer.Header().Set("Content-Type", "application/json")
  c.JSON(201, content)
}

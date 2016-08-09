package companydata

import (
	"net/http"
	"./models"
	"./controllers"
	"github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
)

var (
	Companys *gin.RouterGroup
  Company *gin.RouterGroup
)

func AddRoutes() {
  dc := controllers.DBController{}
	dc.InitDB()
  dc.InitSchema()

	// Get a TodolistController instance
	companyCtr := controllers.CompanyController{}
	companyCtr.SetDB(dc.GetDB())

  Companys.GET("/", companyCtr.ListCompanys)

	Company.GET("/:id", companyCtr.GetCompany)
  Company.PUT("/:id", companyCtr.UpdateCompanyById)
  Company.POST("/", companyCtr.CreateCompany)
	Company.DELETE("/:id", companyCtr.DeleteCompany)
}

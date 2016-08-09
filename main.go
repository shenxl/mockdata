package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shenxl/mockdata/controllers"
)

func main() {

	// Get DBController
	dc := controllers.DBController{}
	dc.InitDB()
  dc.InitSchema()

	// Get a TodolistController instance
	companyCtl := controllers.CompanyController{}
	companyCtl.SetDB(dc.GetDB())

	// Get a todolist resource
	router := gin.Default()

	router.GET("/companys", companyCtl.ListCompanys)
  router.GET("/company/:id", companyCtl.GetCompany)
  router.DELETE("/users/:id", companyCtl.DeleteCompany)
  router.POST("/users", companyCtl.CreateCompany)
  router.PUT("/users/:id", companyCtl.UpdateCompany)

  router.Run(":8888")
}

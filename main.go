package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shenxl/mockdata/controllers"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {

	// Get DBController
	dc := controllers.DBController{}
	dc.InitDB()
	dc.InitSchema()

	// Get a TodolistController instance
	companyCtl := controllers.CompanyController{}
	companyCtl.SetDB(dc.GetDB())

	companySnCtl := controllers.CompanySnController{}
	companySnCtl.SetDB(dc.GetDB())

	serverCtl := controllers.ServerController{}
	serverCtl.SetDB(dc.GetDB())

	companydailyCtl := controllers.CompanyDailyController{}
	companydailyCtl.SetDB(dc.GetDB())

	companymonthCtl := controllers.CompanyMonthlyController{}
	companymonthCtl.SetDB(dc.GetDB())

	companyinstallCtl := controllers.CompanyInstallController{}
	companyinstallCtl.SetDB(dc.GetDB())

	router := gin.New()
	router.Use(CORSMiddleware())
	// Get a todolist resource
	industry := router.Group("/industry")
	{
		industry.GET("/", companyCtl.GetIndustry)
	}

	company := router.Group("/companys")
	{
		company.GET("/", companyCtl.List)
		company.POST("/", companyCtl.Create)
		company.GET("/:id", companyCtl.GetCompany)
		company.DELETE("/:id", companyCtl.Delete)
		company.PUT("/:id", companyCtl.Update)
	}

	companysn := router.Group("/company_sns")
	{
		companysn.GET("/", companySnCtl.List)
		companysn.POST("/", companySnCtl.Create)
		companysn.GET("/:id", companySnCtl.GetCompanySn)
		companysn.DELETE("/:id", companySnCtl.Delete)
		companysn.PUT("/:id", companySnCtl.Update)
	}

	server := router.Group("/servers")
	{
		server.GET("/", serverCtl.List)
		server.POST("/", serverCtl.Create)
		server.GET("/:id", serverCtl.GetServer)
		server.DELETE("/:id", serverCtl.Delete)
		server.PUT("/:id", serverCtl.Update)
	}

	companydaily := router.Group("/companydailys")
	{
		companydaily.GET("/", companydailyCtl.List)
		companydaily.GET("/:id", companydailyCtl.GetCompanyDaily)
	}

	companymonthly := router.Group("/companymonthlys")
	{
		companymonthly.GET("/", companymonthCtl.List)
		companymonthly.GET("/:id", companymonthCtl.GetCompanyMonthly)
	}

	companyinstall := router.Group("/companyinstalls")
	{
		companyinstall.GET("/", companyinstallCtl.List)
		companyinstall.GET("/:id", companyinstallCtl.GetCompanyInstall)
	}

	router.Run(":8888")
}

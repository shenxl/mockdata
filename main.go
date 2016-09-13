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

	dataCtl := controllers.DataController{}
	dataCtl.SetDB(dc.GetDB())

	chartCtl := controllers.ChartController{}
	chartCtl.SetDB(dc.GetDB())

	router := gin.New()
	router.Use(CORSMiddleware())

	report := router.Group("/api/report")
	{
		report.GET("/mapdata/", chartCtl.MapData)
		report.GET("/piedata/installbyindustry", chartCtl.InstallForIndustry)
		report.GET("/piedata/buybyindustry", chartCtl.BuyForIndustry)
		report.GET("/summarydata", chartCtl.SummaryForMonth)
		report.GET("/linedata/", chartCtl.InstallForLine)
		report.GET("/linedata/:id", chartCtl.InstallForLineByID)
	}

	data := router.Group("/api/companys")
	{
		data.GET("/", dataCtl.CompanyListByQuery)
		data.GET("/daily/:id", dataCtl.CompanyDaily)
		data.GET("/types/", companyCtl.GetType)
		// groups.GET("/group_:gid/companies", companymonthCtl.GroupByQuery)
		// groups.GET("/group_:gid/companies/company_:cid/sns", companymonthCtl.GroupByQuery)
	}

	// Get a todolist resource
	industry := router.Group("/api/industry")
	{
		industry.GET("/", companyCtl.GetIndustry)
	}

	company := router.Group("/companys")
	{
		company.GET("/", companyCtl.List)
		//company.GET("/types", companyCtl.GetType)
		company.GET("/industries/:name", companyCtl.GetListByIndustry)
		company.POST("/", companyCtl.Create)
		company.GET("/company/:id", companyCtl.GetCompany)
		company.DELETE("/company/:id", companyCtl.Delete)
		company.PUT("/company/:id", companyCtl.Update)
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

	companymonthly := router.Group("/monthlydata")
	{
		companymonthly.GET("/", companymonthCtl.List)
		companymonthly.GET("/all/:year/:month", companymonthCtl.AllByHere)
		// companymonthly.GET("/:id", companymonthCtl.GetCompanyMonthly)
	}

	companyinstall := router.Group("/companyinstalls")
	{
		companyinstall.GET("/", companyinstallCtl.List)
		companyinstall.GET("/:id", companyinstallCtl.GetCompanyInstall)
	}

	router.StaticFile("/font/iconfont.woff", "./iconfont/iconfont.woff")
	router.StaticFile("/font/iconfont.ttf", "./iconfont/iconfont.ttf")

	router.Run(":8888")
}

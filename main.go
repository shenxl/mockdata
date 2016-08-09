package main

import (
  "./companydata/restapi"
  "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	companydata.Companys = r.Group("/api/companys")
  companydata.Company = r.Group("/api/company")
	companydata.AddRoutes()

	r.Run(":8888")
}

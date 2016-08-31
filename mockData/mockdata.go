package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shenxl/mockdata/controllers"
	"github.com/shenxl/mockdata/models"
)

func main() {
	mocklogfile, err := os.OpenFile("/mock.log", os.O_RDWR|os.O_CREATE, 0)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer mocklogfile.Close()
	logger := log.New(mocklogfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("数据库模块初始化..")
	dc := controllers.DBController{}
	dc.InitDB()
	dc.InitSchema()
	db := dc.GetDB()
	db.LogMode(false)
	logger.Println("数据库模块初始化成功")

	// for index := 0; index < 2000; index++ {
	// 	monthly := models.CompanyMonthly{}
	// 	monthly.CompanyId = int64(rand.Intn(12587) + 12909)
	// 	monthly.ServerId = int64(1)
	// 	monthly.Month = rand.Intn(11) + 1
	// 	monthly.Year = 2015
	// 	monthly.ActivityMax = int64(rand.Intn(1000) + 200)
	// 	monthly.ActivitySum = int64(rand.Intn(3000) + 1000)
	// 	monthly.ActivityAvg = int64(rand.Intn(500) + 300)
	// 	monthly.InstallMax = int64(rand.Intn(1000) + 200)
	// 	monthly.InstallSum = int64(rand.Intn(3000) + 1000)
	// 	monthly.InstallAvg = int64(rand.Intn(500) + 300)
	// 	saveerr := db.Save(&monthly).Error
	// 	if saveerr != nil {
	// 		logger.Printf("添加groupid[%v]失败： %v \n", monthly.Id, saveerr)
	// 	}
	// }

	for index := 0; index < 5000; index++ {
		daily := models.CompanyDaily{}
		daily.CompanyId = int64(rand.Intn(12587) + 12909)
		daily.ServerId = int64(1)
		daily.Month = rand.Intn(11) + 1
		daily.Year = 2015
		daily.Day = rand.Intn(30) + 1
		daily.ActivitySum = int64(rand.Intn(1000) + 500)
		daily.InstallSum = int64(rand.Intn(500) + 200)
		saveerr := db.Save(&daily).Error
		if saveerr != nil {
			logger.Printf("添加groupid[%v]失败： %v \n", daily.Id, saveerr)
		}
	}

	// for index := 0; index < 3000; index++ {
	// 	install := models.CompanyInstall{}
	// 	install.CompanyId = int64(rand.Intn(12587) + 12909)
	// 	install.ServerId = int64(1)
	// 	install.Sum = int64(rand.Intn(6000) + 2000)
	// 	saveerr := db.Save(&install).Error
	// 	if saveerr != nil {
	// 		logger.Printf("添加groupid[%v]失败： %v \n", install.Id, saveerr)
	// 	}
	// }

}

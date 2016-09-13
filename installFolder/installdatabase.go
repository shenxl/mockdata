package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shenxl/mockdata/controllers"
	"github.com/shenxl/mockdata/models"
)

func main() {

	logfile, err := os.OpenFile("/install.log", os.O_RDWR|os.O_CREATE, 0)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("数据库模块初始化..")
	dc := controllers.DBController{}
	dc.InitDB()
	dc.InitSchema()
	db := dc.GetDB()
	db.LogMode(false)
	logger.Println("数据库模块初始化成功")

	type sourceData struct {
		Sales              string
		Sn                 string
		SnType             string
		SnArea             string
		SnNumber           int64
		GroupName          string
		FinalName          string
		CompanyName        string
		Important          int
		Region             string
		Province           string
		City               string
		Address            string
		CompanyType        string
		CompanyIndustry    string
		ServiceLength      int
		ServiceDate        time.Time
		AuthorizationYears int
		AuthorizationDate  time.Time
		Notes              string
	}

	// datas := []sourceData{}
	// db.Limit(100).Find(&datas)
	//
	// for _, data := range datas {
	// 	fmt.Printf("%v\n", data)
	// }
	logger.Println("开始生成数据库..")
	fmt.Println("开始生成数据库..")
	rows, err := db.Model(&sourceData{}).Rows()
	if err == nil {
		for rows.Next() {
			var (
				data      sourceData
				group     models.Group
				company   models.Company
				companySn models.CompanySn
				order     models.Order
			)
			// rows.Scan(&data)
			db.ScanRows(rows, &data)
			// fmt.Printf("获取data [%v]\n", data)
			//  1、获得 Company_group 并在 group表里查询，如果找到则返回 GroupID 若未找到则添加 Group 并返回 GroupID
			if db.First(&group, "group_name = ?", data.GroupName).RecordNotFound() {
				group.GroupName = data.GroupName
				group.Region = data.Region
				group.Province = data.Province
				group.City = data.City
				group.Address = data.Address
				group.Type = data.CompanyType
				group.Industry = data.CompanyIndustry
				group.Important = data.Important
				gsaveE := db.Create(&group).Error
				if gsaveE != nil {
					logger.Printf("添加groupid[%v]失败： %v \n", group.Id, gsaveE)
				}
			}

			//  2、若未找到 则 根据 GroupID 创建 company 表,并返回 companyID
			if db.First(&company, "name = ?", data.CompanyName).RecordNotFound() {
				company.Name = data.CompanyName
				company.GroupId = group.Id
				company.Region = data.Region
				company.Type = data.CompanyType
				company.Industry = data.CompanyIndustry
				company.Province = data.Province
				company.City = data.City
				company.Address = data.Address
				company.Important = data.Important
				csaveE := db.Create(&company).Error
				if csaveE != nil {
					logger.Printf("添加company[%v]失败： %v \n", company.Id, csaveE)
				}
			}

			//	3.根据 company_id 新建订单记录
			order.CompanyId = company.Id
			order.GroupId = group.Id
			order.Sales = data.Sales
			order.Sns = data.Sn
			order.OrderType = data.SnType
			order.OrderArea = data.SnArea
			order.OrderName = data.FinalName
			order.OrderNumber = data.SnNumber
			order.AuthorizationDate = data.AuthorizationDate
			order.AuthorizationYears = data.AuthorizationYears
			order.LengthOfService = data.ServiceLength
			order.ServiceDate = data.ServiceDate
			osaveE := db.Create(&order).Error
			if osaveE != nil {
				logger.Printf("添加order[%v]失败： %v \n", company.Id, osaveE)
			}

			//  4、对 sn 分列 ,将分列后的每一条插入 company_sn 表
			sns := strings.Split(data.Sn, ",")
			for _, sn := range sns {
				companySn = models.CompanySn{}
				if sn != "SERIALCODE_YJ" {
					if db.Where("company_id = ? AND sn = ?", company.Id, sn).Find(&companySn).RecordNotFound() {
						fmt.Printf("向：公司%v 加入序列号[%v]\n", company.Name, sn)
						companySn.CompanyId = company.Id
						companySn.Sn = sn
						snsaveE := db.Create(&companySn).Error
						if snsaveE != nil {
							logger.Printf("添加company_sn[%v]失败： %v\n", companySn.Id, snsaveE)
						}
					}
				}
			}
		}
		logger.Println("生成数据库结束..")
		fmt.Println("生成数据库结束..")
	}
}

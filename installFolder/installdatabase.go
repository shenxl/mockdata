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
		Sales        string `gorm:"sales"`
		Sn           string `gorm:"sn"`
		SnNumber     int    `gorm:"sales"`
		CompanyGroup string `gorm:"sales"`
		Company      string `gorm:"company"`
		Region       string `gorm:"sales"`
		Province     string `gorm:"sales"`
		City         string `gorm:"sales"`
		Address      string `gorm:"sales"`
		Industry     string `gorm:"sales"`
		ServiceYear  int
		ServiceDate  time.Time
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
			)
			// rows.Scan(&data)
			db.ScanRows(rows, &data)
			//  1、获得 Company_group 并在 group表里查询，如果找到则返回 GroupID 若未找到则添加 Group 购买量 并返回 GroupID
			if db.First(&group, "group_name = ?", data.CompanyGroup).RecordNotFound() {
				group.GroupName = data.CompanyGroup
				group.Industry = data.Industry
				group.PurchaseNumber = data.SnNumber
				group.Sales = data.Sales
				gsaveE := db.Save(&group).Error
				if gsaveE != nil {
					logger.Printf("添加groupid[%v]失败： %v \n", group.Id, gsaveE)
				}

			} else {
				// logger.Printf("找到重复集团[%v]：\n", group.GroupName)
				gupdateE := db.Model(&group).Where("id = ?", group.Id).Update("purchase_number", group.PurchaseNumber+data.SnNumber).Error
				if gupdateE != nil {
					logger.Printf("更新groupid[%v]失败： %v \n", group.Id, gupdateE)
				}
			}

			//  3、若未找到 则 根据 GroupID 创建 company 表,并返回 companyID
			if db.First(&company, "name = ?", data.Company).RecordNotFound() {
				company.Name = data.Company
				company.BuyNumber = data.SnNumber
				company.GroupId = group.Id
				company.Region = data.Region
				company.Industry = data.Industry
				company.Province = data.Province
				company.City = data.City
				company.Address = data.Address
				company.LengthOfService = data.ServiceYear
				company.ServiceDate = data.ServiceDate
				csaveE := db.Save(&company).Error
				if csaveE != nil {
					logger.Printf("添加company[%v]失败： %v \n", company.Id, csaveE)
				}
			} else {
				// logger.Printf("找到重复公司[%v]：\n", company.Name)
				//  2、获得 company 并在 company 表里查询 ， 如果找到则返回 company ID 并更新 SnNumber = old + now ， 服务截止日期取离当前时间最远的值，服务器与服务截止日期在同一个实体上。并返回 companyId
				if company.ServiceDate.Before(data.ServiceDate) {
					cupdateE1 := db.Model(&company).Where("id = ?", company.Id).
						Update(models.Company{BuyNumber: company.BuyNumber + data.SnNumber, LengthOfService: data.ServiceYear, ServiceDate: data.ServiceDate}).Error
					if cupdateE1 != nil {
						logger.Printf("更新company[%v]失败： %v \n", company.Id, cupdateE1)
					}
				} else {
					cupdateE2 := db.Model(&company).Where("id = ?", company.Id).
						Update(models.Company{BuyNumber: company.BuyNumber + data.SnNumber, LengthOfService: company.LengthOfService, ServiceDate: company.ServiceDate}).Error
					if cupdateE2 != nil {
						logger.Printf("更新company[%v]失败： %v \n", company.Id, cupdateE2)
					}
				}
			}
			//  4、对 sn 分列 ,将分列后的每一条插入 company_sn 表
			sns := strings.Split(data.Sn, ",")
			// logger.Printf("需要向：公司%v 加入[%v]个序列号\n", company.Name, len(sns))
			for _, sn := range sns {
				companySn.CompanyId = company.Id
				companySn.Sn = sn
				snsaveE := db.Save(&companySn).Error
				if snsaveE != nil {
					logger.Printf("添加company_sn[%v]失败： %v\n", companySn.Id, snsaveE)
				}
			}
		}
		logger.Println("生成数据库结束..")
		fmt.Println("生成数据库结束..")
	}
}

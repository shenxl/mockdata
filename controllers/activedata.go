package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	logging "github.com/op/go-logging"
	"github.com/shenxl/mockdata/models"
)

const (
	FITLERALL  = "all"
	FITLERNONE = ""
)

type ActiveDataController struct {
	DB *gorm.DB
}

func (ac *ActiveDataController) SetDB(d *gorm.DB) {
	ac.DB = d
	ac.DB.LogMode(true)
}

var avitvedataLog = logging.MustGetLogger("DataControl")

func (ac *ActiveDataController) GetActiveDataByCompanys(c *gin.Context) {
	results := []models.VCompanyMonthly{}
	linq := ac.DB.Model(&models.VCompanyMonthly{})
	count := 0
	// 根据关键字查询
	keyword := c.DefaultQuery("keyword", FITLERNONE)
	if keyword != FITLERNONE {
		linq = linq.Where("name LIKE ?", "%"+keyword+"%")
	}

	// 根据企业类型查询
	companyType := c.DefaultQuery("type", FITLERALL)
	typeArr := strings.Split(companyType, ",")
	if companyType != FITLERALL {
		linq = linq.Where("type in (?)", typeArr)
	}
	// 根据重要性查询
	importantStr := c.DefaultQuery("important", FITLERNONE)
	if importantStr != FITLERNONE {
		linq = linq.Where("important = 1")
	}
	// 根据区域查询
	region := c.DefaultQuery("region", FITLERALL)
	regionArr := strings.Split(region, ",")
	if region != FITLERALL {
		linq = linq.Where("region in (?)", regionArr)
	}

	// 根据月份查询
	now := time.Now()
	year := now.Year()
	mon := now.Month()
	Datestr := fmt.Sprintf("%d-%d", year, mon)
	startDate := c.DefaultQuery("startDate", Datestr)
	startYear := strings.Split(startDate, "-")[0]
	startMonth := strings.Split(startDate, "-")[1]
	linq = linq.Where("year = ? AND month = ? ", startYear, startMonth)

	// 排序逻辑
	sorterStr := c.DefaultQuery("field", FITLERNONE)
	orderStr := c.DefaultQuery("order", FITLERNONE)

	if orderStr == "descend" {
		sorterStr += " DESC"
	}

	if sorterStr != FITLERNONE {
		linq = linq.Order(sorterStr)
	}

	// 计算记录总数
	linq = linq.Count(&count)

	// 分页逻辑
	sikp := c.DefaultQuery("start", "0")
	limit := c.DefaultQuery("limit", "5")

	linq = linq.Limit(limit).Offset(sikp)

	err := linq.Find(&results).Error
	if err != nil {
		avitvedataLog.Debugf("获取企业月报活视图失败，错误代码为 '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "获取企业月报活视图失败",
		}
		c.JSON(404, res)
		return
	}
	content := gin.H{
		"status":  "200",
		"success": true,
		"length":  count,
		"data":    results,
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

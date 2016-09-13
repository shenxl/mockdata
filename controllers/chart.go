package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

var chartLog = logging.MustGetLogger("ChartControl")

type ChartController struct {
	DB *gorm.DB
}

func (ac *ChartController) SetDB(d *gorm.DB) {
	ac.DB = d
	ac.DB.LogMode(true)
}

func (ac *ChartController) InstallForLine(c *gin.Context) {
	type resultData struct {
		Day         string `json:"day"`
		ActivitySum int64  `json:"active_sum"`
		InstallSum  int64  `json:"install_sum"`
	}

	now := time.Now()
	year := now.Year()
	mon := now.Month()
	datestr := fmt.Sprintf("%d,%d", year, mon)
	date := c.DefaultQuery("Date", datestr)

	dateYear := strings.Split(date, ",")[0]
	dataMonth := strings.Split(date, ",")[1]

	results := []resultData{}
	rawSQL := `
	select day, sum(activity_sum) as activity_sum , sum(install_sum) as install_sum  from company_daily
	where year = ` + dateYear + ` and month = ` + dataMonth + `
	GROUP BY day
	`
	err := ac.DB.Raw(rawSQL).Scan(&results).Error
	if err != nil {
		dataLog.Debugf("Error when looking up groupDaily List, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "No groud list found",
		}
		c.JSON(404, res)
		return
	}

	content := gin.H{
		"status":  "200",
		"success": true,
		"data":    results,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)

}

func (ac *ChartController) InstallForLineByID(c *gin.Context) {
	type resultData struct {
		Year        int64 `json:"year"`
		Month       int64 `json:"month"`
		Day         int64 `json:"day"`
		ActivitySum int64 `json:"activity_sum"`
		InstallSum  int64 `json:"install_sum"`
	}
	id := c.Params.ByName("id")
	results := []resultData{}
	rawSQL := `
    select year,month,day,sum(activity_sum) as activity_sum , sum(install_sum) as install_sum from company_daily
    where company_id = '` + id + `'
    group by year,month,day,company_id
	`
	err := ac.DB.Raw(rawSQL).Scan(&results).Error
	if err != nil {
		dataLog.Debugf("Error when looking up groupDaily List, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "No groud list found",
		}
		c.JSON(404, res)
		return
	}

	content := gin.H{
		"status":  "200",
		"success": true,
		"data":    results,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)

}

func (ac *ChartController) InstallForIndustry(c *gin.Context) {
	type ResultData struct {
		Type string `json:"name"`
		Sum  int64  `json:"value"`
	}

	results := []ResultData{}
	rawSQL := `
		select company.type,sum(company_install.sum) as sum from company_install
		LEFT JOIN company on company.id = company_install.company_id
		group by company.type
	`
	err := ac.DB.Raw(rawSQL).Scan(&results).Error
	if err != nil {
		dataLog.Debugf("Error when looking up groupDaily List, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "No groud list found",
		}
		c.JSON(404, res)
		return
	}

	content := gin.H{
		"status":  "200",
		"success": true,
		"data":    results,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)

}

func (ac *ChartController) SummaryForMonth(c *gin.Context) {
	type nowData struct {
		InstallSum  int64 `json:"install_sum"`
		ActivitySum int64 `json:"activity_sum"`
	}

	now := time.Now()
	year := now.Year()
	mon := now.Month()
	datestr := fmt.Sprintf("%d,%d", year, mon)
	date := c.DefaultQuery("Date", datestr)

	dateYear := strings.Split(date, ",")[0]
	dataMonth := strings.Split(date, ",")[1]
	nowdata := []nowData{}
	predata := []nowData{}
	rawSQL := `
		select Sum(install_sum) as install_sum ,Sum(activity_sum) as activity_sum from company_monthly
		WHERE year = ` + dateYear + ` and month = ` + dataMonth + `
	`
	predateYear, _ := strconv.Atoi(dateYear)
	preMonth, _ := strconv.Atoi(dataMonth)
	preMonth = preMonth - 1
	if preMonth == 0 {
		preMonth = 12
		predateYear = predateYear - 1
	}
	prerawSQL := `
		select Sum(install_sum) as install_sum ,Sum(activity_sum) as activity_sum from company_monthly
		WHERE year = ` + strconv.Itoa(predateYear) + ` and month = ` + strconv.Itoa(preMonth) + `
	`

	err := ac.DB.Raw(rawSQL).Scan(&nowdata).Error
	err2 := ac.DB.Raw(prerawSQL).Scan(&predata).Error
	if err != nil && err2 != nil {
		dataLog.Debugf("Error when looking up groupDaily List, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "No groud list found",
		}
		c.JSON(404, res)
		return
	}

	content := gin.H{
		"status":  "200",
		"success": true,
		"nowdata": nowdata,
		"predata": predata,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)

}

func (ac *ChartController) BuyForIndustry(c *gin.Context) {
	type ResultData struct {
		Type           string `json:"name"`
		PurchaseNumber int64  `json:"value"`
	}

	results := []ResultData{}
	rawSQL := `
		select type,sum(total) as  purchase_number from v_group_order
		GROUP BY type
	`
	err := ac.DB.Raw(rawSQL).Scan(&results).Error
	if err != nil {
		dataLog.Debugf("Error when looking up groupDaily List, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "No groud list found",
		}
		c.JSON(404, res)
		return
	}

	content := gin.H{
		"status":  "200",
		"success": true,
		"data":    results,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)

}

func (ac *ChartController) MapData(c *gin.Context) {

	type MapResult struct {
		Province string `json:"name"`
		Total    string `json:"value"`
	}
	results := []MapResult{}

	err := ac.DB.Table("v_company_order").Select("province , sum(total) as total").Group("province").
		Order("total").Scan(&results).Error
	if err != nil {
		logcompanySn.Debugf("按地域查找购买量发生错误 '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "No company data found",
		}
		c.JSON(404, res)
		return
	}
	content := gin.H{
		"status":  "200",
		"success": true,
		"count":   len(results),
		"mapdata": results,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

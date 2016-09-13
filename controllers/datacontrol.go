package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

var dataLog = logging.MustGetLogger("DataControl")

type DataController struct {
	DB *gorm.DB
}

func (ac *DataController) SetDB(d *gorm.DB) {
	ac.DB = d
	ac.DB.LogMode(true)
}

func (ac *DataController) CompanyDaily(c *gin.Context) {
	type ResultData struct {
		GroupId     string `json:"group_id"`
		Year        int    `json:"year"`
		Month       int    `json:"month"`
		Day         int    `json:"day"`
		ActivitySum int64  `json:"activity_sum"`
	}
	groupId := c.Params.ByName("id")

	results := []ResultData{}
	rawSQL := `
		select T1.group_id,T1.year,T1.month,T1.day,Sum(T1.activity_sum) as activity_sum
		FROM
			(SELECT company_daily.activity_sum,company_daily.year,company_daily.month,company_daily.day,company.group_id
			FROM company_daily
			LEFT JOIN company ON company.id = company_daily.company_id ) as T1
	  WHERE T1.group_id = ` + groupId + `
		GROUP BY T1.group_id,t1.T1.year,T1.month,T1.day
		ORDER BY T1.year,T1.month,T1.day
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
		"status":    "200",
		"success":   true,
		"dailyData": results,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)

}

func (ac *DataController) CompanyListByQuery(c *gin.Context) {

	type ResultData struct {
		Name         string  `json:"name"`
		CompanyId    string  `json:"company_id"`
		GroupId      string  `json:"group_id"`
		Type         string  `json:"type"`
		Industry     string  `json:"industry"`
		Important    int     `json:"important"`
		Total        string  `json:"buy_number"`
		ServerID     string  `json:"server_id"`
		Year         int     `json:"year"`
		Month        int     `json:"month"`
		CompanyNum   string  `json:"company_num"`
		ActivitySum  int64   `json:"activity_sum"`
		ActivityAvg  int64   `json:"activity_avg"`
		ActivityMax  int64   `json:"activity_max"`
		InstallSum   int64   `json:"install_sum"`
		InstallAvg   int64   `json:"install_avg"`
		InstallMax   int64   `json:"install_max"`
		InstallTotal int64   `json:"install_total"`
		UserRate     float64 `json:"user_rate"`
		InstallRate  float64 `json:"install_rate"`
	}

	type PageOption struct {
		Current  string `json:"current"`
		Total    string `json:"total"`
		PageSize string `json:"pageSize"`
	}

	type FilterOption struct {
		Industry string `json:"industry"`
	}

	now := time.Now()
	year := now.Year()
	mon := now.Month()
	Datestr := fmt.Sprintf("%d,%d", year, mon)
	startDate := c.DefaultQuery("startDate", Datestr)
	// endDate := c.DefaultQuery("endDate", Datestr)

	start := c.DefaultQuery("start", "0")
	limit := c.DefaultQuery("limit", "5")

	companyType := c.DefaultQuery("type", "all")

	// industry := c.DefaultQuery("industry", "all")
	// TODO: 根据排序
	// orderBy := c.DefaultQuery("orderBy", "")
	// orderByDesc := c.DefaultQuery("orderBy", "")
	// sortbyDesc := c.DefaultQuery("sortbyDesc", "[]")

	// 根据类型检索
	typeArr := strings.Split(companyType, ",")
	typeStr := ""
	if typeArr[0] == "all" {
		typeStr = ""
	} else if len(typeArr) > 0 {
		typeStr = "WHERE type in ("
		for index := 0; index < len(typeArr); index++ {
			if index == len(typeArr)-1 {
				typeStr = typeStr + `"` + typeArr[index] + `")`
			} else {
				typeStr = typeStr + `"` + typeArr[index] + `" ,`
			}
		}
	}

	// 根据行业检索
	// industryArr := strings.Split(industry, ",")
	// dataLog.Debugf("获得的参数为industryArr: %v ", industryArr)
	// industryStr := ""
	// if industryArr[0] == "all" {
	// 	industryStr = ""
	// } else if len(industryArr) > 0 {
	// 	industryStr = "WHERE industry in ("
	// 	for index := 0; index < len(industryArr); index++ {
	// 		if index == len(industryArr)-1 {
	// 			industryStr = industryStr + `"` + industryArr[index] + `")`
	// 		} else {
	// 			industryStr = industryStr + `"` + industryArr[index] + `" ,`
	// 		}
	// 	}
	// }

	// 根据关键字检索
	keywordStr := c.DefaultQuery("keyword", "")
	if keywordStr != "" && typeStr != "" {
		keywordStr = "AND name like " + `"%` + keywordStr + `%"`
	} else if keywordStr != "" && typeStr == "" {
		keywordStr = "Where name like " + `"%` + keywordStr + `%"`
	}

	importantStr := c.DefaultQuery("important", "")
	if importantStr != "" && typeStr != "" {
		importantStr = "AND important = 1"
	} else if importantStr != "" && typeStr == "" {
		importantStr = "Where important = 1"
	}

	sorterStr := c.DefaultQuery("field", "")
	orderStr := c.DefaultQuery("order", "")
	if sorterStr != "" {
		sorterStr = "Order by " + sorterStr
		if orderStr == "descend" {
			sorterStr += " DESC"
		}
	} else if sorterStr == "" && orderStr == "" {
		sorterStr = "Order by total DESC "
	}

	startYear := strings.Split(startDate, ",")[0]
	startMonth := strings.Split(startDate, ",")[1]
	// endYear := strings.Split(endDate, ",")[0]
	// endMonth := strings.Split(endDate, ",")[1]

	results := []ResultData{}
	rawSQL := `
	from (select monthly.* ,company_install.sum as install_total
		from (
			SELECT *
			from company_monthly
			where year = ` + startYear + ` and month = ` + startMonth + ` )as monthly
		left join company_install on company_install.company_id = monthly.company_id ) as totalinfo
	left join v_company_order as company on company.id = totalinfo.company_id
	`

	err := ac.DB.Raw(`
		select company.name,company.group_id,company.type,company.industry,company.total, company.important,totalinfo.*,
			totalinfo.install_total/company.total as install_rate , totalinfo.activity_sum /company.total as user_rate
		` + rawSQL + `
		` + typeStr + `
		` + keywordStr + `
		` + importantStr + `
		` + sorterStr + `
		LIMIT ` + limit + ` OFFSET ` + start + `
		`).Scan(&results).Error

	if err != nil {
		dataLog.Debugf("获得企业报活数据列表发现错误： '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "获得企业报活数据列表发现错误",
		}
		c.JSON(404, res)
		return
	}

	pageOption := PageOption{}

	ac.DB.Raw(`
		SELECT count(1) as total
		` + rawSQL + `
		` + typeStr + `
		` + keywordStr + `
		` + importantStr + `
		`).Scan(&pageOption)

	pageOption.PageSize = limit

	content := gin.H{
		"status":      "200",
		"success":     true,
		"companylist": results,
		"byPage":      pageOption,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

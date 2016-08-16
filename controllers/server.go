package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/shenxl/mockdata/models"
)

var logserver = logging.MustGetLogger("Server")

type ServerController struct {
	DB *gorm.DB
}

func (ac *ServerController) SetDB(d *gorm.DB) {
	ac.DB = d
	ac.DB.LogMode(true)
}

func (ac *ServerController) List(c *gin.Context) {

	type ServerData struct {
		models.Server
		models.Company
	}

	results := []ServerData{}

	err := ac.DB.Table("server").
		Select("server.* , company.*").
		Joins("left join company on server.company = company.id ").
		Scan(&results).Error

	if err != nil {
		logserver.Debugf("Error when looking up servers, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "No server list found",
		}
		c.JSON(404, res)
		return
	}
	content := gin.H{
		"status":  "200",
		"success": true,
		"servers": results,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

// Get a company
func (ac *ServerController) GetServer(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
	entity := models.Server{}
	err := ac.DB.Where("id=?", id).Find(&entity).Error

	if err != nil {
		logserver.Debugf("Error when looking up server, the error is '%v'", err)
		res := gin.H{
			"status": "404",
			"error":  "server not found",
		}
		c.JSON(404, res)
		return
	}

	content := gin.H{
		"status":  "201",
		"success": true,
		"server":  entity,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

// Create a company
func (ac *ServerController) Create(c *gin.Context) {
	var entity models.Server
	c.BindJSON(&entity)
	err := ac.DB.Save(&entity).Error

	if err != nil {
		logserver.Debugf("Error while creating a server, the error is '%v'", err)
		res := gin.H{
			"status": "403",
			"error":  "Unable to create server",
		}
		c.JSON(404, res)
		return
	}

	content := gin.H{
		"status":   "201",
		"success":  true,
		"serverid": entity.Id,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)
}

func (ac *ServerController) Update(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
	var entity models.Server

	c.BindJSON(&entity)
	err := ac.DB.Table("server").Where("id = ?", id).Updates(&entity).Error
	if err != nil {
		logserver.Debugf("Error while updating a server, the error is '%v'", err)
		res := gin.H{
			"status": "403",
			"error":  "Unable to update server",
		}
		c.JSON(403, res)
		return
	}

	content := gin.H{
		"status":   "201",
		"success":  true,
		"serverid": id,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)
}

func (ac *ServerController) Delete(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
	var entity models.Server

	c.BindJSON(&entity)

	// Update Timestamps
	//user.UpdateDate = time.Now().String()

	err := ac.DB.Where("id = ?", id).Delete(&entity).Error
	if err != nil {
		logserver.Debugf("Error while deleting a server, the error is '%v'", err)
		res := gin.H{
			"status": "403",
			"error":  "Unable to delete server",
		}
		c.JSON(403, res)
		return
	}

	content := gin.H{
		"success":  true,
		"serverid": id,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)
}

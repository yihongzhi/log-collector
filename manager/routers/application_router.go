package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yihongzhi/log-collector/manager/database"
	"github.com/yihongzhi/log-collector/manager/models"
	"net/http"
)

//查询应用列表
func ApplicationList(c *gin.Context) {
	var apps []models.Application
	database.DB.Find(&apps)
	c.JSON(http.StatusOK, apps)
}

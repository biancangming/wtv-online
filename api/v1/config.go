package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wtv-online/model"
)

func IsRequiredLogin(c *gin.Context) {
	config := model.ConfigData
	if config.Username == "" && config.Password == "" {
		c.String(http.StatusOK, "0")
	} else {
		c.String(http.StatusOK, "1")
	}
}

func GetConfigData(c *gin.Context) {
	config := model.ConfigData
	c.JSON(http.StatusOK, gin.H{
		"title":       config.Title,
		"qrcodeLink":  config.QrcodeLink,
		"description": config.Description,
	})
}

func Login(c *gin.Context) {
	config := model.Config{}
	_ = c.ShouldBindJSON(&config)
	if config.Username == model.ConfigData.Username && config.Password == model.ConfigData.Password {
		c.String(http.StatusOK, "1")
	} else {
		c.String(http.StatusOK, "0")
	}
}

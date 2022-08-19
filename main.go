package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"wtv-online/model"
	"wtv-online/router"
)

func main() {
	//编译linux gox -os="linux"
	//编译windows gox -os="windows"
	//编译darwin gox -os="darwin"
	gin.SetMode(os.Getenv(gin.ReleaseMode))
	model.InitModel()
	router.InitRouter()
}

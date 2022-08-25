package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"wtv-online/model"
	"wtv-online/router"
)

func main() {
	//编译 gox -os="linux windows darwin" -output="build"
	gin.SetMode(os.Getenv(gin.ReleaseMode))
	model.InitModel()
	router.InitRouter()
}

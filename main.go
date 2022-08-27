package main

import (
	"fmt"
	"github.com/go-ini/ini"
)

import (
	"github.com/gin-gonic/gin"
	"os"
	"wtv-online/model"
	"wtv-online/router"
)

func InitAppIni() {
	Cfg, err := ini.Load("app.ini")
	if err != nil {
		fmt.Println("未设置配置文件，跳过此步骤")
		return
	}
	userSection := Cfg.Section("user")
	configSection := Cfg.Section("config")
	username := userSection.Key("username").MustString("")
	password := userSection.Key("password").MustString("")

	title := configSection.Key("title").MustString("超级文本链接分享工具 - 公众号 一个橙子pro")
	qrcodeLink := configSection.Key("qrcode_link").MustString("http://online.bianbingdang.com/da428b3b6d6e058eb738e6b77a08e9b.jpg")
	description := configSection.Key("description").MustString("一个橙子出品, 上边是我的微信公众号，本网站支持私有部署，https://github.com/biancangming/wtv-online 。本网站是本人纯手工制作，请勿无脑攻击，请勿使用本站发布皇都读、政治敏感、道德沦丧等信息，本人一经发现立即删除。")

	model.ConfigData = model.Config{
		Username: username,
		Password: password,
		Title: title,
		QrcodeLink: qrcodeLink,
		Description: description,
	}
	fmt.Println(model.ConfigData)
}

func main() {
	//编译 gox -os="linux windows darwin" -output="build"
	gin.SetMode(os.Getenv(gin.ReleaseMode))
	InitAppIni()
	model.InitModel()
	router.InitRouter()
}

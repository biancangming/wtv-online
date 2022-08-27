package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	v1 "wtv-online/api/v1"
)

func getAssetContent(name string) string {
	b, _ := Asset(name)
	return string(b)
}

func InitRouter() {

	r := gin.New()
	fmt.Println()

	r.NoRoute(func(c *gin.Context) {
		rp := c.Request.URL.Path

		if rp == "/" {
			c.Writer.Header().Add("Content-Type", "text/html; charset=utf-8")
			c.String(http.StatusOK, getAssetContent("ui/dist/index.html"))
			return
		}

		if strings.HasSuffix(rp, ".js") {
			c.Writer.Header().Add("Content-Type", "application/javascript")
		} else if strings.HasSuffix(rp, ".css") {
			c.Writer.Header().Add("Content-Type", "text/css")
		}

		c.String(http.StatusOK, getAssetContent("ui/dist"+rp))
	})

	r.Static("/static", "./static")

	{
		v1g := r.Group("/api")

		// 链接分享
		{
			g := v1g.Group("/share")
			g.POST("update", v1.UpdateOrAddShare)
			g.POST("updateStatus", v1.UpdateUseStatus)
			g.GET("/:uuid", v1.GetShareUrl)
			g.GET("/get", v1.GetShareUrlData)
			g.GET("/urls", v1.GetShareUrls)
		}

		// 配置
		{
			g := v1g.Group("/config")
			g.GET("isRequiredLogin", v1.IsRequiredLogin)
			g.GET("config", v1.GetConfigData)
			g.POST("login", v1.Login)
		}

	}

	r.Run(":1999")
}

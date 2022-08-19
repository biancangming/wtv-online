package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay/pkg/util"
	"net/http"
	"strings"
	"wtv-online/model"
	"wtv-online/utils/result"
)

// GetShareUrl 获取单条信息内容
func GetShareUrl(c *gin.Context) {
	up := c.Param("uuid")

	var r = strings.Split(up, ".")
	uuid := r[0]
	end := r[1]
	s, err := model.GetShareUrl(uuid)

	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	if s.Uuid == "" {
		c.String(http.StatusOK, "链接不存在，欢迎关注微信公众号 一个橙子pro")
		return
	}

	if s.UseStatus == 1 {
		c.String(http.StatusOK, "链接已经被原作者废弃，如果您是原作者，请还原废弃状态再使用，欢迎关注微信公众号 一个橙子pro")
		return
	}

	if s.FileType != end {
		c.String(http.StatusOK, "链接不存在，检查后缀是否正确，欢迎关注微信公众号 一个橙子pro")
		return
	}

	c.String(http.StatusOK, s.Content)
}

func GetShareUrlData(c *gin.Context) {
	u := c.Query("uuid")
	s, err := model.GetShareUrl(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail("联系管理员"))
		return
	}
	c.JSON(http.StatusOK, result.Ok(s, "Ok"))
}

func GetShareUrls(c *gin.Context) {
	us := util.String2Int(c.Query("useStatus"))
	ss, err := model.GetShareUrls(us)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail("Ok"))
		return
	}

	c.JSON(http.StatusOK, result.Ok(ss, "Ok"))
}

func UpdateOrAddShare(c *gin.Context) {
	share := model.Share{}
	_ = c.ShouldBindJSON(&share)

	r, err := model.UpdateOrAddShare(share)

	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail("联系管理员"))
		return
	}
	var msg string
	if share.Uuid == "" {
		msg = "添加成功"
	} else {
		msg = "链接更新内容成功"
	}
	c.JSON(http.StatusOK, result.Ok(r, msg))
}

func UpdateUseStatus(c *gin.Context) {
	share := model.Share{}

	_ = c.ShouldBindJSON(&share)

	r, err := model.UpdateUseStatus(share)

	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail("联系管理员"))
		return
	}

	c.JSON(http.StatusOK, result.Ok(r, "成功"))
}

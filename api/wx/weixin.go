package wx

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"wtv-online/utils/util"
)

// WeiXinTextMsg 微信文本消息结构体
type WeiXinTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	MsgId        int64
	XMLName      xml.Name `xml:"xml"`
}

// WeiXinReceive 包含id name 的下拉框
func WeiXinReceive(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	ok := util.CheckSignature(signature, timestamp, nonce, "yigechengzi")
	if !ok {
		log.Println("微信公众号接入校验失败!")
		c.String(http.StatusInternalServerError, "")
		return
	}

	log.Println("微信公众号接入校验成功!")
	c.String(http.StatusOK, echostr)
}

// WeiXinMsgReceive 微信消息接收
func WeiXinMsgReceive(c *gin.Context) {
	var textMsg WeiXinTextMsg
	err := c.ShouldBindXML(&textMsg)
	if err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		return
	}

	log.Printf("[消息接收] - 收到消息, 消息类型为: %s, 消息内容为: %s\n", textMsg.MsgType, textMsg.Content)

	WXMsgReply(c, textMsg.ToUserName, textMsg.FromUserName)
}

// WXMsgReply 微信消息回复
func WXMsgReply(c *gin.Context, fromUser, toUser string) {
	repTextMsg := WeiXinTextMsg{
		ToUserName:   toUser,
		FromUserName: fromUser,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      fmt.Sprintf("[消息回复] - %s", time.Now().Format("2006-01-02 15:04:05")),
	}

	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	fmt.Println(string(msg))
	//c.XML(http.StatusOK, msg)
	_, _ = c.Writer.Write(msg)
}

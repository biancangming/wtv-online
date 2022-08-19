package alipay

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	"net/http"
	"os"
	"strconv"
	"time"
	"wtv-online/model"
	"wtv-online/utils/result"
)

var signType string
var payUrl string
var appid string
var notifyUrl string
var rsaPrivateKey string
var alipayPublicKey string
var client alipay.Client

func InitAliPayInfo() {
	/*** 请填写以下配置信息 ***/
	signType = os.Getenv("ALI_PAY_SIGN_TYPE") //签名算法类型，支持RSA2和RSA，推荐使用RSA2
	payUrl = os.Getenv("ALI_PAY_URL")
	appid = os.Getenv("ALI_PAY_APPID")          //  沙箱 2016110300789385 https://open.alipay.com 账户中心->密钥管理->开放平台密钥，填写添加了当面付应用的APPID
	notifyUrl = os.Getenv("ALI_PAY_NOTIFY_URL") //付款成功后的异步回调地址 	//付款金额，单位:元 	//订单标题
	//商户私钥，填写对应签名算法类型的私钥，如何生成密钥参考：https://docs.open.alipay.com/291/105971和https://docs.open.alipay.com/200/105310
	rsaPrivateKey = os.Getenv("ALI_PAY_RAS_PRIVATE_KEY")
	//支付宝公钥，登录支付宝开放平台，账户中心->密钥管理->开放平台密钥，找到对应的应用，在接口内容加密方式处查看支付宝公钥
	alipayPublicKey = os.Getenv("ALI_PAY_RAS_PUBLIC_KEY")

	_client, err := alipay.NewClient(appid, rsaPrivateKey, false)
	client = *_client
	if err != nil {
		xlog.Error(err)
		return
	}
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetNotifyUrl(notifyUrl)
}

//生成订单号
func order() string {
	now := time.Now()
	return fmt.Sprintf("%s%08x%05x", "", now.Unix(), now.UnixNano()%0x100000)
}

func DoPay(c *gin.Context) {
	var payAmount float64
	var orderName string

	ctx := context.Background()
	if c.Query("totalFee") != "" {
		payAmount, _ = strconv.ParseFloat(c.Query("totalFee"), 64)
	}
	outTradeNo := order()
	if c.Query("outTradeNo") != "" {
		outTradeNo = c.Query("outTradeNo")
	}
	if c.Query("orderName") != "" {
		orderName = c.Query("orderName")
	}

	bm := make(gopay.BodyMap)
	bm.Set("subject", orderName)
	bm.Set("out_trade_no", outTradeNo)
	bm.Set("total_amount", payAmount)
	//创建订单
	aliRsp, err := client.TradePrecreate(ctx, bm)
	if err != nil {
		xlog.Error("err:", err)
		c.JSON(http.StatusInternalServerError, result.Fail(err.Error()))
		return
	}

	_ = model.AddFacePay(model.FacePay{OutTradeNo: outTradeNo, TradeName: orderName, OutTradeMoney: payAmount})
	xlog.Debug("aliRsp:", *aliRsp)
	xlog.Debug("aliRsp.QrCode:", aliRsp.Response.QrCode)
	xlog.Debug("aliRsp.OutTradeNo:", aliRsp.Response.OutTradeNo)

	c.JSON(http.StatusOK, result.Ok(gin.H{
		"qrCode":     aliRsp.Response.QrCode,
		"outTradeNo": aliRsp.Response.OutTradeNo,
	}, "生成订单成功"))
}

// Notify 异步回调通知处理
func Notify(c *gin.Context) {

	notifyReq, err := alipay.ParseNotifyToBodyMap(c.Request) // c.Request 是 gin 框架的写法
	if err != nil {
		xlog.Error(err)
	}

	fmt.Println(notifyReq)

	// 公钥模式验签
	//    注意：APP支付，手机网站支付，电脑网站支付 不支持同步返回验签
	//    aliPayPublicKey：支付宝平台获取的支付宝公钥
	//    signData：待验签参数，aliRsp.SignData
	//    sign：待验签sign，aliRsp.Sign
	_, err2 := alipay.VerifySign(alipayPublicKey, notifyReq)

	if err2 != nil {
		xlog.Error(err)
	}
	outTradeNo := notifyReq.Get("out_trade_no")
	outTradeMoney, _ := strconv.ParseFloat(notifyReq.Get("out_trade_money"), 64)
	appId := notifyReq.Get("app_id")
	authAppId := notifyReq.Get("auth_app_id")
	buyerId := notifyReq.Get("buyer_id")
	buyerLogonId := notifyReq.Get("buyer_logon_id")
	buyerPayAmount, _ := strconv.ParseFloat(notifyReq.Get("buyer_pay_amount"), 64)
	gmtCreate, _ := time.Parse("2006-01-02 15:04:05", notifyReq.Get("gmt_create"))
	gmtPayment, _ := time.Parse("2006-01-02 15:04:05", notifyReq.Get("gmt_payment"))
	sellerEmail := notifyReq.Get("seller_email")
	tradeNo := notifyReq.Get("trade_no")
	tradeStatus := notifyReq.Get("trade_status")
	//tradeName := notifyReq.Get("trade_name")

	_ = model.UpdateFacePay(outTradeNo, model.FacePay{
		OutTradeMoney:  outTradeMoney,
		AppId:          appId,
		AuthAppId:      authAppId,
		BuyerId:        buyerId,
		BuyerLogonId:   buyerLogonId,
		BuyerPayAmount: buyerPayAmount,
		GmtCreate:      gmtCreate,
		GmtPayment:     gmtPayment,
		SellerEmail:    sellerEmail,
		TradeNo:        tradeNo,
		TradeStatus:    tradeStatus,
	})
	//程序执行完后必须打印输出“success”（不包含引号）。如果商户反馈给支付宝的字符不是success这7个字符，支付宝服务器会不断重发通知，直到超过24小时22分钟。一般情况下，25小时以内完成8次通知（通知的间隔频率一般是：4m,10m,10m,1h,2h,6h,15h）；
	c.String(http.StatusOK, "success")
}

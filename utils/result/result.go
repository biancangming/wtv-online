package result

import (
	"github.com/gin-gonic/gin"
	cons "wtv-online/utils"
)

func Ok(data interface{}, msg string) gin.H {
	return gin.H{
		"code": cons.SuccessCode,
		"data": data,
		"msg":  msg,
	}
}

func Fail(msg string) gin.H {
	return gin.H{
		"code": cons.FailCode,
		"data": nil,
		"msg":  msg,
	}
}

func PageOk(records interface{}, total int64, current, size int, msg string) gin.H {
	return gin.H{
		"code": cons.SuccessCode,
		"data": gin.H{
			"total":   total,
			"records": records,
			"current": current,
			"size":    size,
		},
		"msg": msg,
	}
}

func PageFail(msg string) gin.H {
	return gin.H{
		"code": cons.FailCode,
		"data": gin.H{
			"total":   0,
			"records": nil,
			"current": 0,
			"size":    0,
		},
		"msg": msg,
	}
}

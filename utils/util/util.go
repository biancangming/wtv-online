package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type CitySN struct {
	Cip   string
	Cid   string
	Cname string
}

// GetCitySn 获取IP 地区等信息
func GetCitySn() (CitySN, error) {
	city := new(CitySN)

	res, err := http.Get("https://pv.sohu.com/cityjson?ie=utf-8")
	defer func() {
		if err := res.Body.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	if err != nil {
		return *city, err
	}

	r, ioErr := ioutil.ReadAll(res.Body)
	if ioErr != nil {
		return *city, err
	}

	reg := regexp.MustCompile(`var returnCitySN = (.*?);`)
	citySnStr := reg.FindStringSubmatch(string(r))[1]

	fmt.Println(citySnStr)

	cityErr := json.Unmarshal([]byte(reg.FindStringSubmatch(string(r))[1]), city)
	if cityErr != nil {
		return *city, err
	}

	return *city, nil
}

func GetPage(c *gin.Context) (int, int, error) {
	var err error
	current, errC := strconv.Atoi(c.DefaultQuery("current", "1"))
	err = errC

	size, errS := strconv.Atoi(c.DefaultQuery("size", "10"))
	err = errS
	return current, size, err
}

//Randomstring 取得随机字符串:使用字符串拼接
func Randomstring(length int) string {
	if length < 1 {
		return ""
	}
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))

	var rchar string = ""
	for i := 1; i <= length; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}
	return rchar
}

// CheckSignature 微信公众号签名检查
func CheckSignature(signature, timestamp, nonce, token string) bool {
	arr := []string{timestamp, nonce, token}
	// 字典序排序
	sort.Strings(arr)

	n := len(timestamp) + len(nonce) + len(token)
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(arr); i++ {
		b.WriteString(arr[i])
	}

	return Sha1(b.String()) == signature
}

// Sha1 进行Sha1编码
func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5 进行Md5编码
func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

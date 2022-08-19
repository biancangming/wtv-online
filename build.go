package main

import (
	"fmt"
	"net/url"
)

//filepath: 要编译的文件的路径
func build() {
	u, _ := url.Parse("http://12156.vod.adultiptv.net/ph55de7a932b9c7/play.m3u8")
	fmt.Println(u.Host)
}

//
//func main() {
//	build() //要编译的文件的路径
//}

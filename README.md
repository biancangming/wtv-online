# 简介

使用本程序可在线托管txt、m3u、m3u8、yml、json等文本，可用于iptv电视连接配置，geojson地图数据保存，私人程序固定数据配置等。

# 部署教程

https://mp.weixin.qq.com/s/pK-8EbJ_4s3C8FsFgXSb3w

# 演示网站


http://online.bianbingdang.com/

本站只用作演示、请勿用于个人真实生产，数据可能随时被删除

# 打包教程
安装 
```shell
go get github.com/jteeuwen/go-bindata/...
go get github.com/elazarl/go-bindata-assetfs/...
```

执行 `go-bindata-assetfs .\ui\dist` 生成 `bindata.go` 文件，并复制到router目录下

打包命令，全局需要安装`gox`

- 编译linux gox -os="linux"
- 编译windows gox -os="windows"
- 编译darwin gox -os="darwin"

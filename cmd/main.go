package main

import (
	"duryun-blog/api/v1"
	"duryun-blog/model"
)

func main() {
	// 引用数据库
	model.InitDb()
	// 引入路由组件
	v1.InitRouter()

}

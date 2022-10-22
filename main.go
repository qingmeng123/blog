package main

import (
	"duryun-blog/model"
	"duryun-blog/routes"
)

func main() {
	// 引用数据库
	model.InitDb()
	// 引入路由组件
	routes.InitRouter()

}

// @Author:冯铁城 [17615007230@163.com] 2023-10-12 11:22:20
package main

import (
	"log"

	"golang.org/x/sync/errgroup"
	"net-project-1/base"
	"net-project-1/routers"
)

var (
	g errgroup.Group
)

func main() {

	//1.初始化数据库
	base.InitDB()

	//2.初始化路由
	routers.InitRoutersV1()

	//3.运行服务
	g.Go(func() error {
		return routers.Router1.Run(":8080")
	})

	//4.监听异常
	if err := g.Wait(); err != nil {
		log.Fatalf("server run err:[%v]\n", err)
	}
}

package main

import (
	"./common"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"net/http"
)

//统一验证拦截器，每个接口都需要提前验证
func Auth(respW http.ResponseWriter, req *http.Request) error {
	return errors.New("认证失败")
}

//执行正常业务逻辑
func Check(respW http.ResponseWriter, req *http.Request) {
	fmt.Println("执行check!")
}

func main() {
	//1. 过滤器
	filter := common.NewFilter()
	//注册拦截器
	filter.RegisterUriFilter("/check", Auth)
	//2,启动服务
	http.HandleFunc("/check", filter.Handler(Check))
	http.ListenAndServe(":8083", nil)
}

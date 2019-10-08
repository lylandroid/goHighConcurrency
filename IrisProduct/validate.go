package main

import (
	"./common"
	"./encrypt"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"net/http"
)

//统一验证拦截器，每个接口都需要提前验证
func Auth(respW http.ResponseWriter, req *http.Request) error {
	return CheckUserInfo(req)
}

func CheckUserInfo(req *http.Request) error {
	uidCookie, err := req.Cookie("uid")
	if err != nil {
		return errors.New("用户未登录！")
	}
	signCookie, err := req.Cookie("sign")
	if err != nil {
		return errors.New("用户加密串获取失败！")
	}
	deSignByte, err := encrypt.DePwdCode(signCookie.Value)
	if err != nil {
		return errors.New("加密串已被篡改！")
	}
	fmt.Println("结果比对：uid=", uidCookie.Value, " sign: ", string(deSignByte))
	if CheckIdInfo(uidCookie.Value, string(deSignByte)) {
		return nil
	}
	return errors.New("身份校验失败!")
}

//自定义逻辑判断
func CheckIdInfo(checkStr string, signStr string) bool {
	if checkStr == signStr {
		return true
	}
	return false
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

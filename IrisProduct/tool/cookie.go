package tool

import (
	"github.com/kataras/iris"
	"net/http"
)

//设置全局cookie
func GlobalCookie(ctx iris.Context,name string,value string)  {
	ctx.SetCookie(&http.Cookie{Name:name,Value:value,Path:"/"})
}

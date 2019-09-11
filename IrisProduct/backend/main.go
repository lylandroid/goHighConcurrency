package main

import "github.com/kataras/iris"

const rootWebPath = "./IrisProduct/backend/web/"

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	template := iris.HTML(rootWebPath+"views", ".html").
		Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	//设置模板目录
	app.StaticWeb("/assets", rootWebPath+"assets")
	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message",
			ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	//注册控制权
	//TODO
	//启动服务
	app.Run(iris.Addr(":8080"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations, )

}

package main

import (
	"../services"
	"../common"
	"./web/controller"
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"time"
	"./middlerware"
)

const rootWebPath = "./IrisProduct/frontend/web/"

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	template := iris.HTML(rootWebPath+"views", ".html").
		Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	//设置模板目录
	app.StaticWeb("/public", rootWebPath+"public")
	//访问生成html静态文件
	app.StaticWeb("/html", rootWebPath+"htmlProductOut")
	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message",
			ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	session := sessions.New(sessions.Config{
		Cookie:  "helloword",
		Expires: 60 * time.Minute,
	})

	//注册控制权
	userService := services.NewUserService("user")
	userParty := app.Party("/user")
	user := mvc.New(userParty)
	user.Register(userService, ctx, session.Start)
	user.Handle(new(controller.UserController))

	db, err := common.NewMysqlConn()
	if err != nil {
		panic(err)
	}
	orderService := services.NewOrderService("order", db)

	productService := services.NewProductServiceImp(db)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	productParty.Use(middlerware.AuthProduct)
	product.Register(ctx, orderService, productService, session.Start)
	product.Handle(new(controller.ProductController))

	//启动服务
	app.Run(iris.Addr(":8082"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations, )

}

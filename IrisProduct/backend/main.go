package main

import (
	"context"
	"github.com/kataras/iris"
	"../common"
	"../services"
	"github.com/kataras/iris/mvc"
	"./web/contorllers"
)

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
	//连接数据库成功
	db, err := common.NewMysqlConn()
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	productServiceImp := services.NewProductServiceImp("product", db)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productServiceImp)
	product.Handle(new(contorllers.ProductController))

	//注册控制权

	//启动服务
	app.Run(iris.Addr(":8080"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations, )

}

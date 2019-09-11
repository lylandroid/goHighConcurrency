package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"./web/controllers"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./Iris/web/views", ".html"))
	mvc.New(app.Party("/movie")).Handle(new(controllers.MovieController))
	app.Run(iris.Addr(":8080") /*, iris.WithoutServerError(iris.ErrServerClosed)*/)
}

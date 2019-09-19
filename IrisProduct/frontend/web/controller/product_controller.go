package controller

import (
	"github.com/kataras/iris"
	"../../../services"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.IProductService
	Session        *sessions.Session
}

func (p *ProductController) GetDetail() mvc.View {
	product, err := p.ProductService.GetProduct(1)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/view.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

package controllers

import (
	"../../../services"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"../../../datamodels"
	"../../../common"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.IProductService
}

func (p *ProductController) GetAll() mvc.View {
	products, err := p.ProductService.GetProductAll()
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	fmt.Println(errMsg)
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			//"errMsg":     errMsg,
			"productArray": products,
		},
	}
}

func (p *ProductController) Get() mvc.View {
	products, err := p.ProductService.GetProductAll()
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	fmt.Println(errMsg)
	return mvc.View{
		Name: "product/product.html",
		Data: iris.Map{
			//"errMsg":     errMsg,
			"productArray": products,
		},
	}
}

//修改商品
func (p *ProductController) PostUpdate() {
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "imooc"})
	if err := dec.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	err := p.ProductService.UpdateProduct(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

func (p *ProductController) GetAdd() mvc.View {
	return mvc.View{
		Name: "product/add.html",
	}
}

func (p *ProductController) PostAdd() {
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "imooc"})
	if err := dec.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	_, err := p.ProductService.InsertProduct(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

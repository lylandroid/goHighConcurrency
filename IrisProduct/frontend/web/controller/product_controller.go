package controller

import (
	"../../../datamodels"
	"../../../services"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"strconv"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.IProductService
	OrderService   services.IOrderService
	Session        *sessions.Session
}

func (p *ProductController) GetDetail() mvc.View {
	product, err := p.ProductService.GetProduct(1)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	fmt.Println(product)
	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/view.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

func (p *ProductController) GetOrder() mvc.View {
	productId, pIdErr := p.Ctx.URLParamInt64("productID")
	userID, uIdErr := strconv.Atoi(p.Ctx.GetCookie("uid"))
	if pIdErr != nil || uIdErr != nil {
		p.Ctx.Application().Logger().Debug(pIdErr, uIdErr)
	}
	product, err := p.ProductService.GetProduct(productId)
	fmt.Println("product: ", product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	var orderId int64
	showMessage := "抢购失败！"
	if product.ProductNum > 0 {
		var err error
		product.ProductNum -= 1
		err = p.ProductService.UpdateProduct(product)
		if err != nil {
			p.Ctx.Application().Logger().Debug(err)
		}
		order := &datamodels.Order{
			UserId:      int64(userID),
			ProductId:   productId,
			OrderStatus: datamodels.OrderSuccess,
		}
		fmt.Println("order: ", order)
		orderId, err = p.OrderService.InsertOrder(order)
		if err != nil {
			p.Ctx.Application().Logger().Debug(err)
		} else {
			showMessage = "抢购成功！"
		}
	}

	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/result.html",
		Data: iris.Map{
			"orderID":     orderId,
			"showMessage": showMessage,
		},
	}

}

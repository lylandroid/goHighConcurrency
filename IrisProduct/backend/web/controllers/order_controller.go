package controllers

import (
	"github.com/kataras/iris"
	"../../../services"
	"github.com/kataras/iris/mvc"
)

type OrderController struct {
	Ctx          iris.Context
	OrderService services.IOrderService
}

func (o *OrderController) Get() mvc.View {
	orderMap, err := o.OrderService.GetAllOrderInfo()
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	return mvc.View{
		Name: "order/view.html",
		Data: iris.Map{
			"order": orderMap,
		},
	}

}

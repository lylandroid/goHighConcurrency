package controller

import (
	"../../../datamodels"
	"../../../services"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.IProductService
	OrderService   services.IOrderService
	Session        *sessions.Session
}

func (p *ProductController) GetGenerateHtml() {
	//1.获取模板
	template, err := template.ParseFiles(filepath.Join(templatePath,"product.html"))
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	//2. 获取html生成路径
	fileName := filepath.Join(htmlOutPath, "htmlProduct.html")
	//3.获取模板渲染数据
	productId, err := p.Ctx.URLParamInt64("productID")
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	product, err := p.ProductService.GetProduct(productId)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	//4.生成静态文件
	generateStaticHtml(p.Ctx, template, fileName, product)
}

var (
	htmlOutPath  = "./IrisProduct/frontend/web/generate/htmlProductOut/" //生成Html保存目录
	templatePath = "./IrisProduct/frontend/web/views/template/"          //静态文件模板目录
)

func generateStaticHtml(ctx iris.Context, template *template.Template,
	fileName string, product *datamodels.Product) {
	//判断文件是否存在
	if exist(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			ctx.Application().Logger().Debug(err)
		}
	}
	//2,生成静态文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		ctx.Application().Logger().Debug(err)
	}
	defer file.Close()
	template.Execute(file, &product)
}

//判断文件是否存在
func exist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
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

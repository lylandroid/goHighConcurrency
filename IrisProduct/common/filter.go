package common

import (
	"net/http"
)

//声明一个新的数据类型（函数类型）
type FilterHandler func(respW http.ResponseWriter, req *http.Request) error

//拦截器结构体
type Filter struct {
	//用来存储需要拦截的URI
	filterUriMap map[string]FilterHandler
}

//拦截器初始化
func NewFilter() *Filter {
	return &Filter{
		filterUriMap: make(map[string]FilterHandler),
	}
}

//注册拦截器
func (f *Filter) RegisterUriFilter(uri string, filterHandler FilterHandler) {
	f.filterUriMap[uri] = filterHandler
}

//获取拦截器对应的Handler
func (f *Filter) GetFilter(uri string) FilterHandler {
	return f.filterUriMap[uri]
}

type WebHandler func(respW http.ResponseWriter, req *http.Request)

//执行拦截器，返回函数类型
func (f *Filter) Handler(webHandler WebHandler) func(http.ResponseWriter, *http.Request) {
	return func(respW http.ResponseWriter, req *http.Request) {
		filterHandler := f.filterUriMap[req.RequestURI]
		if filterHandler != nil {
			if err := filterHandler(respW, req); err != nil {
				respW.Write([]byte(err.Error()))
				return
			}
		}
		//执行正常注册函数
		webHandler(respW, req)
		/*for uri, filterHandler := range f.filterUriMap {
			if uri == req.RequestURI {
				err := filterHandler(respW, req)
				if err != nil {
					respW.Write([]byte(err.Error()))
					return
				}
				break
			}
		}*/
	}

}

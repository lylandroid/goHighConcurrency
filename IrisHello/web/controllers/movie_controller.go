package controllers

import (
	"github.com/kataras/iris/mvc"
	"../../repositories"
	"../../services"
)

type MovieController struct {
}

func (m *MovieController) Get() mvc.View {
	movieRepository := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManager(movieRepository)
	movieResult := movieService.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		//Name: "../view/movie/index.html",
		Data: movieResult,
	}
}

package services

import (
	"../repositories"
	"fmt"
)

type MovieService interface {
	ShowMovieName() string
}
type MovieServiceManager struct {
	Repo repositories.MovieRepository
}

func NewMovieServiceManager(repo repositories.MovieRepository) *MovieServiceManager {
	return &MovieServiceManager{
		Repo: repo,
	}
}

func (m *MovieServiceManager) ShowMovieName() string {
	name := "获取到models数据：" + m.Repo.GetMovieName()
	fmt.Println(name)
	return name
}

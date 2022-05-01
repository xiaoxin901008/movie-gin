package routers

import (
	"github.com/gin-gonic/gin"
	"movie/apps/actress"
	"movie/apps/film"
	"movie/apps/genre"
	"movie/apps/image"
	"movie/apps/link"
	"movie/apps/login"
	"movie/apps/register"
)

type Option func(engine *gin.Engine)

var options = make([]Option, 0)

func Include(opts ...Option) {
	options = append(options, opts...)
}

func IncludeRouter() {
	Include(film.Routers)
	Include(actress.Routers)
	Include(link.Routers)
	Include(image.Routers)
	Include(genre.Routers)
	Include(login.Routers)
	Include(register.Routers)
}


func Init(r *gin.Engine) *gin.Engine {
	IncludeRouter()
	for _, opt := range options {
		opt(r)
	}
	return r
}

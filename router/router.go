package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name            string
	Method          string
	Path            string
	HandlerFunction func(*gin.Context)
}
type routes struct {
	router *gin.Engine
}
type Routes []Route

func (r routes) User(g *gin.RouterGroup) {
	users := g.Group("/user")
	for _, singleRoute := range user {
		switch singleRoute.Method {
		case http.MethodGet:
			users.GET(singleRoute.Path, singleRoute.HandlerFunction)
		case http.MethodPost:
			users.POST(singleRoute.Path, singleRoute.HandlerFunction)
		case http.MethodPut:
			users.PUT(singleRoute.Path, singleRoute.HandlerFunction)
		case http.MethodPatch:
			users.PATCH(singleRoute.Path, singleRoute.HandlerFunction)
		case http.MethodDelete:
			users.DELETE(singleRoute.Path, singleRoute.HandlerFunction)
		}
	}
}

func Routing() {
	r := routes{
		router: gin.Default(),
	}
	grouping := r.router.Group(os.Getenv("API_VERSION"))
	r.User(grouping)
	r.router.Run(":" + os.Getenv("PORT"))
}

package router

import (
	"goStructure/constant"
	usercontroller "goStructure/controller/userController"
	"net/http"
)

var user = Routes{
	Route{"Hello World", http.MethodGet, constant.HelloWorldPath, usercontroller.HelloWorld},
}

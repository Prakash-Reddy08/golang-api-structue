package usercontroller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	fmt.Println("Hello World")
}

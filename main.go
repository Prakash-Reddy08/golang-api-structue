package main

import (
	"goStructure/initializers"
	"goStructure/router"
)

func init() {
	initializers.LoadSecrets()
}
func main() {
	router.Routing()
}

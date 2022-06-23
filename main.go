package main

import (
	"os"
	"specialteam/config"
	"specialteam/module/special"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

var (
	IP_ADDR string
	PORT    string
)

func init() {
	//init app env
	godotenv.Load(".env")
	IP_ADDR = os.Getenv("IP_ADDR")
	PORT = os.Getenv("PORT")
	//ini database
	config.ConnectDB()
}

func main() {

	//routing
	route := echo.New()

	// v1 := route.Group("/v1")

	route.GET("special/search_specialteam", special.Search)
	route.POST("special/create_specialteam", special.Create)
	route.PUT("special/update_specialteam/:id", special.Update)
	route.DELETE("special/delete_specialteam/:id", special.Delete)

	route.Start(IP_ADDR + ":" + PORT)
}

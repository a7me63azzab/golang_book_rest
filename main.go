package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	
	//Middleware

	// Log what happens in your app to keep track of everything from request / response data && status code && Request processing time, etc ... 
    e.Use(middleware.Logger())

	// Catch UnExpected Errors and return a 500 server error.
	e.Use(middleware.Recover())
	
	// DB
	
	// Routes

	// Server
	e.Logger.Fatal(e.Start(":8080"))
}

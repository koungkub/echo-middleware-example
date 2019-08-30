package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.RequestID())

	e.Use(middleware.CORS())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			fmt.Println("middleware")
			// return &echo.HTTPError{
			// 	Code:    200,
			// 	Message: "eiei",
			// }
			return next(c)
		})
	})

	// https://forum.labstack.com/t/how-to-use-standard-middleware-func-http-handler-http-handler/272/2

	// net/http (middleware)
	e.Use(echo.WrapMiddleware(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("middleware2")
			h.ServeHTTP(w, r) // pass this middleware
		})
	}))

	e.GET("/", func(c echo.Context) error {
		fmt.Println("routing")
		return c.JSON(200, c.Response().Header().Get(echo.HeaderXRequestID))
	})
	
	e.GET("/a", echo.WrapHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	})))

	log.Fatal(e.Start(":1323"))
}

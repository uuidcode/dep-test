package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		var book = new(Book)
		context.Bind(book)

		fmt.Printf("%+v\n", book)
		return context.String(http.StatusOK, "Hello")
	})

	e.Start(":17003")
	fmt.Println("OK")
}

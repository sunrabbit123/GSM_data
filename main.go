package main

import (
	"GSM_data/router"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo"
)

var (
	port  string
	debug bool
	e     *echo.Echo
)

func init() {
	debug = true
	e = router.Router(debug)
}
func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	startMessage := "Go server at http://127.0.0.1" + port
	fmt.Println(startMessage)

	e.Logger.Fatal(e.Start(":" + port))
}

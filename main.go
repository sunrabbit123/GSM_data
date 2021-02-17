package main

import (
	"GSM_data/router"
	"fmt"

	"github.com/labstack/echo"
)

var (
	port  string
	debug bool
	e     *echo.Echo
)

func init() {
	port = ":4995"
	debug = true
	e = router.Router(debug)
}
func main() {
	startMessage := "Go server at http://127.0.0.1" + port
	fmt.Println(startMessage)

	e.Logger.Fatal(e.Start(port))
}

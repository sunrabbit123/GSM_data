package main

import (
	"fmt"

	"github.com/Goolgae/GSM_data/server/router"
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

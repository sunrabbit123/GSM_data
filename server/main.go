package main

import (
	"fmt"

	"github.com/Goolgae/GSM_data/server/router"
)

func main() {
	port := ":9876"
	debug := true

	e := router.Router(debug)

	startMessage := "Go server at http://127.0.0.1" + port
	fmt.Println(startMessage)

	e.Logger.Fatal(e.Start(port))
}

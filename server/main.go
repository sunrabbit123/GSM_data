package main

import (
	"fmt"

	router "github.com/Goolgae/GSM_data/server/router"
)

func main() {
	port := "9876"
	e = router.Router()
	startMessage := "Go server at http://127.0.0.1" + port
	fmt.Println(startMessage)

	e.Logger.Fatal(e.Start(port))
}

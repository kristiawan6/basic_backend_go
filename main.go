package main

import (
	"backend_golang/src/config"
	"backend_golang/src/helper"
	"backend_golang/src/routes"
	"fmt"
	"net/http"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helper.Migration()
	defer config.DB.Close()
	routes.Router()
	fmt.Print("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

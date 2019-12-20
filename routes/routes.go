package routes

import (
	"goapi/controllers"
	"net/http"
)

func GetRoutes() {
	http.HandleFunc("/", controllers.Index)
}

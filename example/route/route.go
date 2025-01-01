package route

import (
	"net/http"

	"github.com/gocroot/example/controller"
)

func URL() {
	http.HandleFunc("/", controller.GetHome)
}
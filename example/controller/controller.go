package controller

import (
	"net/http"

	"github.com/gocroot/reqs"

	"github.com/gocroot/example/model"
)

func GetHome(respw http.ResponseWriter, req *http.Request) {
	var resp model.Response
	resp.Response = reqs.GetIPaddress()
	reqs.WriteJSON(respw, http.StatusOK, resp)
}
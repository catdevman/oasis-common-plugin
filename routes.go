package main

import (
	"github.com/catdevman/oasis/shared"
	"log"
	"net/http"
)

var routes = []shared.Route{
	{
		Method:    "GET",
		Pattern:   "GET /common",
		HandlerID: "getHandler",
	},
	{
		Method:    "POST",
		Pattern:   "POST /common",
		HandlerID: "postHandler",
	},
}

var handlers = map[string]func(shared.SerializedRequest) shared.SerializedResponse{
	"getHandler":  getHandler,
	"postHandler": postHandler,
}
var getHandler = func(req shared.SerializedRequest) shared.SerializedResponse {
	return shared.SerializedResponse{
		Header: http.Header{
			"Content-Type": {
				"application/json",
			},
		},
		StatusCode: http.StatusOK,
		Body:       `{"common": true}`,
	}
}

var postHandler = func(req shared.SerializedRequest) shared.SerializedResponse {
	log.Print("In the post method :)")
	return shared.SerializedResponse{
		Header: http.Header{
			"Content-Type": {
				"application/json",
			},
		},
		StatusCode: http.StatusOK,
		Body:       `{"common": true}`,
	}
}

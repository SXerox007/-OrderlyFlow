package server

import (
	"net/http"

	"github.com/SXerox007/-OrderlyFlow/base/router"
	"github.com/gorilla/handlers"
)

/**
*
* start the server
*
**/
func StartServer(port string) {
	headersOk := handlers.AllowedHeaders([]string{""})
	http.ListenAndServe(":"+port, handlers.CORS(headersOk)(router.HeadNodeRouter))
}

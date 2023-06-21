package main

import (
	"fmt"
	"net/http"

	routes "countdowntimer/router"

	"github.com/gorilla/mux"
)

// @version 1.0

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// schemes https

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", mux.Vars(req)["name"])
}

func main() {

	router := routes.NewRouter()
	router.HandleFunc("/hello/{name}", helloHandler)

	http.ListenAndServe(":80", router)

}

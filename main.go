package main

import (
	"fmt"
	"net/http"

	routes "countdowntimer/router"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// @version 1.0

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// schemes https

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", mux.Vars(req)["name"])
}

func main() {
	viper.SetConfigName("appsettings")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		} else {
			// Config file was found but another error was produced
		}
	}

	router := routes.NewRouter()
	router.HandleFunc("/hello/{name}", helloHandler)

	http.ListenAndServe(":80", router)

}

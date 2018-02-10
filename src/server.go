package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	DefaultFunctionDirectory string = "/openbrisk/"
	DevFunctionDirectory     string = "../examples/"
)

var (
	moduleName         string      // figlet.sh
	moduleDependencies string      // figlet.deps.sh
	functionHandler    string      // not nedded?
	functionTimeout    int    = 10 // NOTE: Define default value.
)

func main() {
	fmt.Println(os.Getenv("MODULE_NAME"))
	fmt.Println(os.Getenv("FUNCTION_HANDLER"))
	fmt.Println(os.Getenv("FUNCTION_DEPENDENCIES"))

	/*if os.Getenv("FUNCTION_TIMEOUT") != "" {
		functionTimeout, err := strconv.Atoi(os.Getenv("FUNCTION_TIMEOUT"))
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println(functionTimeout)
	}*/

	http.HandleFunc("/healthz", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "text/plain")
		response.Header().Set("Connection", "close")
		response.WriteHeader(http.StatusOK)
	})

	fmt.Println("Listening on port 8080 ...")
	var error = http.ListenAndServe(":8080", nil)
	if error != nil {
		panic(error)
	}
}

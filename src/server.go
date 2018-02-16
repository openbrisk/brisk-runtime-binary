package main

import (
	"fmt"
	"net/http"
	//"os"
	"os/exec"
)

const (
	DefaultFunctionDirectory string = "/openbrisk/"
	DevFunctionDirectory     string = "../examples/"
)

var (
	moduleName         string      // $MODULE_NAME.sh
	moduleDependencies string      // $MODULE_NAME.deps.sh
	functionHandler    string      // not nedded?
	functionTimeout    int    = 10 // NOTE: Define default value.
)

func main() {
	// TODO: Read env variables and set for warapper script.
	// TODO: Pipe input into script.
	command := exec.Command("./function-wrapper.sh")
	output, e := command.Output();
	if e == nil {
		fmt.Printf("%s", output);
	}

	http.HandleFunc("/healthz", func(response http.ResponseWriter, request *http.Request) {
		if(request.Method == "GET") {
			response.Header().Set("Content-Type", "text/plain")
			response.Header().Set("Connection", "close")
			response.WriteHeader(http.StatusOK)
		} else {
			response.WriteHeader(http.StatusNotFound)
		}
	})

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		if(request.Method == "GET") {
			// TODO: Handle function invocation without parameters. 
		} else if(request.Method == "POST") {
			// TODO: Handle function invocation with parameters.
		} else {
			response.WriteHeader(http.StatusNotFound)
		}
	})

	fmt.Println("Listening on port 8080 ...")
	/*var error = http.ListenAndServe(":8080", nil)
	if error != nil {
		panic(error)
	}*/
}

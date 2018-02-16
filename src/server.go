package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/openbrisk/brisk-runtime-binary/src/util"
)

var (
	moduleName         string // $MODULE_NAME.sh
	moduleDependencies string // $MODULE_NAME.deps.sh
	functionHandler    string // $FUNCTION_HANDLER
	functionTimeout    = 10   // $FUNCTION_TIMEOUT
	envError           error
)

func main() {
	readEnvironment()

	// Handle GET requests to the health check endpoint.
	http.HandleFunc("/healthz", func(response http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			response.WriteHeader(http.StatusOK)
		} else {
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Handle GET and POST requests to the function endpoint.
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case
			"GET", "POST":
			executeFunction(response, request)
			break
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Println("Listening on port 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// readEnvironment reads the environment variables of the function and logs the results.
func readEnvironment() {
	moduleName = os.Getenv("MODULE_NAME")
	log.Printf("Using module script: %s.sh", moduleName)

	moduleDependencies = os.Getenv("MODULE_NAME")
	log.Printf("Using module deps script: %s.deps.sh", moduleDependencies)

	functionHandler = os.Getenv("FUNCTION_HANDLER")
	log.Printf("Using function handler: %s", functionHandler)

	if os.Getenv("FUNCTION_TIMEOUT") != "" {
		functionTimeout, envError = strconv.Atoi(os.Getenv("FUNCTION_TIMEOUT"))
		if envError != nil {
			log.Fatal(envError)
		}
	}
	log.Printf("Using function timeout: %d seconds", functionTimeout)
}

// executeFunction executes the function script using the http requests body as input.
// The duration header and timeout features are supported.
func executeFunction(response http.ResponseWriter, request *http.Request) {
	start := time.Now()

	log.Println("Running function wrapper script")
	command := exec.Command("./function-wrapper.sh")
	pipeWriter, _ := command.StdinPipe()

	var isHeaderWritten = false
	var functionInput []byte
	var functionOutput []byte
	var err error
	var waitGroup sync.WaitGroup
	var timer *time.Timer

	functionInput = getRequestBody(request)

	waitGroup.Add(2)

	// Create timer that kills the script after the defined timeout.
	if functionTimeout > 0 {
		timer = time.NewTimer(time.Duration(functionTimeout * 1000000000))

		go func() {
			<-timer.C
			log.Println("Function timed out: killing function process")

			if command != nil && command.Process != nil {
				isHeaderWritten = true
				response.WriteHeader(http.StatusRequestTimeout)
				response.Write([]byte("Function timed out: killed function process\n"))

				processKillError := command.Process.Kill()
				if processKillError != nil {
					log.Fatal(processKillError.Error())
				}
			}
		}()
	}

	// Pipe input body to function in go routine to prevent blocking.
	go func() {
		defer waitGroup.Done()
		pipeWriter.Write(functionInput)
		pipeWriter.Close()
	}()

	// Get the output value from the combined stdout and stderr of the function script.
	go func() {
		defer waitGroup.Done()
		functionOutput, err = command.CombinedOutput()

		// Check if the result contains the forward structure.
		functionResult, marshalError := util.UnmarshalFunctionResult(functionOutput)
		if marshalError == nil {
			log.Println("Found function result structure: writing forward header")
			decodedResult, decodingError := base64.StdEncoding.DecodeString(functionResult.Result)
			if decodingError == nil {
				log.Println("Decoding base64 encoded function result")
				functionOutput = decodedResult
			}
			if !isHeaderWritten {
				response.Header().Set("X-OpenBrisk-Forward", functionResult.Forward.To)
			}
		}
	}()

	waitGroup.Wait()
	if timer != nil {
		timer.Stop()
	}

	// Write the duration header in nanoseconds.
	duration := time.Since(start)
	if !isHeaderWritten {
		response.Header().Set("X-OpenBrisk-Duration", fmt.Sprint(duration.Nanoseconds()))
		response.WriteHeader(http.StatusOK)
		response.Write(functionOutput)
		isHeaderWritten = true
	}
	log.Printf("Function execution duration: %s nanoseconds", fmt.Sprint(duration.Nanoseconds()))
}

// getRequestBody returns the request body data (the function input) as byte array.
func getRequestBody(request *http.Request) []byte {
	var body []byte
	var err error

	// No request body: return the empty byte array.
	if request.Body == nil {
		return body
	}

	// Close the body stream at the end of the function.
	defer request.Body.Close()

	// Read request body.
	body, err = ioutil.ReadAll(request.Body)
	// TODO: Better error handling: Send back per http response?
	if err != nil {
		return body
	}

	return body
}

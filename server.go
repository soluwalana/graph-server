package main

import (
    "net/http"
    "os"
    "log"
)

// DefaultPort is the default port to use if once is not specified by the SERVER_PORT environment variable
const DefaultPort = "7893";


func getServerPort() (string) {
    port := os.Getenv("SERVER_PORT");
    if port != "" {
        return port;
    }

    return DefaultPort;
}

// EchoHandler echos back the request as a response
func EchoHandler(writer http.ResponseWriter, request *http.Request) {

    log.Println("Sending Tree on " + request.URL.Path + " to client (" + request.RemoteAddr + ")")

    writer.Header().Set("Access-Control-Allow-Origin", "*")

    // allow pre-flight headers
    writer.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")

    writer.Write([]byte(`{
      "value": 5,
      "left": {
        "value": 3,
        "left": { "value": 2, "left": {"left": {"value": 10}}, "right": {"value": 0}},
        "right": { "value": 13, "right": {"value": 10} }
      },
      "right": {
        "value": 3,
        "left": {
          "value": 4,
          "right": { "value": 13 }
        },
        "right": { "value": 10 }
      }
    }`))
}

func main() {

    log.Println("starting server, listening on port " + getServerPort())

    http.HandleFunc("/", EchoHandler)
    http.ListenAndServe(":" + getServerPort(), nil)
}

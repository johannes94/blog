package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

type Service struct {
	Name, Port   string
	ResponseTime int
}

func (service Service) HandleAllRequest(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(service.ResponseTime) * time.Millisecond)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Response from %v", service.Name)))
}

func main() {

	var serviceName = flag.String("s", "", "Name of this Service")
	var port = flag.String("p", "", "HTTP port to listen to")
	var responseTime = flag.Int("r", 0, "Time in ms to wait before response")
	flag.Parse()

	service := Service{*serviceName, *port, *responseTime}
	http.HandleFunc("/", service.HandleAllRequest)
	go http.ListenAndServe(fmt.Sprintf(":%v", service.Port), nil)

	fmt.Printf("%v listening on port: %v, Press <Enter to exit>\n", service.Name, service.Port)
	fmt.Scanln()
}

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

type PizzaOrder struct {
	Pizza, Store, Price string
}

func main() {

	var pizza = flag.String("pizza", "", "Pizza to order")
	var store = flag.String("store", "", "Name of the Pizza Store")
	var price = flag.String("price", "", "Price")

	flag.Parse()

	order := PizzaOrder{*pizza, *store, *price}
	body, _ := json.Marshal(order)

	start := time.Now()

	orderChan := make(chan *http.Response)
	paymentChan := make(chan *http.Response)
	storeChan := make(chan *http.Response)

	errGrp, _ := errgroup.WithContext(context.Background())

	// OrderService is expected at 8081
	errGrp.Go(func() error { return SendPostAsync("http://localhost:8081", body, orderChan) })

	// PaymentService is expected at 8082
	errGrp.Go(func() error { return SendPostAsync("http://localhost:8082", body, paymentChan) })

	// StoreService is expected at 8083
	errGrp.Go(func() error { return SendPostAsync("http://localhost:8083", body, storeChan) })

	err := errGrp.Wait()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error with submitting the order, try again later...")
		os.Exit(1)
	}

	orderResponse := <-orderChan
	defer orderResponse.Body.Close()
	bytes, _ := ioutil.ReadAll(orderResponse.Body)
	fmt.Println(string(bytes))

	paymentResponse := <-paymentChan
	defer paymentResponse.Body.Close()
	bytes, _ = ioutil.ReadAll(paymentResponse.Body)
	fmt.Println(string(bytes))

	storeResponse := <-storeChan
	defer storeResponse.Body.Close()
	bytes, _ = ioutil.ReadAll(storeResponse.Body)
	fmt.Println(string(bytes))

	end := time.Now()

	fmt.Printf("Order processed after %v seconds\n", end.Sub(start).Seconds())
}

func SendPostAsync(url string, body []byte, rc chan *http.Response) error {
	response, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err == nil {
		go func() { rc <- response }()
	}

	return err
}

func SendPostRequest(url string, body []byte) *http.Response {
	response, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		panic(err)
	}

	return response
}

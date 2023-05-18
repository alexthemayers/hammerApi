package main

import (
	"context"
	"flag"
	"fmt"
	"hammerApi/api"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	debug         bool
	ratePerSecond int
)

type client struct {
	http.Client
}

func main() {
	ratePerSecond = *flag.Int("rate", 20, "number of calls made per second")
	debug = *flag.Bool("debug", false, "show debug output")
	flag.Parse()
	callCount := 0
	notifyOn := 500
	protocol := "http"
	host := "localhost"
	port := ":87"
	//base := fmt.Sprintf("%s://%s:%s", protocol, host, port)
	//client := client{
	//	http.Client{Timeout: 5 * time.Second},
	//}
	client, err := api.NewHTTPClient(&api.Config{
		Port:        port,
		Host:        host,
		Protocol:    protocol,
		Token:       "testToken",
		ContentType: "application/json",
	}, func(r *http.Request) error {
		token := getToken()
		header := "Authorization"
		r.Header.Add(header, "Bearer "+token)
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	errChan := make(chan error)
	go func(<-chan error) {
		for {
			select {
			case err := <-errChan:
				fmt.Println("ERROR:: " + err.Error())
			}
		}
	}(errChan)

	for {
		for i := 0; i < ratePerSecond; i++ {
			// go client.callTarget(makeRequest(base, errChan), errChan)
			go callTarget(client, errChan)
			callCount++
			if callCount%notifyOn == 0 {
				fmt.Printf("%d calls made\n", callCount)
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func getToken() string {
	return "testToken"
}

func callTarget(c *api.Client, errChan chan error) {
	res, err := c.ValidationFixtureEndpointCustomerGet(context.Background(),
		&api.ValidationFixtureEndpointCustomerGetParameters{
			Customer: "b58b1400-d280-434e-bfe7-a1d544d9b8a1",
			Id:       "Heres a param id"},
		&api.ValidationFixtureEndpointCustomerGetRequestBody{})
	fmt.Printf("GOT BYTES :: %s\n", res)
	if err != nil {
		errChan <- err
		return
	}
}

func (c *client) callTarget(r http.Request, errChan chan<- error) {
	resp, err := c.Do(&r)
	if err != nil {
		errChan <- err
		return
	}
	var data []byte
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			errChan <- err
		}
	}(resp.Body)
	if resp.Body != nil {
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			errChan <- err
			return
		}
	}
	if debug {
		fmt.Printf("GOT :: %s | Body :: %s\n", resp.Status, string(data))
	}
}

func makeRequest(base string, errChan chan<- error) http.Request {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, base, nil)
	if err != nil {
		errChan <- err
		return http.Request{}
	}
	return *req
}

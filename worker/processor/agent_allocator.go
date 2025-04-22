package processor

import (
	"context"
	"errors"
	"github.com/dionofrizal88/go-allocator/config"
	"github.com/dionofrizal88/go-allocator/pkg/integration"
	"github.com/dionofrizal88/go-allocator/pkg/integration/provider/qiscus/activeagent"
	"github.com/dionofrizal88/go-allocator/pkg/integration/provider/qiscus/assignagent"
	"github.com/dionofrizal88/go-allocator/pkg/integration/provider/qiscus/auth"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"strings"
	"time"
)

// AgentAllocator is a struct which is implementation of AgentAllocator worker.
type AgentAllocator struct {
	config config.Configuration
	redis  *redis.Client
}

// Agent is a struct which is implementation of agent worker.
type Agent struct {
	ID                  int `json:"id"`
	CustomerHandleCount int `json:"customer_handle_count"`
}

// NewAgentAllocator is a constructor to initialize agent allocator.
func NewAgentAllocator(c config.Configuration, r *redis.Client) *AgentAllocator {
	return &AgentAllocator{
		config: c,
		redis:  r,
	}
}

// Run is a function to run agent allocator worker.
func (a *AgentAllocator) Run() {
	ctx := context.Background()

	allAgents, err := a.AllAgentAvailable(ctx)
	if err != nil {
		log.Print("error while get all agents error:", err)

		return
	}

	if len(allAgents) == 0 {
		log.Print("Warn all agent is not available\n")

		return
	}

	for {
		log.Printf("Allocator ready to serve \n")

		result, err := a.redis.BRPop(ctx, 0*time.Second, "agent-allocator-queue").Result()
		if err != nil || len(result) < 2 {
			log.Print("No data or Redis error:", err)
			continue
		}

		resultRequest := strings.Split(result[1], ":")
		if len(resultRequest) != 2 {
			log.Print("Value from redis is invalid")
			continue
		}

		log.Printf("Processing request for room id %s \n", resultRequest[1])

		// randomize agents order
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(allAgents), func(i, j int) {
			allAgents[i], allAgents[j] = allAgents[j], allAgents[i]
		})

		// filter base on current_customer_count not more than max assign config
		var agentID, agentIndex int
		var isAgentExist bool

		for i, aa := range allAgents {
			if aa.CustomerHandleCount < a.config.AgentAllocatorWorkerMaxAssign {
				agentID = aa.ID
				isAgentExist = true
				agentIndex = i

				break
			}
		}

		if !isAgentExist {
			log.Print("Warn all agent active is more than max assign\n")
			log.Printf("Retrying allocate again after %d minutes...", a.config.AgentAllocatorWorkerSleep)
			time.Sleep(time.Duration(a.config.AgentAllocatorWorkerSleep) * time.Minute)
			log.Print("Allocator ready to serve again ...\n")

			log.Print("Reset customer handle count all agent\n")

			for _, aa := range allAgents {
				aa.CustomerHandleCount = 0
			}

			agentID = allAgents[0].ID
			agentIndex = 0
		}

		if err := a.processAgent(ctx, agentID, resultRequest); err != nil {
			log.Printf("Processing failed: %v", err)

			// Re-queue
			a.redis.RPush(ctx, "agent-allocator-queue", result[1])
		}

		if err == nil {
			allAgents[agentIndex].CustomerHandleCount += 1
		}
	}
}

// processAgent is a function to process agent allocator worker.
func (a *AgentAllocator) processAgent(ctx context.Context, agentID int, req []string) error {
	// assign agent id into room id
	assignAgentProvider := assignagent.AssignAgent{
		AppConfig: a.config,
	}

	err := assignAgentProvider.SetInput(assignagent.InputAssignAgent{
		RoomID:  req[1],
		AgentID: agentID,
	})
	if err != nil {
		log.Printf("Error while set input post assign agent\n")

		return err
	}

	assignAgentProvider.SetRequestID(uuid.New().String())

	assignAgentIntegration := integration.NewIntegration(&assignAgentProvider)
	_, err = assignAgentIntegration.Info(ctx)
	if err != nil {
		log.Print("Error while send request to assign agent from qiscus")

		return err
	}

	log.Printf("Success assign agent into room id %s\n", req[1])

	return nil
}

// AllAgentAvailable is a function to get all available agent data for allocator worker.
func (a *AgentAllocator) AllAgentAvailable(ctx context.Context) ([]*Agent, error) {
	// auth to get auth token
	authProvider := auth.Auth{
		AppConfig: a.config,
	}

	err := authProvider.SetInput(auth.InputAuth{
		Email:    a.config.QiscusUsername,
		Password: a.config.QiscusPassword,
	})
	if err != nil {
		log.Printf("Error while set input auth\n")

		return nil, err
	}

	authProvider.SetRequestID(uuid.New().String())

	authIntegration := integration.NewIntegration(&authProvider)
	authResponse, err := authIntegration.Info(ctx)
	if err != nil {
		log.Print("Error while send request auth from qiscus")

		return nil, err
	}

	// get available agent using auth token
	resultAuthResponse, ok := authResponse.(*auth.InfoResponse)
	if !ok {
		log.Print("Error while set transform auth response\n")

		return nil, errors.New("error while set transform auth response")
	}

	activeAgentProvider := activeagent.ActiveAgent{
		AppConfig:     a.config,
		Authorization: resultAuthResponse.Data.User.AuthenticationToken,
	}

	activeAgentProvider.SetRequestID(uuid.New().String())

	activeAgentIntegration := integration.NewIntegration(&activeAgentProvider)
	resultActiveAgentResponse, err := activeAgentIntegration.Info(ctx)
	if err != nil {
		log.Print("Error while send request get active agent from qiscus")

		return nil, err
	}

	resultActiveAgentTransform, ok := resultActiveAgentResponse.(*activeagent.InfoResponse)
	if !ok {
		log.Print("Error while set transform active agent response\n")

		return nil, errors.New("error while set transform active agent response")
	}

	var agents []*Agent

	for _, aa := range resultActiveAgentTransform.Data.Agents {
		if aa.IsAvailable {
			agents = append(agents, &Agent{
				ID:                  aa.Id,
				CustomerHandleCount: 0,
			})
		}
	}

	return agents, nil
}

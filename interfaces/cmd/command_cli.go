package cmd

import (
	"github.com/dionofrizal88/go-allocator/config"
	"github.com/dionofrizal88/go-allocator/worker/processor"
	"github.com/go-redis/redis/v8"
	"github.com/urfave/cli/v2"
)

// NewCommand construct a CLI commands.
func NewCommand(
	config config.Configuration,
	redis *redis.Client,
) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "allocator:start",
			Usage: "Start the agent allocator worker",
			Action: func(c *cli.Context) error {

				agentAllocatorProcessor := processor.NewAgentAllocator(config, redis)
				agentAllocatorProcessor.Run()

				return nil
			},
		},
	}
}

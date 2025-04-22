package tests

import (
	"github.com/dionofrizal88/go-allocator/config"
	"github.com/dionofrizal88/go-allocator/pkg/db"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

// TestSuite is a struct represent is self.
type TestSuite struct {
	Config      config.Configuration
	RedisClient *redis.Client
}

// InitTestSuite knows how to initialize test suite.
func InitTestSuite() *TestSuite {
	conf := config.GetConfig()
	if conf.AppName == "" {
		conf = config.GetConfig()
	}

	redisConnection, errRedisConnection := db.NewRedisConnection(conf)
	if errRedisConnection != nil {
		log.Fatalf("Failed to connect into redis db %v", errRedisConnection)
	}

	time.Sleep(5 * time.Second)

	return &TestSuite{
		Config:      conf,
		RedisClient: redisConnection,
	}
}

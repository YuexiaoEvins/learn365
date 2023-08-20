package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const addr = "127.0.0.1:6379"

func getRedisClient(t *testing.T) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:    addr,
		Network: "tcp",
	})
	err := cli.Ping(context.TODO()).Err()
	assert.Nil(t, err)
	return cli
}

func Test_ReplicasRedis(t *testing.T) {
	cli := getRedisClient(t)
	assert.NotNil(t, cli)
}

func Test_Counter(t *testing.T) {
	cli := getRedisClient(t)

	updateCounterWithTimePrecision(cli, 1)
	getCounter(cli)
	time.Sleep(1 * time.Second)
}

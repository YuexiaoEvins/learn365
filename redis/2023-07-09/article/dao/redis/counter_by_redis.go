package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

const dockerRedis = "127.0.0.1:6379"

var counterBucket = []time.Duration{1 * time.Second, 5 * time.Second, 1 * time.Minute, 5 * time.Minute}

var cc = redis.NewClient(&redis.Options{
	Addr:    dockerRedis,
	Network: "tcp",
})

func updateCounterWithTimePrecision(cli *redis.Client, count int) {
	pipeline := cli.Pipeline()

	counterName := "demo_counter"
	ctx := context.TODO()
	for _, bucket := range counterBucket {
		curTs := time.Now().Add(bucket).Unix()
		hName := fmt.Sprintf("%s:%s", bucket, counterName)
		pipeline.ZAdd(ctx, "known:", redis.Z{Member: hName, Score: 0})
		pipeline.HIncrBy(ctx, "count:"+hName, strconv.FormatInt(curTs, 10), int64(count))
	}
	_, err := pipeline.Exec(ctx)
	if err != nil {
		fmt.Printf("pipeline exec failed:%v", err)
	}

}

func getCounter(cli *redis.Client) {
	context.TODO()
	pipeline := cli.Pipeline()

	//counterName := "demo_counter"
	ctx := context.TODO()
	slice := pipeline.HKeys(ctx, "count:")
	_, err := pipeline.Exec(ctx)
	if err != nil {
		fmt.Printf("pipeline exec failed:%v", err)
	}

	for i, s := range slice.Val() {
		fmt.Printf("i:%d s:%s", i, s)
	}

}

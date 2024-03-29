package belajar_golang_redis

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

func TestConnection(t *testing.T) {
	assert.NotNil(t, client)
}

var ctx = context.Background()

func TestPing(t *testing.T) {
	result, err := client.Ping(ctx).Result()
	assert.Nil(t, err)
	assert.Equal(t, "PONG", result)
	fmt.Println(result)
}

func TestString(t *testing.T) {
	client.SetEx(ctx, "name", "Muhammad Zulhaditya Hapiz", 3*time.Second)

	result, err := client.Get(ctx, "name").Result()
	assert.Nil(t, err)
	assert.Equal(t, "Muhammad Zulhaditya Hapiz", result)

	time.Sleep(5 * time.Second)

	result, err = client.Get(ctx, "name").Result()
	assert.NotNil(t, err)
	assert.Equal(t, "", result)
}

func TestList(t *testing.T) {
	// menambahkan data ke list sebelah kanan
	client.RPush(ctx, "names", "Inayah")
	client.RPush(ctx, "names", "Fitri")
	client.RPush(ctx, "names", "Wulandari")

	// lihat data dari sebelah kiri
	assert.Equal(t, "Inayah", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Fitri", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Wulandari", client.LPop(ctx, "names").Val())

	// hapus data list
	client.Del(ctx, "names")
}

func TestSet(t *testing.T) {
	client.SAdd(ctx, "students", "Inayah")
	client.SAdd(ctx, "students", "Inayah")
	client.SAdd(ctx, "students", "Fitri")
	client.SAdd(ctx, "students", "Fitri")
	client.SAdd(ctx, "students", "Wulandari")
	client.SAdd(ctx, "students", "Wulandari")

	assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
	assert.Equal(t, []string{"Inayah", "Fitri", "Wulandari"}, client.SMembers(ctx, "students").Val())
}

func TestSortedSet(t *testing.T) {
	client.ZAdd(ctx, "scores", redis.Z{Score: 100, Member: "Inayah"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 85, Member: "Fitri"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 95, Member: "Wulandari"})

	assert.Equal(t, []string{"Fitri", "Wulandari", "Inayah"}, client.ZRange(ctx, "scores", 0, 2).Val())
	assert.Equal(t, "Inayah", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Wulandari", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Fitri", client.ZPopMax(ctx, "scores").Val()[0].Member)
}

func TestHash(t *testing.T) {
	client.HSet(ctx, "user:1", "id", "1")
	client.HSet(ctx, "user:1", "name", "Inayah")
	client.HSet(ctx, "user:1", "email", "inayah@example.com")

	user := client.HGetAll(ctx, "user:1").Val()
	assert.Equal(t, "1", user["id"])
	assert.Equal(t, "Inayah", user["name"])
	assert.Equal(t, "inayah@example.com", user["email"])

	client.Del(ctx, "user:1")
}

func TestGeoPoint(t *testing.T) {
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Store A",
		Longitude: 106.818489,
		Latitude:  -6.178966,
	})

	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Store B",
		Longitude: 106.821568,
		Latitude:  -6.180662,
	})

	// cari jarak store a dan store b
	distance := client.GeoDist(ctx, "sellers", "Store A", "Store B", "km").Val()
	assert.Equal(t, 0.3892, distance)

	sellers := client.GeoSearch(ctx, "sellers", &redis.GeoSearchQuery{
		Longitude:  106.819143,
		Latitude:   -6.180182,
		Radius:     5,
		RadiusUnit: "km",
	}).Val()

	assert.Equal(t, []string{"Store A", "Store B"}, sellers)
}

func TestHyperLogLog(t *testing.T) {
	client.PFAdd(ctx, "visitors", "inayah", "fitri", "wulandari")
	client.PFAdd(ctx, "visitors", "muhammad", "zulhaditya", "hapiz")
	client.PFAdd(ctx, "visitors", "muhammad", "syapiq", "alfarazi")
	client.PFAdd(ctx, "visitors", "muhammad", "iqbal", "ramadhan")

	assert.Equal(t, int64(10), client.PFCount(ctx, "visitors").Val())
}

func TestPipeLine(t *testing.T) {
	client.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.SetEx(ctx, "name", "Inayah", time.Second*5)
		p.SetEx(ctx, "address", "Jakarta", time.Second*5)
		return nil
	})

	assert.Equal(t, "Inayah", client.Get(ctx, "name").Val())
	assert.Equal(t, "Jakarta", client.Get(ctx, "address").Val())
}

func TestTransaction(t *testing.T) {
	_, err := client.TxPipelined(ctx, func(p redis.Pipeliner) error {
		p.SetEx(ctx, "name", "Hapiz", time.Second*5)
		p.SetEx(ctx, "address", "Batam", time.Second*5)
		return nil
	})

	assert.Nil(t, err)

	assert.Equal(t, "Hapiz", client.Get(ctx, "name").Val())
	assert.Equal(t, "Batam", client.Get(ctx, "address").Val())
}

func TestPublishStream(t *testing.T) {
	for i := 0; i < 10; i++ {
		client.XAdd(ctx, &redis.XAddArgs{
			Stream: "members",
			Values: map[string]interface{}{
				"name":    "Inayah",
				"address": "Indonesia",
			},
		})
	}
}

func TestCreateConsumerGroup(t *testing.T) {
	client.XGroupCreate(ctx, "members", "group-1", "0")
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-1")
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-2")
}

func TestGetStream(t *testing.T) {
	result := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    "group-1",
		Consumer: "consumer-1",
		Streams:  []string{"members", ">"},
		Count:    2,
		Block:    time.Second * 5,
	}).Val()

	for _, stream := range result {
		for _, message := range stream.Messages {
			fmt.Println(message.Values)
		}
	}
}

func TestSubsribePubSub(t *testing.T) {
	pubSub := client.Subscribe(ctx, "channel-1")
	for i := 0; i < 10; i++ {
		message, _ := pubSub.ReceiveMessage(ctx)
		fmt.Println(message.Payload)
	}

	pubSub.Close()
}

func TestPublishPubSub(t *testing.T) {
	for i := 0; i < 10; i++ {
		client.Publish(ctx, "channel-1", "Hello "+strconv.Itoa(i))
	}
}

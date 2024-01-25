package belajar_golang_redis

import (
	"context"
	"fmt"
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

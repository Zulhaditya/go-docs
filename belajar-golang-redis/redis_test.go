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

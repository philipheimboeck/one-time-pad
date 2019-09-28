package persistence

import (
	"encoding/json"
	"fmt"
	"time"

	"../domain"
	"github.com/go-redis/redis"
)

// RedisRepository fetches and stores values from a Redis cluster
type RedisRepository struct {
	client *redis.Client
}

// MakeRedisRepository creates a new instance
func MakeRedisRepository(address string) RedisRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return RedisRepository{client: client}
}

// Get returns a value object by key
func (r *RedisRepository) Get(key string) (domain.Value, error) {
	var value domain.Value

	val, err := r.client.Get(key).Result()
	if err != nil {
		return value, err
	}

	if err = json.Unmarshal([]byte(val), &value); err != nil {
		panic(err)
	}

	return value, nil
}

// Put stores a value object
func (r *RedisRepository) Put(key string, value domain.Value) {
	ttl := time.Duration(value.TTL * 1000 * 1000 * 1000)

	payload, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	err = r.client.Set(key, payload, ttl).Err()
	if err != nil {
		panic(err)
	}
}

// Delete a value
func (r *RedisRepository) Delete(key string) {
	r.client.Del(key)
}

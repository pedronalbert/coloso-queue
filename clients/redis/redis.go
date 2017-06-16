package redis

import (
  "github.com/go-redis/redis"
)

// Client - exported
var Client = redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})

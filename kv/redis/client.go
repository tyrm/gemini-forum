package redis

import (
	"errors"
	"time"

	redisCon "github.com/gomodule/redigo/redis"
)

// Client is a redis clint
type Client struct {
	redisAddress  string
	redisDB       int
	redisPassword string

	db *redisCon.Pool
}

var errCantConnect = errors.New("can't connect to redis")

// NewClient creates a new redis client.
func NewClient(address string, db int, password string) (*Client, error) {
	client := Client{
		redisAddress:  address,
		redisDB:       db,
		redisPassword: password,
	}

	// connect to redis
	client.db = &redisCon.Pool{
		Dial: func() (redisCon.Conn, error) {
			var opts []redisCon.DialOption
			if client.redisPassword != "" {
				opts = append(opts, redisCon.DialPassword(client.redisPassword))
			}
			opts = append(opts, redisCon.DialDatabase(client.redisDB))
			opts = append(opts, redisCon.DialConnectTimeout(100*time.Millisecond))
			opts = append(opts, redisCon.DialReadTimeout(1000*time.Millisecond))

			return redisCon.Dial("tcp", client.redisAddress, opts...)
		},
	}

	logger.Tracef("new redis client created: %s(%d)", address, db)
	return &client, nil
}

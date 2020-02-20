package redis

import (
	"github.com/Chewy-Inc/lets-go/util"
	"github.com/go-redis/redis"
)

var (
	log, _ = util.InitLoggerWithLevel(nil)
)

//Creates a pointer to a redis cluster client. By default looks for a redis cluster at
//localhost on port 7000. Address can be configured with the REDIS_HOST environment
//variable.
func ClusterClient() *redis.ClusterClient {

	clusterURL := util.GetEnv("REDIS_HOST", "127.0.0.1:7000")
	log.Debug("INIT REDIS CONNECT")
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{clusterURL},
	})

	pong, err := clusterClient.Ping().Result()
	log.Debug("REDIS CONNECT ! ", pong)
	if err != nil {
		log.Error("Could not connect to redis!", err)
		return nil
	}

	return clusterClient
}

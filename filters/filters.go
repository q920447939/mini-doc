package filters

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	config "wahaha/conf"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cache"
	"time"
)

func RegisterSession() gin.HandlerFunc {
	store, _ := sessions.NewRedisStore(
		10,
		"tcp",
		config.GetEnv().RedisConfig.REDIS_IP+":"+config.GetEnv().RedisConfig.REDIS_PORT,
		config.GetEnv().RedisConfig.REDIS_PASSWORD,
		[]byte("secret"))
	return sessions.Sessions("mysession", store)
}

func RegisterCache() gin.HandlerFunc {
	var cacheStore persistence.CacheStore
	cacheStore = persistence.NewRedisCache(config.GetEnv().RedisConfig.REDIS_IP+":"+config.GetEnv().RedisConfig.REDIS_PORT, config.GetEnv().RedisConfig.REDIS_PASSWORD, time.Minute)
	return cache.Cache(&cacheStore)
}

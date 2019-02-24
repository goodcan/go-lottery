package conf

type RedisConfig struct {
	Host      string
	Port      int
	IsRunning bool
}

var RedisCacheList = []RedisConfig{
	{
		Host:      "127.0.0.1",
		Port:      6379,
		IsRunning: true,
	},
}

var RedisCache = RedisCacheList[0]

package conf

type RdsConfig struct {
	Host      string
	Port      int
	IsRunning bool
}

var RdsCacheList = []RdsConfig{
	{
		Host:      "127.0.0.1",
		Port:      6379,
		IsRunning: true,
	},
}

var RdsCache = RdsCacheList[0]

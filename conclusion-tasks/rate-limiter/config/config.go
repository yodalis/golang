package config

import (
	"os"
	"strconv"
)

type Config struct {
	RedisHost         string
	RedisPort         string
	IPRequestLimit    int
	TokenRequestLimit int
	BlockTimeSeconds  int
}

func LoadConfig() Config {
	ipLimit, _ := strconv.Atoi(os.Getenv("IP_REQUEST_LIMIT"))
	tokenLimit, _ := strconv.Atoi(os.Getenv("TOKEN_REQUEST_LIMIT"))
	blockTime, _ := strconv.Atoi(os.Getenv("BLOCK_TIME_SECONDS"))

	return Config{
		RedisHost:         os.Getenv("REDIS_HOST"),
		RedisPort:         os.Getenv("REDIS_PORT"),
		IPRequestLimit:    ipLimit,
		TokenRequestLimit: tokenLimit,
		BlockTimeSeconds:  blockTime,
	}
}

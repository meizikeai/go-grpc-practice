package config

import (
	"os"

	"go-grpc-practice/libs/types"
)

var Address = "127.0.0.1:9527"

var mysqlConfig = map[string]types.ConfMySQL{
	"ailab-test": {
		Master:   []string{"127.0.0.1:3306"},
		Slave:    []string{"127.0.0.1:3306"},
		Username: "blued",
		Password: "g3bkshwqcj4wcSMr",
		Database: "blued",
	},
	"ailab-release": {
		Master:   []string{"127.0.0.1:3306"},
		Slave:    []string{"127.0.0.1:3306", "127.0.0.1:3306"},
		Username: "test",
		Password: "yintai@123",
		Database: "test",
	},
}

var redisConfig = map[string]types.ConfRedis{
	"ailab-test": {
		Master:   []string{"127.0.0.1:6379"},
		Password: "",
		Db:       0,
	},
	"ailab-release": {
		Master:   []string{"127.0.0.1:6379"},
		Password: "",
		Db:       0,
	},
}

func isProduction() bool {
	result := false

	mode := os.Getenv("GGP_MODE")

	if mode == "release" {
		result = true
	}

	return result
}

func GetMySQLConfig() types.FullConfMySQL {
	env := isProduction()

	confMySQL := mysqlConfig["ailab-release"]

	if env == false {
		confMySQL = mysqlConfig["ailab-test"]
	}

	result := types.FullConfMySQL{
		"default": confMySQL,
	}

	return result
}

func GetRedisConfig() types.FullConfRedis {
	env := isProduction()

	confRedis := redisConfig["ailab-release"]

	if env == false {
		confRedis = redisConfig["ailab-test"]
	}

	result := types.FullConfRedis{
		"default": confRedis,
	}

	return result
}

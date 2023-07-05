package config

import (
	"os"

	"go-grpc-practice/libs/types"
)

var Address = "127.0.0.1:9527"

var mysqlConfig = map[string]types.ConfMySQL{
	"default-test": {
		Master:   []string{"127.0.0.1:3306"},
		Slave:    []string{"127.0.0.1:3306"},
		Username: "blued",
		Password: "g3bkshwqcj4wcSMr",
		Database: "blued",
	},
	"default-release": {
		Master:   []string{"127.0.0.1:3306"},
		Slave:    []string{"127.0.0.1:3306", "127.0.0.1:3306"},
		Username: "test",
		Password: "yintai@123",
		Database: "test",
	},
}

var redisConfig = map[string]types.ConfRedis{
	"default-test": {
		Master:   []string{"127.0.0.1:6379"},
		Password: "",
		Db:       0,
	},
	"default-release": {
		Master:   []string{"127.0.0.1:6379"},
		Password: "",
		Db:       0,
	},
}

func getMode() string {
	mode := os.Getenv("GGP_MODE")

	if mode == "" {
		mode = "test"
	}

	return mode
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
	mode := getMode()
	result := types.FullConfMySQL{}

	data := []string{
		"default",
	}

	for _, v := range data {
		result[v] = mysqlConfig[v+"-"+mode]
	}

	return result
}

func GetRedisConfig() types.FullConfRedis {
	mode := getMode()
	result := types.FullConfRedis{}

	data := []string{
		"default",
	}

	for _, v := range data {
		result[v] = redisConfig[v+"-"+mode]
	}

	return result
}

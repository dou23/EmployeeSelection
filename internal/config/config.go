package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func GetWebServicePort() int {
	portStr := os.Getenv("WEB_SERVICE_PORT")
	portInt, err := strconv.Atoi(portStr)
	if err != nil {
		panic("Error: Fail to Get WEB_SERVICE_PORT," + err.Error())
	}
	return portInt
}

func GetNodeId() int {
	nodeIdStr := os.Getenv("DB_NODE_ID")
	nodeIdInt, err := strconv.Atoi(nodeIdStr)
	if err != nil {
		panic("Error: Fail to Get WEB_SERVICE_PORT," + err.Error())
	}
	return nodeIdInt
}

// 加载env.default到.env环境变量中
func LoadEnv() {
	err := godotenv.Load("env.default")
	if err != nil {
		panic("godotenv Error: " + err.Error())
	}
}

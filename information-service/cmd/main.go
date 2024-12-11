package main

import (
	"fmt"
	"information-service/pkg/utils"
)

func main() {
	fmt.Println("User service started!")

	utils.InitLogs()
	utils.LoadEnv()
	utils.CreateRedisConn()
	utils.InitDBConnection()
	utils.StartMigrations()
	utils.InitValidator()
}

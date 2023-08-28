package main

import (
	"fmt"
	"net"
	"os"

	"go/protos"

	"go-grpc-practice/config"
	"go-grpc-practice/controllers"
	"go-grpc-practice/libs/log"
	"go-grpc-practice/libs/tool"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	// tool.HandleMySQLClient()
	// tool.HandleRedisClient()

	log.HandleLogger("go-grpc-practice")
}

func main() {
	tool.SignalHandler(func() {
		tool.CloseMySQL()
		tool.CloseRedis()

		tool.Stdout("Server Shutdown")

		os.Exit(0)
	})

	client, err := net.Listen("tcp", config.Server)

	if err != nil {
		panic(fmt.Sprintf("Failed to listen: %v", err))
	}

	server := grpc.NewServer()

	protos.RegisterTestRpcServiceServer(server, &controllers.Server{})

	reflection.Register(server)

	fmt.Println("Listen and Server running on", config.Server)

	if err := server.Serve(client); err != nil {
		panic(fmt.Sprintf("Failed to serve: %v", err))
	}
}

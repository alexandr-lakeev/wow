package main

import (
	"context"
	"fmt"
	"github.com/alexandr-lakeev/wow.git/internal/pkg/protocol/emoji"
	"math"
	"os/signal"
	"syscall"

	proofOfWork "github.com/alexandr-lakeev/wow.git/internal/pkg/proof_of_work"
	"github.com/alexandr-lakeev/wow.git/internal/pkg/proof_of_work/hasher"
	tcpServer "github.com/alexandr-lakeev/wow.git/internal/server"
)

func main() {
	port := "8888"
	hashcash := proofOfWork.New(hasher.New(), "ver1", 1, "0", math.MaxInt)
	protocolServer := emoji.NewServer(hashcash)
	server := tcpServer.New(protocolServer)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	fmt.Println("TCP server is starting on port " + port)

	go func() {
		if err := server.Run(ctx, port); err != nil {
			fmt.Println("TCP server error:", err)
		}
	}()

	<-ctx.Done()

	fmt.Println("TCP server shutdown")
}

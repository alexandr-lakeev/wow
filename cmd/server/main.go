package main

import (
	"context"
	"log"
	"math"
	"os/signal"
	"syscall"

	pkgLogger "github.com/alexandr-lakeev/wow/internal/pkg/logger"
	proofOfWork "github.com/alexandr-lakeev/wow/internal/pkg/proof_of_work"
	"github.com/alexandr-lakeev/wow/internal/pkg/proof_of_work/hasher"
	"github.com/alexandr-lakeev/wow/internal/pkg/protocol/emoji"
	"github.com/alexandr-lakeev/wow/internal/pkg/quotes"
	tcpServer "github.com/alexandr-lakeev/wow/internal/server"
)

func main() {
	logger, err := pkgLogger.New()
	if err != nil {
		log.Fatal(err)
	}

	// todo configuration
	port := "8888"
	hashcash := proofOfWork.New(hasher.New(), "ver1", 6, "0", math.MaxInt)
	protocolServer := emoji.NewServer(
		hashcash,
		// https://www.wix.com/blog/motivational-quotes
		quotes.NewService([]quotes.Quote{
			"I learned that courage was not the absence of fear, but the triumph over it. The brave man is not he who does not feel afraid, but he who conquers that fear.",
			"If you believe it will work, you'll see opportunities. If you believe it won't, you will see obstacles.",
			"Believe you can and you're halfway there.",
			"Do one thing every day that scares you.",
			"I didn't get there by wishing for it or hoping for it, but by working for it.",
		}),
		logger,
	)

	server := tcpServer.New(protocolServer)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	logger.Info("TCP server is starting on port " + port)

	go func() {
		if err := server.Run(ctx, port); err != nil {
			logger.Error("TCP server error:" + err.Error())
		}
	}()

	<-ctx.Done()
}

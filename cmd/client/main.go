package main

import (
	"context"
	"log"
	"math"

	tcpClient "github.com/alexandr-lakeev/wow/internal/client"
	pkgLogger "github.com/alexandr-lakeev/wow/internal/pkg/logger"
	proofOfWork "github.com/alexandr-lakeev/wow/internal/pkg/proof_of_work"
	"github.com/alexandr-lakeev/wow/internal/pkg/proof_of_work/hasher"
	"github.com/alexandr-lakeev/wow/internal/pkg/protocol/emoji"
)

func main() {
	logger, err := pkgLogger.New()
	if err != nil {
		log.Fatal(err)
	}

	// todo configuration
	hashcash := proofOfWork.New(hasher.New(), "ver1", 1, "0", math.MaxInt)
	protocol := emoji.NewClient(hashcash, logger)
	client := tcpClient.New(protocol)

	if err := client.Run(context.Background(), "server:8888"); err != nil {
		logger.Error(err.Error())
	}
}

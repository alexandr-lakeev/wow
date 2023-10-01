package main

import (
	"context"
	"fmt"

	tcpClient "github.com/alexandr-lakeev/wow.git/internal/client"
)

func main() {
	client := tcpClient.New()

	if err := client.Run(context.Background(), "server:8888"); err != nil {
		fmt.Println(err)
	}
}

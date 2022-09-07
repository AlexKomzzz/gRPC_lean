package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/AlexKomzzz/gRPC_lean/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal("error type X")
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal("error type Y")
	}

	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect")
	}

	defer conn.Close()

	client := api.NewAdderClient(conn)
	res, err := client.Add(context.Background(), &api.AddRequest{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatal("error responce")
	}

	fmt.Println(res.GetResult())
}

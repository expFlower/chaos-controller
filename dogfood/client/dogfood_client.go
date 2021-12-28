// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2021 Datadog, Inc.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "github.com/DataDog/chaos-controller/dogfood/chaosdogfood"
	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var serverAddr string

func init() {
	var serverPort int

	var serverHostname string

	flag.StringVar(&serverHostname, "server_hostname", "<service>.<namespace>.svc.cluster.local", "Hostname of dogfood server")
	flag.IntVar(&serverPort, "server_port", 50000, "Port where gRPC server is running")
	flag.Parse()

	serverAddr = fmt.Sprintf("%s:%d", serverHostname, serverPort)
}

func orderWithTimeout(client pb.ChaosDogfoodClient, animal string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.Order(ctx, &pb.FoodRequest{Animal: animal})
	if err != nil {
		return "", err
	}

	return res.Message, nil
}

func getCatalogWithTimeout(client pb.ChaosDogfoodClient) ([]*pb.CatalogItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.GetCatalog(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return res.Items, nil
}

func main() {
	// create and eventually close connection
	fmt.Printf("connecting to %v...\n", serverAddr)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("error closing gRPC connection to dogfood server", "error", err)
		}
	}()

	// generate and use client
	client := pb.NewChaosDogfoodClient(conn)

	for {
		// visually mark a new loop in logs
		fmt.Println("x")

		items, err := getCatalogWithTimeout(client)
		if err != nil {
			fmt.Printf("| ERROR getting catalog:%v\n", err.Error())
		}

		fmt.Printf("| got catalog: %v items returned\n", strconv.Itoa(len(items)))
		time.Sleep(time.Second)

		order, err := orderWithTimeout(client, "cat")
		if err != nil {
			fmt.Printf("| ERROR ordering food: %v\n", err.Error())
		}

		fmt.Printf("| ordered: %v\n", order)
		time.Sleep(time.Second)
	}
}
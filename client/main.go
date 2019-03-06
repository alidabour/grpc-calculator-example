package main

import (
	"bufio"
	"client/entity"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	port := flag.String("port", "8000", "Port to listen on")
	flag.Parse()

	conn, err := grpc.Dial(fmt.Sprintf(":%s", *port), grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v\n", err)
		os.Exit(1)
	}
	client := entity.NewCalcClient(conn)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		r, err := client.Eval(context.Background(), &entity.Request{Value: scanner.Text()})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Fprintln(os.Stderr, r.Value)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}

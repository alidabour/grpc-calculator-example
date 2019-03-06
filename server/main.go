package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"server/adapters/grpcAPI"
	"server/usecase"

	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	port := flag.String("port", "8000", "Port to listen on")
	flag.Parse()

	srv := grpc.NewServer()
	evalUC := usecase.New()
	grpcAPI.NewServer(srv, evalUC)

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Fatalf("could not listen to %v:- %v", port, err)
	}
	log.Fatal(srv.Serve(l))
}

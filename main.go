package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/sharmarajdaksh/tryst-with-gRPC/protos/currency"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := currency.NewCurrency(log)

	currency.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9902")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}

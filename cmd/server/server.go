package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/iwdmb/kvstore-grpc/proto"
	"github.com/iwdmb/kvstore-grpc/service"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var RootCmd = &cobra.Command{
	Use:   "kvStore",
	Short: "kvStore Server",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start GRPC Server",
	Run: func(cmd *cobra.Command, args []string) {
		server()
	},
}

var (
	GRPCHost string
	HTTPHost string
	GRPCPort string
	HTTPPort string
	Debug    bool
)

func server() {
	// Init zap
	var c zap.Config

	if Debug {
		c = zap.NewDevelopmentConfig()
	} else {
		c = zap.NewProductionConfig()
		c.DisableStacktrace = true
	}

	l, e := c.Build()
	if e != nil {
		panic(e)
	}

	zap.ReplaceGlobals(l)

	// Init gRPC
	addr := net.JoinHostPort(GRPCHost, GRPCPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	forver := make(chan struct{})

	// New GRPC Server
	s := grpc.NewServer()
	proto.RegisterKVServiceServer(s, service.GetService())

	go func() {
		zap.L().Info(
			fmt.Sprintf("gRPC server started on %s", addr),
		)

		err := s.Serve(lis)
		if err != nil {
			zap.L().Panic(
				"s.Serve(lis) error",
				zap.Error(err),
			)
		}
	}()
	// New GRPC Server

	// New HTTP Server
	conn, err := grpc.DialContext(context.Background(), net.JoinHostPort(HTTPHost, HTTPPort), grpc.WithInsecure())
	if err != nil {
		zap.L().Info(
			"grpc DialContext error",
			zap.Error(err),
		)
	}

	go func() {
		gwMux := runtime.NewServeMux()
		err := proto.RegisterKVServiceHandler(context.Background(), gwMux, conn)
		if err != nil {
			zap.L().Info(
				"RegisterKVServiceHandler",
				zap.Error(err),
			)
		}

		zap.L().Info(
			fmt.Sprintf("http server started on %s", net.JoinHostPort(HTTPHost, HTTPPort)),
		)

		err = http.ListenAndServe(net.JoinHostPort(HTTPHost, HTTPPort), gwMux)
		if err != nil {
			zap.L().Panic(
				"http.ListenAndServe error",
				zap.Error(err),
			)
		}
	}()

	<-forver
}

func init() {
	RootCmd.AddCommand(serverCmd)
	// Here you will define your flags and configuration settings.
	serverCmd.Flags().StringVar(&GRPCHost, "grpchost", "127.0.0.1", "GRPC Server Host")
	serverCmd.Flags().StringVar(&HTTPHost, "httphost", "127.0.0.1", "HTTP Server Host")
	serverCmd.Flags().StringVar(&GRPCPort, "grpcport", "6666", "GRPC Server Port")
	serverCmd.Flags().StringVar(&HTTPPort, "httpport", "7777", "HTTP Server Port")
	serverCmd.Flags().BoolVarP(&Debug, "debug", "d", false, "Start Debug Mode")
}

func main() {
	Execute()
}

package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/iwdmb/kvStore/proto"
	"github.com/iwdmb/kvStore/service"
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
	Hostname string
	Port     string
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
	addr := fmt.Sprintf("%s:%s", Hostname, Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	forver := make(chan struct{})

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

	<-forver
}

func init() {
	RootCmd.AddCommand(serverCmd)
	// Here you will define your flags and configuration settings.
	serverCmd.Flags().StringVarP(&Hostname, "hostname", "n", "127.0.0.1", "Server Hostname")
	serverCmd.Flags().StringVarP(&Port, "port", "p", "7777", "Server Port")
	serverCmd.Flags().BoolVarP(&Debug, "debug", "d", false, "Start Debug Mode")
}

func main() {
	Execute()
}

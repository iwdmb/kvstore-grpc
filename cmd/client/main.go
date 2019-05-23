package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/iwdmb/kvstore-grpc/proto"
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

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start GRPC Client",
}

var (
	Host  string
	Port  string
	Debug bool
)

func InitZap() {
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
}

func init() {
	RootCmd.AddCommand(clientCmd)
	// Here you will define your flags and configuration settings.
	clientCmd.Flags().StringVarP(&Host, "host", "n", "127.0.0.1", "gRPC Server Host")
	clientCmd.Flags().StringVarP(&Port, "port", "p", "7777", "gRPC Server Port")
	clientCmd.Flags().BoolVarP(&Debug, "debug", "d", false, "Start Debug Mode")
}

func main() {
	// Init
	InitZap()
	// Init

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, net.JoinHostPort(Host, Port), grpc.WithInsecure())
	if err != nil {
		zap.L().Panic(
			"grpc.DialContext error",
			zap.Error(err),
		)
	}
	defer conn.Close()

	c := pb.NewKVServiceClient(conn)
	pbSReq := pb.SetRequest{
		Key:   "Hello",
		Value: []byte(" world!"),
	}

	pbSResp, err := c.Set(ctx, &pbSReq)
	if err != nil {
		zap.L().Error(
			"kvStore.Set error",
			zap.Error(err),
		)
	}

	zap.L().Info(
		"kvStore.Set",
		zap.Any("", pbSResp),
	)

	pbGReq := pb.GetRequest{
		Key: "Hello",
	}

	pbGResp, err := c.Get(ctx, &pbGReq)
	if err != nil {
		zap.L().Error(
			"kvStore.Get error",
			zap.Error(err),
		)
	}

	zap.L().Info(
		"kvStore.Get",
		zap.Any("", pbGResp),
	)

}

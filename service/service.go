package service

import (
	"context"

	"github.com/gogo/protobuf/types"
	pb "github.com/iwdmb/kvstore-grpc/proto"
	"github.com/ming-go/pkg/kvstore"
)

var kv *kvstore.KVStore

func init() {
	kv = kvstore.NewKVStore()
}

func GetService() pb.KVServiceServer {
	return &svc{}
}

type svc struct{}

func (s *svc) Set(ctx context.Context, in *pb.SetRequest) (*pb.SetResponse, error) {
	kv.SetBytes(in.Key, in.Value)

	p := &pb.SetResponse{
		Key: in.Key,
		Status: &pb.Status{
			Code:      "200",
			Message:   "OK",
			Timestamp: types.TimestampNow(),
		},
	}

	return p, nil
}

func (s *svc) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	b, err := kv.GetBytes(in.Key)
	if err != nil {

	}

	p := &pb.GetResponse{
		Key:   in.Key,
		Value: b,
		Status: &pb.Status{
			Code:      "200",
			Message:   "OK",
			Timestamp: types.TimestampNow(),
		},
	}

	return p, nil
}

func (s *svc) Del(ctx context.Context, in *pb.DelRequest) (*pb.DelResponse, error) {
	return nil, nil
}

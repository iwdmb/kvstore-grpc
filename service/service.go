package service

import (
	"context"

	"github.com/iwdmb/kvStore/proto"
	"github.com/ming-go/pkg/kvstore"
)

var kv *kvstore.KVStore

func init() {
	kv = kvstore.NewKVStore()
}

func GetService() proto.KVServiceServer {
	return &svc{}
}

type svc struct{}

func (s *svc) Set(ctx context.Context, in *proto.SetRequest) (*proto.SetResponse, error) {
	return nil, nil
}

func (s *svc) Get(ctx context.Context, in *proto.GetRequest) (*proto.GetResponse, error) {
	b, err := kv.GetBytes(in.Key)
	if err != nil {

	}

	p := &proto.GetResponse{
		Key:   in.Key,
		Value: b,
	}

	return p, nil
}

func (s *svc) Del(ctx context.Context, in *proto.DelRequest) (*proto.DelResponse, error) {
	return nil, nil
}

// +build tools

package tools

import (
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/protobuf"
	_ "github.com/gogo/protobuf/protoc-gen-gofast"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
)

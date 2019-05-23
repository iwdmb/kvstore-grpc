generate-proto:
	@protoc \
		-I ./proto \
		-I ./vendor/github.com/gogo/googleapis/ \
		--gogofaster_out=plugins=grpc,\ \
	Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
	Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
	Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:./proto \
	kv.proto

generate-grpc-gateway:
	@protoc \
		-I ./proto \
		-I ./vendor/github.com/gogo/googleapis/ \
		--grpc-gateway_out=allow_patch_feature=false,\
	Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
	Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
	Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:./proto \
	kv.proto

package grpc

import (
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
  "micro/go-micro/codec"
  "micro/go-micro/codec/jsonrpc"
  "micro/go-micro/codec/protorpc"
	"google.golang.org/grpc"
)

type jsonCodec struct{}
type protoCodec struct{}
type bytesCodec struct{}

var (
	defaultGRPCCodecs = map[string]grpc.Codec{
		"application/grpc+json":  jsonCodec{},
		"application/grpc+proto": protoCodec{},
		"application/grpc+bytes": bytesCodec{},
	}

	defaultRPCCodecs = map[string]codec.NewCodec{
		"application/json":         jsonrpc.NewCodec,
		"application/json-rpc":     jsonrpc.NewCodec,
		"application/protobuf":     protorpc.NewCodec,
		"application/proto-rpc":    protorpc.NewCodec,
		"application/octet-stream": protorpc.NewCodec,
	}
)

func (protoCodec) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (protoCodec) Unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}

func (protoCodec) String() string {
	return "proto"
}

func (bytesCodec) Marshal(v interface{}) ([]byte, error) {
	b, ok := v.(*[]byte)
	if !ok {
		return nil, fmt.Errorf("failed to marshal: %v is not type of *[]byte", v)
	}
	return *b, nil
}

func (bytesCodec) Unmarshal(data []byte, v interface{}) error {
	b, ok := v.(*[]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal: %v is not type of *[]byte", v)
	}
	*b = data
	return nil
}

func (bytesCodec) String() string {
	return "bytes"
}

func (jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (jsonCodec) String() string {
	return "json"
}

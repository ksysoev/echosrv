package main

import (
	"context"

	"github.com/ksysoev/echosrv/gen/proto"
	"github.com/redis/go-redis/v9"
)

type Service struct{}

func (s *Service) Echo(ctx context.Context, in *proto.StringMessage) (*proto.StringMessage, error) {
	return in, nil
}

func main() {
	redis := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	var service = proto.NewEchoServiceService(&Service{})

	_ = proto.NewRedisEchoService(redis, service).Serve()
}

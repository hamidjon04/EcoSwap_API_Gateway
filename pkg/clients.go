package pkg

import (
	"api_gateway/config"
	"api_gateway/genproto/users"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUsersClient(cfg config.Config) users.UsersServiceClient {
	grpc, err := grpc.NewClient(cfg.USER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return nil
	}
	conn := users.NewUsersServiceClient(grpc)
	return conn
}

package pkg

import (
	"api_gateway/config"
	"api_gateway/genproto/users"
)


func NewUsersClient(cfg config.Config){
	conn, err := users.NewUsersServiceClient(cfg.)
}
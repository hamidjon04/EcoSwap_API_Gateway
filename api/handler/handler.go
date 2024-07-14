package handler

import (
	"api_gateway/config"
	"api_gateway/genproto/users"
	"database/sql"
	"log/slog"
)



type Handler struct{
	UsersService users.UsersServiceClient
	Logger *slog.Logger
}

func NewHandler(db *sql.DB)*Handler{
	cfg := config.Load()
	return &Handler{
		UsersService: ,
	}
}


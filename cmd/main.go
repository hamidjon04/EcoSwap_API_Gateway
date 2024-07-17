package main

import (
	"api_gateway/api"
	"api_gateway/config"
)

func main(){
	router := api.Router()
	router.Run(config.Load().API_GATEWAY)
}
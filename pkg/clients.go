package pkg

import (
	"api_gateway/config"
	"api_gateway/genproto/challenges"
	"api_gateway/genproto/items"
	"api_gateway/genproto/rating"
	"api_gateway/genproto/recycling_center"
	"api_gateway/genproto/swaps"
	"api_gateway/genproto/users"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUsersClient(cfg config.Config) users.UsersServiceClient {
	grpc, err := grpc.NewClient(cfg.USER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Error", err)
		log.Println(err)
		return nil
	}
	conn := users.NewUsersServiceClient(grpc)
	return conn
}

func NewChallengeClient(cfg config.Config) challenges.ChallengesClient{
	grpc, err := grpc.NewClient(cfg.ITEM_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Println("Error", err)
		log.Println(err)
		return nil
	}
	conn := challenges.NewChallengesClient(grpc)
	return conn
}

func NewItemService(cfg config.Config) items.ItemsClient{
	grpc, err := grpc.NewClient(cfg.ITEM_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Println("Error", err)
		log.Println(err)
		return nil
	}
	conn := items.NewItemsClient(grpc)
	return conn
}

func NewRatingService(cfg config.Config) rating.RatingClient{
	grpc, err := grpc.NewClient(cfg.ITEM_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Println("Error", err)
		log.Println(err)
		return nil
	}
	conn := rating.NewRatingClient(grpc)
	return conn
}

func NewCenterService(cfg config.Config) recycling_center.RecyclingCenterClient{
	grpc, err := grpc.NewClient(cfg.ITEM_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Println("Error", err)
		log.Println(err)
		return nil
	}
	conn := recycling_center.NewRecyclingCenterClient(grpc)
	return conn
}

func NewSwapService(cfg config.Config) swaps.SwapsClient{
	grpc, err := grpc.NewClient(cfg.ITEM_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Println("Error", err)
		log.Println(err)
		return nil
	}
	conn := swaps.NewSwapsClient(grpc)
	return conn
}
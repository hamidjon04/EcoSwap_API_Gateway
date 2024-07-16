package handler

import (
	"api_gateway/config"
	"api_gateway/genproto/challenges"
	"api_gateway/genproto/items"
	"api_gateway/genproto/rating"
	"api_gateway/genproto/recycling_center"
	"api_gateway/genproto/swaps"
	"api_gateway/genproto/users"
	"api_gateway/pkg"
	"api_gateway/pkg/logger"
	"log/slog"
)

type Handler struct {
	UsersService     users.UsersServiceClient
	ChallengeService challenges.ChallengesClient
	ItemService      items.ItemsClient
	RatingService    rating.RatingClient
	CenterService    recycling_center.RecyclingCenterClient
	SwapService      swaps.SwapsClient
	Logger           *slog.Logger
}

func NewHandler() *Handler {
	cfg := config.Load()
	return &Handler{
		UsersService:     pkg.NewUsersClient(cfg),
		ChallengeService: pkg.NewChallengeClient(cfg),
		ItemService:      pkg.NewItemService(cfg),
		RatingService:    pkg.NewRatingService(cfg),
		CenterService:    pkg.NewCenterService(cfg),
		SwapService:      pkg.NewSwapService(cfg),
		Logger:           logger.NewLogger(),
	}
}

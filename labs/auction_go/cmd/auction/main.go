package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yodalis/golang/labs/auction_go/config/database/mongodb"
	"github.com/yodalis/golang/labs/auction_go/internal/infra/api/web/controller/auction_controller"
	"github.com/yodalis/golang/labs/auction_go/internal/infra/api/web/controller/bid_controller"
	"github.com/yodalis/golang/labs/auction_go/internal/infra/api/web/controller/user_controller"
	"github.com/yodalis/golang/labs/auction_go/internal/infra/database/auction"
	"github.com/yodalis/golang/labs/auction_go/internal/infra/database/bid"
	"github.com/yodalis/golang/labs/auction_go/internal/infra/database/user"
	"github.com/yodalis/golang/labs/auction_go/internal/usecase/auction_usecase"
	"github.com/yodalis/golang/labs/auction_go/internal/usecase/bid_usecase"
	"github.com/yodalis/golang/labs/auction_go/internal/usecase/user_usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	db, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	router := gin.Default()
	userController, bidController, auctionsController := initDependencies(db)

	checkExpiredAuctions(ctx, db)

	router.GET("/auctions", auctionsController.FindAuctions)
	router.GET("/auctions/:auctionId", auctionsController.FindAuctionById)
	router.POST("/auctions", auctionsController.CreateAuction)
	router.GET("/auction/winner/:auctionId", auctionsController.FindWinningBidByAuctionId)
	router.POST("/bid", bidController.CreateBid)
	router.GET("/bid/:auctionId", bidController.FindBidByAuctionId)
	router.GET("/user/:userId", userController.FindUserById)

	router.Run(":8080")
}

func checkExpiredAuctions(ctx context.Context, db *mongo.Database) {
	auctionRepository := auction.NewAuctionRepository(db)
	go auctionRepository.StartAuctionExpirationChecker(ctx)
}

func initDependencies(db *mongo.Database) (userController *user_controller.UserController, bidController *bid_controller.BidController, auctionController *auction_controller.AuctionController) {
	auctionRepository := auction.NewAuctionRepository(db)
	bidRepository := bid.NewBidRepository(db, auctionRepository)
	userRepository := user.NewUserRepository(db)

	userController = user_controller.NewUserController(user_usecase.NewUserUseCase(userRepository))
	auctionController = auction_controller.NewAuctionController(auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository))
	bidController = bid_controller.NewBidController(bid_usecase.NewBidUseCase(bidRepository))

	return
}

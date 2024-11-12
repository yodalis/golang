package auction

import (
	"context"
	"fmt"
	"fullcycle-auction_go/configuration/logger"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
}
type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) StartAuctionExpirationChecker(ctx context.Context) {
	fmt.Println("Starting check for auctions expired...")
	intervalStr := os.Getenv("CHECK_INTERVAL")
	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		logger.Error(fmt.Sprintf("invalid CHECK_INTERVAL: %v", err), err)
		return
	}
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ar.checkAndCloseExpiredAuctions(ctx)
		case <-ctx.Done():
			fmt.Println("Auction expiration checker stopped")
			return
		}
	}
}

func (ar *AuctionRepository) checkAndCloseExpiredAuctions(ctx context.Context) {
	expiredAuctions, err := ar.findExpiredAuctions(ctx)
	if len(expiredAuctions) == 0 {
		return
	}

	fmt.Println(fmt.Sprintf("expired auctions: %v", expiredAuctions))
	if err != nil {
		logger.Error(fmt.Sprintf("error finding expired auctions: %v\n", err), err)
		return
	}

	for _, auction := range expiredAuctions {
		auction.Status = 1
		fmt.Println("Updating status...", auction)
		if err := ar.updateAuctionStatus(ctx, auction.Id); err != nil {
			logger.Error(fmt.Sprintf("error closing auction %v: %v\n", auction.Id, err), err)
		}
	}
}

func (ar *AuctionRepository) findExpiredAuctions(ctx context.Context) ([]auction_entity.Auction, error) {
	var (
		expiredAuctions []AuctionEntityMongo
		response        []auction_entity.Auction
	)

	durationStr := os.Getenv("AUCTION_DURATION")
	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		return nil, err
	}

	expirationThreshold := time.Now().Add(-time.Duration(duration) * time.Second).Unix()

	filter := bson.M{
		"status":    0,
		"timestamp": bson.M{"$lt": expirationThreshold},
	}

	cursor, err := ar.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &expiredAuctions); err != nil {
		return nil, err
	}

	for _, value := range expiredAuctions {
		response = append(response, auction_entity.Auction{
			Id:          value.Id,
			ProductName: value.ProductName,
			Status:      value.Status,
			Category:    value.Category,
			Description: value.Description,
			Condition:   value.Condition,
		})
	}

	return response, nil
}

func (ar *AuctionRepository) updateAuctionStatus(ctx context.Context, auctionId string) *internal_error.InternalError {
	fmt.Println(auctionId)
	filter := bson.M{"_id": auctionId}

	update := bson.M{"$set": bson.M{"status": 1}}

	_, err := ar.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println("Error trying to update auction status", err)
		return internal_error.NewInternalServerError("Error trying to update auction status")
	}

	fmt.Println(fmt.Sprintf("Auction status updated successfully for auction ID: %s", auctionId))
	return nil
}

func (ar *AuctionRepository) CreateAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction) *internal_error.InternalError {
	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}
	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert auction", err)
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	return nil
}

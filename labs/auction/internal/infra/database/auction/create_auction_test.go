package auction_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/infra/database/auction"
	"fullcycle-auction_go/internal/internal_error"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func Test_Create_Auction_StartAuctionExpirationChecker(t *testing.T) {
	os.Setenv("CHECK_INTERVAL", "1")
	os.Setenv("AUCTION_DURATION", "2")
	collectionName := "auctions"

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	ctx := context.Background()

	mtestDb.Run("should update status", func(mt *mtest.T) {
		auctionEntity := &auction_entity.Auction{
			Id:          uuid.NewString(),
			ProductName: "Product Test",
			Category:    "Some category",
			Description: "Some description",
			Condition:   0,
			Status:      0,
			Timestamp:   time.Now().Add(-2 * time.Second),
		}

		firstBatch := mtest.CreateCursorResponse(1, fmt.Sprintf("%s.%s", mt.DB.Name(), collectionName), mtest.FirstBatch, convertToBson(auctionEntity))
		killCursors := mtest.CreateCursorResponse(0, fmt.Sprintf("%s.%s", mt.DB.Name(), collectionName), mtest.NextBatch)

		mt.AddMockResponses(firstBatch, killCursors)
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		auctionRepo := auction.NewAuctionRepository(mt.DB)
		go auctionRepo.StartAuctionExpirationChecker(ctx)

		time.Sleep(5 * time.Second)

		mt.AddMockResponses(mtest.CreateCursorResponse(1, fmt.Sprintf("%s.%s", mt.DB.Name(), collectionName), mtest.FirstBatch, bson.D{
			{Key: "_id", Value: auctionEntity.Id},
			{Key: "status", Value: 1},
		}))
		var updatedAuction auction_entity.Auction
		dbErr := mt.DB.Collection("auctions").FindOne(ctx, bson.M{"_id": auctionEntity.Id}).Decode(&updatedAuction)

		if dbErr != nil {
			err := &internal_error.InternalError{
				Message: dbErr.Error(),
				Err:     dbErr.Error(),
			}
			log.Fatal("Error fetching updated auction: ", err)
		}

		assert.NotEqual(t, auctionEntity.Status, updatedAuction.Status)
	})
}

func convertToBson(auctionEntity *auction_entity.Auction) bson.D {
	return bson.D{
		{Key: "_id", Value: auctionEntity.Id},
		{Key: "product_name", Value: auctionEntity.ProductName},
		{Key: "category", Value: auctionEntity.Category},
		{Key: "description", Value: auctionEntity.Description},
		{Key: "condition", Value: auctionEntity.Condition},
		{Key: "status", Value: auctionEntity.Status},
		{Key: "timestamp", Value: auctionEntity.Timestamp.Unix()},
	}
}

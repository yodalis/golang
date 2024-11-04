package bid

import (
	"context"
	"fmt"
	"time"

	"github.com/yodalis/golang/labs/auction_go/config/logger"
	"github.com/yodalis/golang/labs/auction_go/internal/entity/bid_entity"
	"github.com/yodalis/golang/labs/auction_go/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (br *BidRepository) FindBidByAuctionId(ctx context.Context, auctionId string) ([]bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auctionId": auctionId}

	cursor, err := br.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error(fmt.Sprintf("Error trying to find bids by auctionId %s", auctionId), err)
		return nil, internal_error.NewInternalServerError(fmt.Sprintf("Error trying to find bids by auctionId %s", auctionId))
	}

	var bidEntitiesMongo []BidEntityMongo
	if err := cursor.All(ctx, &bidEntitiesMongo); err != nil {
		logger.Error(fmt.Sprintf("Error trying to find bids by auctionId %s", auctionId), err)
	}

	var bidEntities []bid_entity.Bid
	for _, bidEntityMongo := range bidEntitiesMongo {
		bidEntities = append(bidEntities, bid_entity.Bid{
			Id:        bidEntityMongo.Id,
			UserId:    bidEntityMongo.UserId,
			AuctionId: bidEntityMongo.AuctionId,
			Amount:    bidEntityMongo.Amount,
			Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
		})
	}

	return bidEntities, nil
}

func (br *BidRepository) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auctionId": auctionId}

	var bidEntityMongo BidEntityMongo
	opts := options.FindOne().SetSort(bson.D{{"amount", -1}})

	if err := br.Collection.FindOne(ctx, filter, opts).Decode(&bidEntityMongo); err != nil {
		logger.Error(fmt.Sprintf("Error trying to find bids by auctionId %s", auctionId), err)
		return nil, internal_error.NewInternalServerError(fmt.Sprintf("Error trying to find bids by auctionId %s", auctionId))
	}

	return &bid_entity.Bid{
		Id:        bidEntityMongo.Id,
		UserId:    bidEntityMongo.UserId,
		AuctionId: bidEntityMongo.AuctionId,
		Amount:    bidEntityMongo.Amount,
		Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
	}, nil
}

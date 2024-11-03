package bid

import (
	"context"
	"sync"

	"github.com/yodalis/golang/labs/auction_go/config/logger"
	"github.com/yodalis/golang/labs/auction_go/internal/entity/auction_entity"
	"github.com/yodalis/golang/labs/auction_go/internal/entity/bid_entity"
	"github.com/yodalis/golang/labs/auction_go/internal/infra/database/auction"
	"github.com/yodalis/golang/labs/auction_go/internal/internal_error"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidEntityMongo struct {
	Id        string  `bson:"_id"`
	UserId    string  `bson:"user_id"`
	AuctionId string  `bson:"auction_id"`
	Amount    float64 `bson:"amount"`
	Timestamp int64   `bson:"timestamp"`
}

type BidRepository struct {
	Collection        *mongo.Collection
	AuctionRepository *auction.AuctionRepository
}

func NewBidRepository(db *mongo.Database, auctionRepository *auction.AuctionRepository) *BidRepository {
	return &BidRepository{
		Collection:        db.Collection("bids"),
		AuctionRepository: auctionRepository,
	}
}

func (br *BidRepository) CreateBid(ctx context.Context, bidEntities []bid_entity.Bid) *internal_error.InternalError {
	var wg sync.WaitGroup

	for _, bid := range bidEntities {
		wg.Add(1)

		go func(bidValue bid_entity.Bid) {
			defer wg.Done()

			auctionEntity, err := br.AuctionRepository.FindAuctionById(ctx, bidValue.AuctionId)
			if err != nil {
				logger.Error("Error trying to find auction id", err)
				return
			}

			if auctionEntity.Status != auction_entity.Active {
				return
			}

			bidEntityMongo := &BidEntityMongo{
				Id:        bidValue.Id,
				UserId:    bid.UserId,
				AuctionId: bidValue.AuctionId,
				Amount:    bid.Amount,
				Timestamp: bidValue.Timestamp.Unix(),
			}

			if _, err := br.Collection.InsertOne(ctx, bidEntityMongo); err != nil {
				logger.Error("Error trying to insert  id", err)
				return
			}

		}(bid)
	}

	wg.Wait()

	return nil
}

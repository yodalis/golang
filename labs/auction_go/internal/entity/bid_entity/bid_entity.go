package bid_entity

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yodalis/golang/labs/auction_go/internal/internal_error"
)

type Bid struct {
	Id        string
	UserId    string
	AuctionId string
	Amount    float64
	Timestamp time.Time
}

func CreateBid(userId, auctionId string, amount float64) (*Bid, *internal_error.InternalError) {
	bid := &Bid{
		Id:        uuid.New().String(),
		UserId:    userId,
		AuctionId: auctionId,
		Amount:    amount,
		Timestamp: time.Now(),
	}

	if err := bid.Validate(); err != nil {
		return nil, err
	}

	return bid, nil
}

func (b *Bid) Validate() *internal_error.InternalError {
	if err := uuid.Validate(b.UserId); err != nil {
		return internal_error.NewBadRequestError("User ID is not a valid id")
	} else if err := uuid.Validate(b.AuctionId); err != nil {
		return internal_error.NewBadRequestError("Auction ID is not a valid id")
	} else if b.Amount <= 0 {
		return internal_error.NewBadRequestError("Amount is not a valid value")
	}

	return nil
}

type BidEntityRepository interface {
	CreateBid(ctx context.Context, bidEntities []Bid) *internal_error.InternalError
	FindBidByAuctionId(ctx context.Context, auctionId string) ([]Bid, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*Bid, *internal_error.InternalError)
}

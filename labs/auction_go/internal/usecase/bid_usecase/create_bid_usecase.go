package bid_usecase

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/yodalis/golang/labs/auction_go/config/logger"
	"github.com/yodalis/golang/labs/auction_go/internal/entity/bid_entity"
	"github.com/yodalis/golang/labs/auction_go/internal/internal_error"
)

var bidBatch []bid_entity.Bid

type BidInputDTO struct {
	UserId    string  `json:"user_id"`
	AuctionId string  `json:"auction_id"`
	Amount    float64 `json:"amount"`
}

type BidOutputDTO struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	AuctionId string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type BidUseCase struct {
	BidRepository bid_entity.BidEntityRepository

	timer               *time.Timer
	maxBatchSize        int
	batchInsertInterval time.Duration
	bidChannel          chan bid_entity.Bid
}

type BidUseCaseInterface interface {
	CreateBid(ctx context.Context, bidInputDTO BidInputDTO) *internal_error.InternalError
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDTO, *internal_error.InternalError)
	FindBidByAuctionId(ctx context.Context, auctionId string) ([]BidOutputDTO, *internal_error.InternalError)
}

func NewBidUseCase(bidRepository bid_entity.BidEntityRepository) BidUseCaseInterface {
	maxSizeInterval := getMaxBatchSizeInterval()
	maxBatchSize := getMaxBatchSize()
	bidUseCase := &BidUseCase{
		BidRepository:       bidRepository,
		maxBatchSize:        maxBatchSize,
		batchInsertInterval: maxSizeInterval,
		timer:               time.NewTimer(maxSizeInterval),
		bidChannel:          make(chan bid_entity.Bid, maxBatchSize),
	}

	bidUseCase.triggerCreateRoutine(context.Background())

	return bidUseCase
}

func (bd *BidUseCase) triggerCreateRoutine(ctx context.Context) {
	go func() {
		defer close(bd.bidChannel)

		for {
			select {
			case bidEntinty, ok := <-bd.bidChannel:
				if !ok {
					if len(bidBatch) > 0 {
						if err := bd.BidRepository.CreateBid(ctx, bidBatch); err != nil {
							logger.Error("error trying to process bid batch list", err)
						}
					}
					return
				}

				bidBatch = append(bidBatch, bidEntinty)
				if len(bidBatch) > bd.maxBatchSize {
					if err := bd.BidRepository.CreateBid(ctx, bidBatch); err != nil {
						logger.Error("error trying to process bid batch list", err)
					}

					bidBatch = nil
					bd.timer.Reset(bd.batchInsertInterval)
				}
			case <-bd.timer.C:
				if err := bd.BidRepository.CreateBid(ctx, bidBatch); err != nil {
					logger.Error("error trying to process bid batch list", err)
				}

				bidBatch = nil
				bd.timer.Reset(bd.batchInsertInterval)
			}

		}
	}()
}

func (bd *BidUseCase) CreateBid(ctx context.Context, bidInputDTO BidInputDTO) *internal_error.InternalError {
	bidEntity, err := bid_entity.CreateBid(bidInputDTO.UserId, bidInputDTO.AuctionId, bidInputDTO.Amount)
	if err != nil {
		return err
	}

	bd.bidChannel <- *bidEntity

	return nil
}

func getMaxBatchSizeInterval() time.Duration {
	batchInsertInterval := os.Getenv("BATCH_INSERT_INTERVAL")

	duration, err := time.ParseDuration(batchInsertInterval)
	if err != nil {
		return 3 * time.Minute
	}

	return duration
}

func getMaxBatchSize() int {
	value, err := strconv.Atoi(os.Getenv("MAX_BATCH_SIZE"))
	if err != nil {
		return 5
	}

	return value
}

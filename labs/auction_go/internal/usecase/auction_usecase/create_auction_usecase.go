package auction_usecase

import (
	"context"
	"time"

	"github.com/yodalis/golang/labs/auction_go/internal/entity/auction_entity"
	"github.com/yodalis/golang/labs/auction_go/internal/entity/bid_entity"
	"github.com/yodalis/golang/labs/auction_go/internal/internal_error"
	"github.com/yodalis/golang/labs/auction_go/internal/usecase/bid_usecase"
)

type AuctionInputDTO struct {
	ProductName string           `json:"product_name" binding:"required,min=1"`
	Category    string           `json:"category" binding:"required,min=2"`
	Description string           `json:"description" binding:"required,min=10,max=200"`
	Condition   ProductCondition `json:"condition"`
}

type AuctionOutputDTO struct {
	Id          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type ProductCondition int64
type AuctionStatus int64

type AuctionUseCase struct {
	auctionRepositoryInterface auction_entity.AuctionRepositoryInterface
	bidRepositoryInterface     bid_entity.BidEntityRepository
}

type WinningInfoOutputDTO struct {
	Auction AuctionOutputDTO          `json:"auction"`
	Bid     *bid_usecase.BidOutputDTO `json:"bid,omitempty"`
}

type AuctionUseCaseInterface interface {
	CreateAuction(
		ctx context.Context,
		auctionInput AuctionInputDTO) *internal_error.InternalError

	FindAuctionById(
		ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError)

	FindAuctions(
		ctx context.Context,
		status AuctionStatus,
		category, productName string) ([]AuctionOutputDTO, *internal_error.InternalError)

	FindWinningBidByAuctionId(
		ctx context.Context,
		auctionId string) (*WinningInfoOutputDTO, *internal_error.InternalError)
}

func NewAuctionUseCase(
	auctionRepositoryInterface auction_entity.AuctionRepositoryInterface,
	bidRepositoryInterface bid_entity.BidEntityRepository) AuctionUseCaseInterface {
	return &AuctionUseCase{
		auctionRepositoryInterface: auctionRepositoryInterface,
		bidRepositoryInterface:     bidRepositoryInterface,
	}
}

func (au *AuctionUseCase) CreateAuction(ctx context.Context, auctionInput AuctionInputDTO) *internal_error.InternalError {
	auction, err := auction_entity.CreateAuction(auctionInput.ProductName, auctionInput.Category, auctionInput.Description, auction_entity.ProductCondition(auctionInput.Condition))
	if err != nil {
		return err
	}

	if err := au.auctionRepositoryInterface.CreateAuction(ctx, auction); err != nil {
		return err
	}

	return nil
}

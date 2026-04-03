package handler

import (
	"context"
	"errors"

	"github.com/DenisMekh/mini-transfer-system/account-svc/internal/domain"
	pb "github.com/DenisMekh/mini-transfer-system/gen/go/account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/DenisMekh/mini-transfer-system/account-svc/internal/service"
)

type AccountHandler struct {
	pb.UnimplementedAccountServiceServer
	serv service.AccountService
}

func NewAccountHandler(service service.AccountService) *AccountHandler {
	return &AccountHandler{
		serv: service,
	}
}

// CreateAccount method for creating an account
func (h *AccountHandler) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.Account, error) {
	name := req.GetName()
	userID := req.GetUserId()
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	if userID == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}
	account := &domain.Account{
		Name:   name,
		UserID: userID,
	}
	err := h.serv.Create(ctx, account)
	if err != nil {
		if errors.Is(err, domain.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "account already exists")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return toPbAccount(account), nil
}

// GetAccount method for getting an account from accountID
func (h *AccountHandler) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.Account, error) {
	accountID := req.GetAccountId()
	if accountID == "" {
		return nil, status.Error(codes.InvalidArgument, "account id is required")
	}
	account, err := h.serv.GetByID(ctx, accountID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "account not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return toPbAccount(account), nil
}

// UpdateBalance method for updating balance for user with accountID
func (h *AccountHandler) UpdateBalance(ctx context.Context, req *pb.UpdateBalanceRequest) (*pb.Account, error) {
	accountID := req.GetAccountId()
	amount := req.GetAmount()
	if accountID == "" {
		return nil, status.Error(codes.InvalidArgument, "account id is required")
	}
	if amount == 0 {
		return nil, status.Error(codes.InvalidArgument, "amount is required")
	}
	account, err := h.serv.UpdateBalance(ctx, accountID, amount)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "account not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return toPbAccount(account), nil
}

// toPbAccount private function for mapping account into grpc pbAccount
func toPbAccount(a *domain.Account) *pb.Account {
	return &pb.Account{
		AccountId: a.AccountID,
		UserId:    a.UserID,
		Name:      a.Name,
		Balance:   a.Balance,
		CreatedAt: timestamppb.New(a.CreatedAt),
	}
}

package grpcserver

import (
	"context"
	"time"

	loansv1 "loans-item-go/gen/loans/v1"
	"loans-item-go/model"
	loansvc "loans-item-go/service/loan"
)

type LoanHandler struct {
	loansv1.UnimplementedLoanServiceServer
	svc loansvc.Service
}

func NewLoanHandler(svc loansvc.Service) *LoanHandler {
	return &LoanHandler{svc: svc}
}

func (h *LoanHandler) CreateLoan(ctx context.Context, req *loansv1.CreateLoanRequest) (*loansv1.CreateLoanResponse, error) {
	created := h.svc.Create(ctx, model.Loan{
		ItemID:     req.ItemId,
		BorrowerID: req.BorrowerId,
	})
	// Create returns the row without Item/Borrower preloaded or derived
	// Status; re-fetch so the response carries the full loan.
	loan := h.svc.FindById(ctx, int(created.ID))
	return &loansv1.CreateLoanResponse{Loan: toProtoLoan(loan)}, nil
}

func (h *LoanHandler) GetLoan(ctx context.Context, req *loansv1.GetLoanRequest) (*loansv1.GetLoanResponse, error) {
	loan := h.svc.FindById(ctx, int(req.Id))
	return &loansv1.GetLoanResponse{Loan: toProtoLoan(loan)}, nil
}

func (h *LoanHandler) ListLoans(ctx context.Context, req *loansv1.ListLoansRequest) (*loansv1.ListLoansResponse, error) {
	page, pageSize := pageOrDefault(req.Page)
	loans, total := h.svc.FindAll(ctx, page, pageSize)
	return &loansv1.ListLoansResponse{
		Loans:    toProtoLoans(loans),
		Page:     int32(page),
		PageSize: int32(pageSize),
		Total:    total,
	}, nil
}

func (h *LoanHandler) ListLoansByBorrower(ctx context.Context, req *loansv1.ListLoansByBorrowerRequest) (*loansv1.ListLoansByBorrowerResponse, error) {
	page, pageSize := pageOrDefault(req.Page)
	loans, total := h.svc.FindByBorrowerID(ctx, int(req.BorrowerId), page, pageSize)
	return &loansv1.ListLoansByBorrowerResponse{
		Loans:    toProtoLoans(loans),
		Page:     int32(page),
		PageSize: int32(pageSize),
		Total:    total,
	}, nil
}

func (h *LoanHandler) UpdateLoan(ctx context.Context, req *loansv1.UpdateLoanRequest) (*loansv1.UpdateLoanResponse, error) {
	returnedAt := time.Now()
	if req.ReturnedAt != nil {
		returnedAt = req.ReturnedAt.AsTime()
	}
	h.svc.Update(ctx, model.Loan{ID: req.Id, ReturnedAt: &returnedAt})
	loan := h.svc.FindById(ctx, int(req.Id))
	return &loansv1.UpdateLoanResponse{Loan: toProtoLoan(loan)}, nil
}

func (h *LoanHandler) DeleteLoan(ctx context.Context, req *loansv1.DeleteLoanRequest) (*loansv1.DeleteLoanResponse, error) {
	h.svc.Delete(ctx, model.Loan{ID: req.Id})
	return &loansv1.DeleteLoanResponse{Message: "loan deleted"}, nil
}

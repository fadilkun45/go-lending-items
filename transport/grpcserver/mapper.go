package grpcserver

import (
	"time"

	loansv1 "loans-item-go/gen/loans/v1"
	"loans-item-go/model"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func toProtoUser(user model.User) *loansv1.User {
	return &loansv1.User{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}

func toProtoUsers(users []model.User) []*loansv1.User {
	result := make([]*loansv1.User, len(users))
	for i, u := range users {
		result[i] = toProtoUser(u)
	}
	return result
}

func toProtoLoan(loan model.Loan) *loansv1.Loan {
	return &loansv1.Loan{
		Id: loan.ID,
		Item: &loansv1.ItemSummary{
			Id:   loan.Item.ID,
			Name: loan.Item.Name,
		},
		Borrower: &loansv1.User{
			Id:    loan.Borrower.Id,
			Name:  loan.Borrower.Name,
			Email: loan.Borrower.Email,
		},
		BorrowedAt: timestamppb.New(loan.BorrowedAt),
		ReturnedAt: toProtoTime(loan.ReturnedAt),
		Status:     loan.Status,
	}
}

func toProtoLoans(loans []model.Loan) []*loansv1.Loan {
	result := make([]*loansv1.Loan, len(loans))
	for i, l := range loans {
		result[i] = toProtoLoan(l)
	}
	return result
}

func toProtoTime(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

// pageOrDefault applies the same defaults as the REST controllers.
func pageOrDefault(page *loansv1.PageRequest) (int, int) {
	p, size := 0, 0
	if page != nil {
		p, size = int(page.Page), int(page.PageSize)
	}
	if p <= 0 {
		p = 1
	}
	if size <= 0 {
		size = 10
	}
	return p, size
}

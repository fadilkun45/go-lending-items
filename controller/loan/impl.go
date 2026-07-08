package loanctrl

import (
	"loans-item-go/helper"
	"loans-item-go/model"
	loansvc "loans-item-go/service/loan"
	"net/http"
	"strconv"
	"time"
)

type ControllerImpl struct {
	LoanService loansvc.Service
}

func NewController(loanService loansvc.Service) Controller {
	return &ControllerImpl{LoanService: loanService}
}

type CreateLoanRequest struct {
	ItemID     int64 `json:"item_id" validate:"required"`
	BorrowerID int64 `json:"borrower_id" validate:"required"`
}

type UpdateLoanRequest struct {
	ReturnedAt *time.Time `json:"returned_at"`
}

type ItemSummaryResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type UserSummaryResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoanResponse struct {
	ID         int64               `json:"id"`
	Item       ItemSummaryResponse `json:"item"`
	Borrower   UserSummaryResponse `json:"borrower"`
	BorrowedAt time.Time           `json:"borrowed_at"`
	ReturnedAt *time.Time          `json:"returned_at"`
	Status     string              `json:"status"`
}

func toLoanResponse(loan model.Loan) LoanResponse {
	return LoanResponse{
		ID: loan.ID,
		Item: ItemSummaryResponse{
			ID:   loan.Item.ID,
			Name: loan.Item.Name,
		},
		Borrower: UserSummaryResponse{
			ID:    loan.Borrower.Id,
			Name:  loan.Borrower.Name,
			Email: loan.Borrower.Email,
		},
		BorrowedAt: loan.BorrowedAt,
		ReturnedAt: loan.ReturnedAt,
		Status:     loan.Status,
	}
}

func toLoanResponses(loans []model.Loan) []LoanResponse {
	responses := make([]LoanResponse, len(loans))
	for i, l := range loans {
		responses[i] = toLoanResponse(l)
	}
	return responses
}

// @Summary Create loan (borrow item)
// @Tags loans
// @Accept json
// @Produce json
// @Param body body CreateLoanRequest true "Loan data"
// @Success 201 {object} helper.WebResponse
// @Failure 400 {object} helper.WebResponse
// @Security BearerAuth
// @Router /api/loans [post]
func (h *ControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	var req CreateLoanRequest
	if err := helper.DecodeRequest(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, helper.FormatDecodeError(err))
		return
	}
	if errs := helper.ValidateStruct(req); errs != nil {
		helper.WriteError(w, http.StatusBadRequest, errs[0])
		return
	}
	h.LoanService.Create(r.Context(), model.Loan{ItemID: req.ItemID, BorrowerID: req.BorrowerID})
	helper.WriteResponse(w, http.StatusCreated, "success borrowed items")
}

// @Summary Get loan by ID
// @Tags loans
// @Produce json
// @Param id path int true "Loan ID"
// @Success 200 {object} LoanResponse
// @Failure 404 {object} helper.WebResponse
// @Security BearerAuth
// @Router /api/loans/{id} [get]
func (h *ControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	helper.WriteResponse(w, http.StatusOK, toLoanResponse(h.LoanService.FindById(r.Context(), id)))
}

// @Summary Get all loans
// @Tags loans
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} helper.PaginatedWebResponse
// @Security BearerAuth
// @Router /api/loans [get]
func (h *ControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	loans, total := h.LoanService.FindAll(r.Context(), page, pageSize)
	helper.WritePaginatedResponse(w, http.StatusOK, toLoanResponses(loans), page, pageSize, total)
}

// @Summary Get loans by borrower ID
// @Tags loans
// @Produce json
// @Param borrower_id path int true "Borrower ID"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} helper.PaginatedWebResponse
// @Security BearerAuth
// @Router /api/loans/borrower/{borrower_id} [get]
func (h *ControllerImpl) FindByBorrowerID(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	borrowerID, err := strconv.Atoi(r.PathValue("borrower_id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid borrower_id")
		return
	}
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	loans, total := h.LoanService.FindByBorrowerID(r.Context(), borrowerID, page, pageSize)
	helper.WritePaginatedResponse(w, http.StatusOK, toLoanResponses(loans), page, pageSize, total)
}

// @Summary Return loan
// @Tags loans
// @Accept json
// @Produce json
// @Param id path int true "Loan ID"
// @Param body body UpdateLoanRequest false "Returned at (optional, defaults to now)"
// @Success 200 {object} helper.WebResponse
// @Failure 404 {object} helper.WebResponse
// @Security BearerAuth
// @Router /api/loans/{id} [put]
func (h *ControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	var req UpdateLoanRequest
	if err := helper.DecodeRequest(r, &req); err != nil && err.Error() != "EOF" {
		helper.WriteError(w, http.StatusBadRequest, helper.FormatDecodeError(err))
		return
	}
	now := time.Now()
	loan := model.Loan{ID: int64(id), ReturnedAt: &now}
	if req.ReturnedAt != nil {
		loan.ReturnedAt = req.ReturnedAt
	}
	h.LoanService.Update(r.Context(), loan)
	helper.WriteResponse(w, http.StatusOK, "success returned items")
}

// @Summary Delete loan
// @Tags loans
// @Produce json
// @Param id path int true "Loan ID"
// @Success 200 {object} helper.WebResponse
// @Failure 404 {object} helper.WebResponse
// @Security BearerAuth
// @Router /api/loans/{id} [delete]
func (h *ControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	h.LoanService.Delete(r.Context(), model.Loan{ID: int64(id)})
	helper.WriteResponse(w, http.StatusOK, "loan deleted")
}

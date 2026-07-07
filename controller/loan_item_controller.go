package controller

import (
	"loans-item-go/helper"
	"loans-item-go/model"
	"loans-item-go/service"
	"net/http"
	"strconv"
)

type LoanItemController struct {
	LoanItemService service.LoanItemService
}

func NewLoanItemController(loanItemService service.LoanItemService) *LoanItemController {
	return &LoanItemController{LoanItemService: loanItemService}
}

type CreateLoanItemRequest struct {
	ItemID int64 `json:"item_id" validate:"required"`
	LoanID int64 `json:"loan_id" validate:"required"`
}

type UpdateLoanItemRequest struct {
	ItemID int64 `json:"item_id" validate:"required"`
	LoanID int64 `json:"loan_id" validate:"required"`
}

func (h *LoanItemController) Create(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	var req CreateLoanItemRequest
	if err := helper.DecodeRequest(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, helper.FormatDecodeError(err))
		return
	}

	if errs := helper.ValidateStruct(req); errs != nil {
		helper.WriteError(w, http.StatusBadRequest, errs[0])
		return
	}

	loanItem := model.LoanItem{
		ItemId: req.ItemID,
		LoanId: req.LoanID,
	}
	result := h.LoanItemService.Create(r.Context(), loanItem)
	helper.WriteResponse(w, http.StatusCreated, result)
}

func (h *LoanItemController) FindById(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	loanItem := h.LoanItemService.FindById(r.Context(), id)
	helper.WriteResponse(w, http.StatusOK, loanItem)
}

func (h *LoanItemController) FindByLoanID(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	loanID, err := strconv.Atoi(r.PathValue("loan_id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid loan_id")
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

	loanItems, total := h.LoanItemService.FindByLoanID(r.Context(), loanID, page, pageSize)
	helper.WritePaginatedResponse(w, http.StatusOK, loanItems, page, pageSize, total)
}

func (h *LoanItemController) FindAll(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	loanItems, total := h.LoanItemService.FindAll(r.Context(), page, pageSize)
	helper.WritePaginatedResponse(w, http.StatusOK, loanItems, page, pageSize, total)
}

func (h *LoanItemController) Update(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req UpdateLoanItemRequest
	if err := helper.DecodeRequest(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, helper.FormatDecodeError(err))
		return
	}

	if errs := helper.ValidateStruct(req); errs != nil {
		helper.WriteError(w, http.StatusBadRequest, errs[0])
		return
	}

	loanItem := model.LoanItem{
		Id:     int64(id),
		ItemId: req.ItemID,
		LoanId: req.LoanID,
	}
	result := h.LoanItemService.Update(r.Context(), loanItem)
	helper.WriteResponse(w, http.StatusOK, result)
}

func (h *LoanItemController) Delete(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	h.LoanItemService.Delete(r.Context(), model.LoanItem{Id: int64(id)})
	helper.WriteResponse(w, http.StatusOK, "loan item deleted")
}

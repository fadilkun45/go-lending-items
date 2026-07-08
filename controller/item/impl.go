package itemctrl

import (
	"loans-item-go/helper"
	"loans-item-go/model"
	"loans-item-go/service/item"
	"net/http"
	"strconv"
)

type ControllerImpl struct {
	ItemService itemsvc.Service
}

func NewController(itemService itemsvc.Service) Controller {
	return &ControllerImpl{ItemService: itemService}
}

type CreateItemRequest struct {
	Name       string `json:"name" validate:"required"`
	CategoryID int64  `json:"category_id" validate:"required"`
	SerialNo   string `json:"serial_no"`
	Condition  string `json:"condition"`
	Status     string `json:"status"`
	OwnerID    int64  `json:"owner_id" validate:"required"`
}

type UpdateItemRequest struct {
	Name      string `json:"name" validate:"required"`
	Condition string `json:"condition" validate:"required"`
	Status    string `json:"status" validate:"required"`
}

// @Summary Create item
// @Tags items
// @Accept json
// @Produce json
// @Param body body CreateItemRequest true "Item data"
// @Success 201 {object} model.Item
// @Failure 400 {object} helper.WebResponse
// @Router /api/items [post]
func (h *ControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	var req CreateItemRequest
	if err := helper.DecodeRequest(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, helper.FormatDecodeError(err))
		return
	}
	if errs := helper.ValidateStruct(req); errs != nil {
		helper.WriteError(w, http.StatusBadRequest, errs[0])
		return
	}
	var serialNo *string
	if req.SerialNo != "" {
		serialNo = &req.SerialNo
	}
	result := h.ItemService.Create(r.Context(), model.Item{
		Name:       req.Name,
		CategoryID: req.CategoryID,
		SerialNo:   serialNo,
		Condition:  req.Condition,
		Status:     req.Status,
		OwnerID:    req.OwnerID,
	})
	helper.WriteResponse(w, http.StatusCreated, result)
}

// @Summary Get item by ID
// @Tags items
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} model.Item
// @Failure 404 {object} helper.WebResponse
// @Router /api/items/{id} [get]
func (h *ControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	helper.WriteResponse(w, http.StatusOK, h.ItemService.FindById(r.Context(), id))
}

// @Summary Get all items
// @Tags items
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} helper.PaginatedWebResponse
// @Router /api/items [get]
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
	items, total := h.ItemService.FindAll(r.Context(), page, pageSize)
	helper.WritePaginatedResponse(w, http.StatusOK, items, page, pageSize, total)
}

// @Summary Get items by owner
// @Tags items
// @Produce json
// @Param ownerId path int true "Owner ID"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} helper.PaginatedWebResponse
// @Router /api/items/owner/{ownerId} [get]
func (h *ControllerImpl) FindByOwner(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	ownerId, err := strconv.Atoi(r.PathValue("ownerId"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid ownerId")
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
	items, total := h.ItemService.FindByOwner(r.Context(), ownerId, page, pageSize)
	helper.WritePaginatedResponse(w, http.StatusOK, items, page, pageSize, total)
}

// @Summary Update item
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param body body UpdateItemRequest true "Update data"
// @Success 200 {object} model.Item
// @Failure 400 {object} helper.WebResponse
// @Router /api/items/{id} [put]
func (h *ControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	var req UpdateItemRequest
	if err := helper.DecodeRequest(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, helper.FormatDecodeError(err))
		return
	}
	if errs := helper.ValidateStruct(req); errs != nil {
		helper.WriteError(w, http.StatusBadRequest, errs[0])
		return
	}
	result := h.ItemService.Update(r.Context(), model.Item{ID: int64(id), Name: req.Name, Condition: req.Condition, Status: req.Status})
	helper.WriteResponse(w, http.StatusOK, result)
}

// @Summary Delete item
// @Tags items
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} helper.WebResponse
// @Failure 404 {object} helper.WebResponse
// @Router /api/items/{id} [delete]
func (h *ControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	h.ItemService.Delete(r.Context(), model.Item{ID: int64(id)})
	helper.WriteResponse(w, http.StatusOK, "item deleted")
}

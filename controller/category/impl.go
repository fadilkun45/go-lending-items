package categoryctrl

import (
	"loans-item-go/helper"
	"loans-item-go/model"
	"loans-item-go/service/category"
	"net/http"
	"strconv"
)

type ControllerImpl struct {
	CategoryService categorysvc.Service
}

func NewController(categoryService categorysvc.Service) Controller {
	return &ControllerImpl{CategoryService: categoryService}
}

type CreateCategoryRequest struct {
	Name    string `json:"name" validate:"required"`
	OwnerID uint64 `json:"owner_id" validate:"required"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

// @Summary Create category
// @Tags categories
// @Accept json
// @Produce json
// @Param body body CreateCategoryRequest true "Category data"
// @Success 201 {object} model.Category
// @Failure 400 {object} helper.WebResponse
// @Security BearerAuth
// @Router /api/categories [post]
func (h *ControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	var req CreateCategoryRequest
	if err := helper.DecodeRequest(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, helper.FormatDecodeError(err))
		return
	}
	if errs := helper.ValidateStruct(req); errs != nil {
		helper.WriteError(w, http.StatusBadRequest, errs[0])
		return
	}
	result := h.CategoryService.Create(r.Context(), model.Category{Name: req.Name, OwnerId: req.OwnerID})
	helper.WriteResponse(w, http.StatusCreated, result)
}

// @Summary Get category by ID
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} model.Category
// @Failure 404 {object} helper.WebResponse
// @Security BearerAuth
// @Router /api/categories/{id} [get]
func (h *ControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	helper.WriteResponse(w, http.StatusOK, h.CategoryService.FindById(r.Context(), id))
}

// @Summary Get all categories
// @Tags categories
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} helper.PaginatedWebResponse
// @Security BearerAuth
// @Router /api/categories [get]
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
	categories, total := h.CategoryService.FindAll(r.Context(), page, pageSize)
	helper.WritePaginatedResponse(w, http.StatusOK, categories, page, pageSize, total)
}

// @Summary Update category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param body body UpdateCategoryRequest true "Update data"
// @Success 200 {object} model.Category
// @Failure 400 {object} helper.WebResponse
// @Security BearerAuth
// @Router /api/categories/{id} [put]
func (h *ControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	var req UpdateCategoryRequest
	if err := helper.DecodeRequest(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, helper.FormatDecodeError(err))
		return
	}
	if errs := helper.ValidateStruct(req); errs != nil {
		helper.WriteError(w, http.StatusBadRequest, errs[0])
		return
	}
	result := h.CategoryService.Update(r.Context(), model.Category{ID: int64(id), Name: req.Name})
	helper.WriteResponse(w, http.StatusOK, result)
}

// @Summary Delete category
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} helper.WebResponse
// @Failure 404 {object} helper.WebResponse
// @Security BearerAuth
// @Router /api/categories/{id} [delete]
func (h *ControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	h.CategoryService.Delete(r.Context(), model.Category{ID: int64(id)})
	helper.WriteResponse(w, http.StatusOK, "category deleted")
}

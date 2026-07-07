package controller

import (
	"loans-item-go/helper"
	"loans-item-go/model"
	"loans-item-go/service"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

// @Summary Register user
// @Tags users
// @Accept json
// @Produce json
// @Param body body RegisterRequest true "Register data"
// @Success 201 {object} model.User
// @Failure 400 {object} helper.WebResponse
// @Router /api/users/register [post]

func (h *UserController) Register(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	var req RegisterRequest
	if err := helper.DecodeRequest(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, helper.FormatDecodeError(err))
		return
	}

	if errs := helper.ValidateStruct(req); errs != nil {
		helper.WriteError(w, http.StatusBadRequest, errs[0])
		return
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	result := h.UserService.Register(r.Context(), user)
	helper.WriteResponse(w, http.StatusCreated, result)
}

// @Summary Login user
// @Tags users
// @Accept json
// @Produce json
// @Param body body LoginRequest true "Login data"
// @Success 200 {object} model.User
// @Failure 400 {object} helper.WebResponse
// @Router /api/users/login [post]
func (h *UserController) Login(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	var req LoginRequest
	if err := helper.DecodeRequest(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, helper.FormatDecodeError(err))
		return
	}

	if errs := helper.ValidateStruct(req); errs != nil {
		helper.WriteError(w, http.StatusBadRequest, errs[0])
		return
	}

	user := h.UserService.Login(r.Context(), req.Email, req.Password)
	helper.WriteResponse(w, http.StatusOK, user)
}

// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 404 {object} helper.WebResponse
// @Router /api/users/{id} [get]
func (h *UserController) FindById(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	user := h.UserService.FindById(r.Context(), id)
	helper.WriteResponse(w, http.StatusOK, user)
}

// @Summary Get all users
// @Tags users
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} helper.PaginatedWebResponse
// @Router /api/users [get]
func (h *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	users, total := h.UserService.FindAll(r.Context(), page, pageSize)
	helper.WritePaginatedResponse(w, http.StatusOK, users, page, pageSize, total)
}

// @Summary Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param body body UpdateUserRequest true "Update data"
// @Success 200 {object} model.User
// @Failure 400 {object} helper.WebResponse
// @Router /api/users/{id} [put]
func (h *UserController) Update(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req UpdateUserRequest
	if err := helper.DecodeRequest(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, helper.FormatDecodeError(err))
		return
	}

	if errs := helper.ValidateStruct(req); errs != nil {
		helper.WriteError(w, http.StatusBadRequest, errs[0])
		return
	}

	user := model.User{
		Id:    int64(id),
		Name:  req.Name,
		Email: req.Email,
	}
	result := h.UserService.Update(r.Context(), user)
	helper.WriteResponse(w, http.StatusOK, result)
}

// @Summary Delete user
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} helper.WebResponse
// @Failure 404 {object} helper.WebResponse
// @Router /api/users/{id} [delete]
func (h *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	h.UserService.Delete(r.Context(), model.User{Id: int64(id)})
	helper.WriteResponse(w, http.StatusOK, "user deleted")
}

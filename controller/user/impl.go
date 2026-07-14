package userctrl

import (
	"loans-item-go/helper"
	"loans-item-go/middleware"
	"loans-item-go/model"
	"loans-item-go/service/user"
	"net/http"
	"strconv"
)

type ControllerImpl struct {
	UserService usersvc.Service
}

func NewController(userService usersvc.Service) Controller {
	return &ControllerImpl{UserService: userService}
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

type LoginResponse struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
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
func (h *ControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
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
	result := h.UserService.Register(r.Context(), model.User{Name: req.Name, Email: req.Email, Password: req.Password})
	helper.WriteResponse(w, http.StatusCreated, result)
}

// @Summary Login user
// @Tags users
// @Accept json
// @Produce json
// @Param body body LoginRequest true "Login data"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} helper.WebResponse
// @Router /api/users/login [post]
func (h *ControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
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
	user, token := h.UserService.Login(r.Context(), req.Email, req.Password)
	helper.WriteResponse(w, http.StatusOK, LoginResponse{User: user, Token: token})
}

// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 404 {object} helper.WebResponse
// @Security BearerAuth
// @Router /api/users/{id} [get]
func (h *ControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	helper.WriteResponse(w, http.StatusOK, h.UserService.FindById(r.Context(), id))
}

// @Summary Get current user from JWT
// @Tags users
// @Produce json
// @Success 200 {object} model.User
// @Failure 401 {object} helper.WebResponse
// @Failure 404 {object} helper.WebResponse
// @Security BearerAuth
// @Router /api/users/me [get]
func (h *ControllerImpl) Me(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok || userID == 0 {
		helper.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	helper.WriteResponse(w, http.StatusOK, h.UserService.FindById(r.Context(), int(userID)))
}

// @Summary Get all users
// @Tags users
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} helper.PaginatedWebResponse
// @Security BearerAuth
// @Router /api/users [get]
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
// @Security BearerAuth
// @Router /api/users/{id} [put]
func (h *ControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
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
	result := h.UserService.Update(r.Context(), model.User{Id: int64(id), Name: req.Name, Email: req.Email})
	helper.WriteResponse(w, http.StatusOK, result)
}

// @Summary Delete user
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} helper.WebResponse
// @Failure 404 {object} helper.WebResponse
// @Security BearerAuth
// @Router /api/users/{id} [delete]
func (h *ControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	defer helper.RecoverError(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	h.UserService.Delete(r.Context(), model.User{Id: int64(id)})
	helper.WriteResponse(w, http.StatusOK, "user deleted")
}

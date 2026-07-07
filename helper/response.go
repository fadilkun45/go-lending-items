package helper

import (
	"encoding/json"
	"net/http"
)

type WebResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func WriteResponse(w http.ResponseWriter, code int, data any) {
	status := "OK"
	if code >= 400 {
		status = "Error"
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(WebResponse{
		Code:   code,
		Status: status,
		Data:   data,
	})
}

func WriteError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(WebResponse{
		Code:    code,
		Status:  "Error",
		Message: message,
	})
}

type PaginatedWebResponse struct {
	Code     int    `json:"code"`
	Status   string `json:"status"`
	Data     any    `json:"data,omitempty"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Total    int64  `json:"total"`
}

func WritePaginatedResponse(w http.ResponseWriter, code int, data any, page int, pageSize int, total int64) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(PaginatedWebResponse{
		Code:     code,
		Status:   "OK",
		Data:     data,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	})
}

func RecoverError(w http.ResponseWriter) {
	if r := recover(); r != nil {
		switch v := r.(type) {
		case AppError:
			WriteError(w, v.Code, v.Message)
		case error:
			WriteError(w, http.StatusInternalServerError, v.Error())
		case string:
			WriteError(w, http.StatusInternalServerError, v)
		default:
			WriteError(w, http.StatusInternalServerError, "internal server error")
		}
	}
}

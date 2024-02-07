package errors

import "net/http"

type ApiError struct {
	Status    int    `json:"status"`
	ErrorType string `json:"error_type"`
	Details   string `json:"details"`
}

func NotFoundError(err error, includeDetails bool) *ApiError {
	details := ""
	if includeDetails {
		details = err.Error()
	}
	return &ApiError{
		Status:    http.StatusNotFound,
		ErrorType: "NOT FOUND",
		Details:   details,
	}
}

func InvalidRequest(err error, includeDetails bool) *ApiError {
	details := ""
	if includeDetails {
		details = err.Error()
	}
	return &ApiError{
		Status:    http.StatusBadRequest,
		ErrorType: "BAD REQUEST",
		Details:   details,
	}
}

func BadGateway(err error, includeDetails bool) *ApiError {
	details := ""
	if includeDetails {
		details = err.Error()
	}
	return &ApiError{
		Status:    http.StatusBadGateway,
		ErrorType: "BAD GATEWAY",
		Details:   details,
	}
}

func NotAuthorized(err error, includeDetails bool) *ApiError {
	details := ""
	if includeDetails {
		details = err.Error()
	}
	return &ApiError{
		Status:    http.StatusUnauthorized,
		ErrorType: "NOT AUTHORIZED",
		Details:   details,
	}
}

func Forbidden(err error, includeDetails bool) *ApiError {
	details := ""
	if includeDetails {
		details = err.Error()
	}
	return &ApiError{
		Status:    http.StatusForbidden,
		ErrorType: "FORBIDDEN",
		Details:   details,
	}
}

func Conflict(err error, includeDetails bool) *ApiError {
	details := ""
	if includeDetails {
		details = err.Error()
	}
	return &ApiError{
		Status:    http.StatusConflict,
		ErrorType: "CONFLICT",
		Details:   details,
	}
}

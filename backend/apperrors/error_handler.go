package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError
	if !errors.As(err, &appErr) {
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	log.Printf("error: %s\n", appErr)

	var statusCode int

	switch appErr.ErrCode {
	case NoData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}

package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/MatsuoTakuro/myapi-go-intermediate/api/contexts"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	var appErr *MyAppError
	if !errors.As(err, &appErr) {
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := contexts.GetTracdID(r.Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int

	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	case RequiredAuthHeader, Unauthorizated:
		statusCode = http.StatusUnauthorized
	case NotMatchUser:
		statusCode = http.StatusForbidden
	case CannotMakeValidator:
		statusCode = http.StatusInternalServerError
	case ResBodyEncodeFailed:
		statusCode = http.StatusInternalServerError
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(appErr)
}

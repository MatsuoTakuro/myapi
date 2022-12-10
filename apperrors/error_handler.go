package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/MatsuoTakuro/myapi-go-intermediate/api/middlewares"
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

	// TODO: fix it later. handlers should not depend on optional middlewares
	traceID := middlewares.GetTracdID(r.Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int

	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	case ResBodyEncodeFailed:
		statusCode = http.StatusInternalServerError
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(appErr)
}

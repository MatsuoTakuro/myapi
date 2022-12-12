package middlewares

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/MatsuoTakuro/myapi-go-intermediate/api/contexts"
	"github.com/MatsuoTakuro/myapi-go-intermediate/apperrors"
	"google.golang.org/api/idtoken"
)

var (
	googleCliID = os.Getenv("GOOGLE_CLIENT_ID")
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		authValues := strings.Split(auth, " ")
		if len(authValues) != 2 {
			err := apperrors.RequiredAuthHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, r, err)
			return
		}

		bearer, idToken := authValues[0], authValues[1]
		if bearer != "Bearer" || idToken == "" {
			err := apperrors.RequiredAuthHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, r, err)
			return
		}

		tokenValidator, err := idtoken.NewValidator(context.Background())
		if err != nil {
			err = apperrors.CannotMakeValidator.Wrap(err, "internal auth error")
			apperrors.ErrorHandler(w, r, err)
			return
		}

		p, err := tokenValidator.Validate(context.Background(), idToken, googleCliID)
		if err != nil {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, r, err)
			return
		}

		name, ok := p.Claims["name"]
		if !ok {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, r, err)
			return
		}

		r = contexts.SetUserName(r, name.(string))

		next.ServeHTTP(w, r)
	})
}

package contexts

import (
	"context"
	"net/http"
)

type userNameKey struct{}

func GetUserName(ctx context.Context) string {
	v := ctx.Value(userNameKey{})

	if un, ok := v.(string); ok {
		return un
	}

	return ""
}

func SetUserName(req *http.Request, name string) *http.Request {
	p := req.Context()

	ctx := context.WithValue(p, userNameKey{}, name)
	req = req.WithContext(ctx)

	return req
}

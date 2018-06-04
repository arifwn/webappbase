package auth

import (
	"context"
	"log"
	"net/http"
)

// user model definition
type User struct {
	username        string
	isAuthenticated bool
}

// context key definition
type contextKey string

func (c contextKey) String() string {
	return "auth context key " + string(c)
}

// assign user data associated with this request to context
func newContextWithUser(ctx context.Context, req *http.Request) context.Context {
	userData := User{
		username:        "test user",
		isAuthenticated: true,
	}

	return context.WithValue(ctx, contextKey("user"), userData)
}

// retrieve user data from context
func UserFromContext(ctx context.Context) User {
	user, ok := ctx.Value(contextKey("user")).(User)
	if !ok {
		return User{
			username:        "",
			isAuthenticated: false,
		}
	}
	return user
}

// User middleware: assign user data to context
func UserContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Println("UserContextMiddleware")
		ctx := newContextWithUser(req.Context(), req)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}

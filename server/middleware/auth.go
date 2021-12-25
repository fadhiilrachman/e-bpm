package middleware

import (
    "context"
    "net/http"
    "strings"

    controller "github.com/fadhiilrachman/e-bpm/controllers"
    utils "github.com/fadhiilrachman/e-bpm/utils"
)

var userCtxKey = &ContextKey{"admin", "username"}
type ContextKey struct {
    role string
    username string
}

func Auth() func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            token := r.Header.Get("authorization")
            
            // allow for anonymous user
            if token == "" {
                next.ServeHTTP(w, r)
                return
            }

            // if login, validate token
            splitToken := strings.Split(token, "Bearer ")
            tokenString := splitToken[1]
            data, err := utils.ParseToken(tokenString)

            if err != nil {
                http.Error(w, "Invalid token", http.StatusForbidden)
                return
            }

            user, _ := controller.IsUsernameExist(data.Username)

            ctx := context.WithValue(r.Context(), userCtxKey, user)

            r = r.WithContext(ctx)
            next.ServeHTTP(w, r)
        })
    }
}

func ForContext(ctx context.Context) (*controller.LoginData, bool) {
    data, ok := ctx.Value(userCtxKey).(*controller.LoginData)
    return data, ok
}
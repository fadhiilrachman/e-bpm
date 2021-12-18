package middleware

import (
    "context"
    "net/http"
    "strings"

    controller "github.com/fadhiilrachman/e-bpm/controllers"
    utils "github.com/fadhiilrachman/e-bpm/utils"
)

var userCtxKey = &contextKey{"admin", "username"}
type contextKey struct {
    Role string
    Username string
}

func Auth() func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(
            func(w http.ResponseWriter, r *http.Request) {
                reqToken := r.Header.Get("authorization")
                
                if reqToken == "" {
                    next.ServeHTTP(w, r)
                    return
                }

                splitToken := strings.Split(reqToken, "Bearer ")
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
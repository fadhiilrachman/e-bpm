package middleware

import (
    "context"
    "net/http"
    "strings"

    controller "github.com/fadhiilrachman/e-bpm/controllers"
)

func Auth() func(http.Handler) http.Handler {
    
}
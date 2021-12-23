package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fadhiilrachman/e-bpm/config"
	"github.com/fadhiilrachman/e-bpm/graph"
	"github.com/fadhiilrachman/e-bpm/graph/generated"
	middleware "github.com/fadhiilrachman/e-bpm/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
    config.Connection()
    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    router := chi.NewRouter()

    routerOptions := cors.Options {
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "Authorization"},
        AllowCredentials: true,
    }

    newRouter := cors.New(routerOptions).Handler

    router.Use(middleware.Auth())
    router.Use(newRouter)

    srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
    websocketTransport := &transport.Websocket{
        Upgrader: websocket.Upgrader{
            CheckOrigin: func (r *http.Request) bool {
                return r.Host == "localhost"
            },
            ReadBufferSize:  1024,
            WriteBufferSize: 1024,
        },
    }

    srv.AddTransport(websocketTransport)

    router.Handle("/", playground.Handler("GraphQL playground", "/query"))
    router.Handle("/query", srv)
    log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

    err := http.ListenAndServe(":"+port, router)
    if err != nil {
        panic(err)
    }
}

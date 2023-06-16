package runtime

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"go.uber.org/fx"

	"github.com/pulkitbhardwaj/matrix/internal"
)

func HTTPServer(lc fx.Lifecycle, gql *handler.Server, db *internal.Client) {

	// Create Router
	mux := chi.NewRouter()

	// Add Middlewares
	mux.Use(middleware.Logger)
	mux.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	}).Handler)

	gql.Use(entgql.Transactioner{TxOpener: db})

	// Add Routes
	mux.Method(http.MethodPost,
		"/", gql,
	)
	mux.Method(http.MethodGet,
		"/", playground.Handler("GraphQL Playground", "/"),
	)
	mux.Get()

	// Create server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}

package graphql

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type GQLParams struct {
	fx.In

	Router chi.Router
	// TxOpener internal.State
	Schema graphql.ExecutableSchema
}

func Handler(p GQLParams) {
	srv := handler.New(p.Schema)

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	// srv.Use(entgql.Transactioner{TxOpener: p.TxOpener})

	p.Router.Post("/", srv.ServeHTTP)
	p.Router.Get("/", playground.Handler("GraphQL Playground", "/"))
}

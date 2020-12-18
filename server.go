package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/ivansukach/octa-web-wallet-api/graph"
	"github.com/ivansukach/octa-web-wallet-api/graph/generated"
	"github.com/ivansukach/octa-web-wallet-api/repositories/validators"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"
)

const defaultPort = "4000"

func CustomCORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Headers", "fingerprint, authorization, development, content-type")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		//upgrader.CheckOrigin = func(r *http.Request) bool { return true}
		next.ServeHTTP(w, r)
	})
}

type Cache struct {
	client redis.UniversalClient
	ttl    time.Duration
}

const apqPrefix = "apq:"

func NewCache(redisAddress string, ttl time.Duration) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, fmt.Errorf("could not create cache: %w", err)
	}

	return &Cache{client: client, ttl: ttl}, nil
}

func (c *Cache) Add(ctx context.Context, key string, value interface{}) {
	c.client.Set(apqPrefix+key, value, c.ttl)
}

func (c *Cache) Get(ctx context.Context, key string) (interface{}, bool) {
	s, err := c.client.Get(apqPrefix + key).Result()
	if err != nil {
		return struct{}{}, false
	}
	return s, true
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("P: ", p)
		log.Println("MessageType: ", messageType)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

type wsHandler struct {
}

func (t wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	//fmt.Fprint(w, "Time: "+time.Now().UTC().String())
	reader(ws)
}

type timeHandler struct {
}

func (t timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Time: "+time.Now().UTC().String())
}

func NewServer(es graphql.ExecutableSchema) *handler.Server {
	srv := handler.New(es)

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
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
	return srv
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db, err := sqlx.Connect("postgres",
		"user=su password=su "+
			"host=localhost dbname=mintscan")
	if err != nil {
		log.Fatal(err)
	}
	//myWsHandler := wsHandler{}
	//myTimeHandler := timeHandler{}
	validatorRps := validators.New(db)
	srv := NewServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(validatorRps)}))
	//graphqlHandlerFunc := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(validatorRps)}), )
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(CustomCORSMiddleware)
	r.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	r.Handle("/graphql", srv)
	//r.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("welcome"))
	//})
	//r.Handle("/graphql", myWsHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

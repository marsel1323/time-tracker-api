package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/marsel1323/timetrackerapi/graph"
	"github.com/marsel1323/timetrackerapi/graph/generated"
	"github.com/marsel1323/timetrackerapi/repository"
	"github.com/marsel1323/timetrackerapi/service"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const defaultPort = 8080

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if authorization == "" {
			fmt.Println("Token is absent")
			//next.ServeHTTP(w, r)
			http.Error(w, "API key absent", http.StatusUnauthorized)
			//json.NewEncoder(w).Encode("error")
			return
		}

		fmt.Println(authorization)
		next.ServeHTTP(w, r)
	})
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", defaultPort, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://time_tracker:password@localhost/time_tracker?sslmode=disable", "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")
	flag.Parse()
	fmt.Println(cfg)
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	logger.Printf("Database connection pool established")

	taskService := service.NewTaskService(repository.NewTaskRepository(db))
	statisticService := service.NewStatisticService(repository.NewStatisticRepository(db))
	categoryService := service.NewCategoryService(repository.NewCategoryRepository(db))

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: graph.NewResolver(
					taskService,
					statisticService,
					categoryService,
				),
			},
		),
	)

	//var clientID = "55044703536-rh9ajuq0q14f5n12otadfaa3dobsvr9a.apps.googleusercontent.com"
	//var clientSecret = "LXAbveh77fTbZaCaNcvwGuZ_"
	//
	//goth.UseProviders(
	//	google.New(clientID, clientSecret, "http://localhost:8080/auth/google/callback"),
	//)

	//store := sessions.NewCookieStore([]byte(""))
	//store.MaxAge(86400 * 30)
	//store.Options.Path = "/"
	//store.Options.HttpOnly = true // HttpOnly should always be enabled
	//store.Options.Secure = false

	//gothic.Store = store

	//router := mux.NewRouter()

	//router.HandleFunc("/auth/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {
	//	user, err := gothic.CompleteUserAuth(w, r)
	//	if err != nil {
	//		fmt.Fprintln(w, err)
	//		return
	//	}
	//	fmt.Fprintln(w, user)
	//})
	//router.HandleFunc("/auth/{provider}", func(w http.ResponseWriter, r *http.Request) {
	//	gothUser, err := gothic.CompleteUserAuth(w, r)
	//	if err == nil {
	//		fmt.Fprintln(w, fmt.Sprintln("Zdarova: ", gothUser))
	//	} else {
	//		gothic.BeginAuthHandler(w, r)
	//	}
	//})

	//http.Handle("/", router)
	//http.ListenAndServe(":8080", nil)

	mux := http.NewServeMux()

	//mux.Handle("/auth/google/callback", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	user, err := gothic.CompleteUserAuth(w, r)
	//	if err != nil {
	//		fmt.Fprintln(w, err)
	//		return
	//	}
	//	fmt.Fprintln(w, user)
	//}))
	//mux.Handle("/auth/google", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	gothUser, err := gothic.CompleteUserAuth(w, r)
	//	if err == nil {
	//		fmt.Fprintln(w, fmt.Sprintln("Zdarova: ", gothUser))
	//	} else {
	//		gothic.BeginAuthHandler(w, r)
	//	}
	//
	//}))

	mux.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", middleware(srv))

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", cfg.port)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), handler)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5+time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

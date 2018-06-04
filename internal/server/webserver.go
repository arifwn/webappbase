package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/qor/auth"
	// "github.com/qor/auth/auth_identity"
	// "github.com/qor/auth/providers/password"
	"github.com/qor/session/manager"

	"github.com/qor/auth_themes/clean"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	// "github.com/arifwn/webappbase/pkg/auth"
	"github.com/arifwn/webappbase/pkg/conf"
)

// Auth middleware: redirect to login page if not logged in
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index\n"))

	// user := auth.UserFromContext(r.Context())
	// log.Println("User: ", user)
}

func RunServer() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()
	// // Add your routes as needed

	r.HandleFunc("/", IndexHandler)

	// authRouter := r.PathPrefix("/auth/").Subrouter()
	// auth.AttachHandlers(authRouter)

	config := conf.Get()

	gormDB, _ := gorm.Open(config.DBType, config.DBConf)

	// Initialize Auth with configuration
	// Auth := auth.New(&auth.Config{
	// 	DB: gormDB,
	// })
	Auth := clean.New(&auth.Config{
		DB: gormDB,
	})

	// gormDB.AutoMigrate(&auth_identity.AuthIdentity{})

	// Register Auth providers
	// Allow use username/password
	// Auth.RegisterProvider(password.New(&password.Config{}))

	// authRouter := r.PathPrefix("/auth/").Subrouter()
	r.PathPrefix("/auth/").Handler(Auth.NewServeMux())
	r.Use(manager.SessionManager.Middleware)

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%s", config.ServerAddress, config.ServerPort),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

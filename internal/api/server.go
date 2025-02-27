package api

import (
	"net/http"
	"time"
	"tube-nutrition/internal/api/handler"
	"tube-nutrition/internal/config"
	"tube-nutrition/internal/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"tube-nutrition/internal/api/middleware"
	"tube-nutrition/internal/utils"
)

const (
	maxHeaderBytes = 1 << 20
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
)

type Server struct {
	httpServer *http.Server
	router     *mux.Router
}

func NewServer(ctx utils.MyContext, config config.Config) *Server {
	router := mux.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	corsRouter := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)

	recoveredRouter := middleware.RecoveryMiddleware(ctx, corsRouter)

	authenticatedRouter := middleware.UserIdentity(ctx, recoveredRouter)

	return &Server{
		httpServer: &http.Server{
			Addr:           config.ServerPort,
			MaxHeaderBytes: maxHeaderBytes,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
			Handler:        authenticatedRouter,
		},
		router: router,
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx utils.MyContext) error {
	return s.httpServer.Shutdown(ctx.Ctx)
}

func (s *Server) HandleAuth(ctx utils.MyContext, service service.NutritionService) {
	s.router.HandleFunc("/api/nutrition/", handler.NutritionPlan(ctx, service)).Methods(http.MethodPost)
	s.router.HandleFunc("/api/nutrition/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // можно заменить на конкретный домен
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodOptions)
}

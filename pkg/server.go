package pkg

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewMux(lc fx.Lifecycle, logger *zap.Logger) *mux.Router {
	logger.Info("Executing NewMux.")

	r := mux.NewRouter()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*", "http://localhost:5001", "localhost:5001", "localhost"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "HEAD", "OPTIONS"},
		Debug:            true,
	}).Handler)

	// adding middleware
	// router.Use(middleware.RequestID)
	// // router.Use(middleware.Logger)
	// router.Use(customMiddleware.AuthMiddleware(userRepo))
	r.Use(AuthMiddleware())
	// r.Use(middleware.AuthMiddleware())

	//

	server := &http.Server{
		// Addr:    fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")),
		Addr:    "localhost:5001",
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting HTTP server.")
			go func() {
				if err := server.ListenAndServe(); err != nil {
					logger.Sugar().Error(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return r
}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gadavy/lhw/zap"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"

	"github.com/lissteron/cloudsuny/config"
	"github.com/lissteron/cloudsuny/internal/app/http/handlers"
	"github.com/lissteron/cloudsuny/internal/app/repositories"
	"github.com/lissteron/cloudsuny/internal/app/usecases"
	"github.com/lissteron/cloudsuny/pkg/server"
	"github.com/lissteron/cloudsuny/pkg/sqlite3"
)

// nolint: funlen,gocritic
func main() {
	// Init config, logger, exit chan
	config.Init()

	logger, err := zap.NewLogger(config.LoggerConfig())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "init logger failed: %v", err)
		os.Exit(1)
	}

	defer logger.Close()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	// Init db
	sqliteDB, err := sqlite3.NewClient(viper.GetString("db.path"))
	if err != nil {
		logger.Fatalf("init db failed: %v", err)
	}

	// Init repository
	repository := repositories.NewRepository(sqliteDB.DB(), logger)

	// Init usecases
	var (
		userCreate = usecases.NewCreateUser(repository, logger)
		userList   = usecases.NewViewUser(repository, logger)

		badgeCreate = usecases.NewCreateBadge(repository, logger)
		badgeUpdate = usecases.NewUpdateBadge(repository, logger)
	)

	// Init handlers
	var (
		userHandlers = handlers.NewUserHandlers(userCreate, userList, logger)
		badgeHandler = handlers.NewBadgeHandlers(badgeCreate, badgeUpdate, logger)
	)

	srv := server.NewHTTP(config.ServerConfig())

	r := srv.Router()
	r1 := r.PathPrefix("/api/v1/").Subrouter()
	r1.HandleFunc("/user/create", userHandlers.CreateHandler)
	r1.HandleFunc("/view", userHandlers.ViewHandler)
	r1.HandleFunc("/badge/create", badgeHandler.CreateHandler)
	r1.HandleFunc("/badge/update", badgeHandler.UpdateHandler)

	errGroup, ctx := errgroup.WithContext(context.Background())

	errGroup.Go(func() error {
		logger.Infof("start http server on: %s", srv.Addr())
		return srv.ListenAndServe()
	})

	select {
	case <-exit:
		logger.Info("stopping application")
	case <-ctx.Done():
		logger.Error("stopping application with error")
	}

	if err = srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error while stopping web server: %v", err)
	}

	if err = errGroup.Wait(); err != nil {
		logger.Errorf("error while waiting for goroutines: %v", err)
	}

	logger.Info("application stopped")
}

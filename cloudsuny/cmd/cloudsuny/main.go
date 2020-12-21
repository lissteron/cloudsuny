package main

import (
	"log"
	"net/http"

	"github.com/loghole/database"
	"github.com/loghole/tron"

	"github.com/lissteron/cloudsuny/config"
	authV1 "github.com/lissteron/cloudsuny/internal/app/controllers/auth/v1"
	badgesV1 "github.com/lissteron/cloudsuny/internal/app/controllers/badges/v1"
	httpV1 "github.com/lissteron/cloudsuny/internal/app/controllers/http"
	usersV1 "github.com/lissteron/cloudsuny/internal/app/controllers/users/v1"
	"github.com/lissteron/cloudsuny/internal/app/repositories"
	"github.com/lissteron/cloudsuny/internal/app/services/auth"
	"github.com/lissteron/cloudsuny/internal/app/services/badges"
	"github.com/lissteron/cloudsuny/internal/app/services/images"
	"github.com/lissteron/cloudsuny/internal/app/services/users"
)

func main() {
	app, err := tron.New(tron.AddLogCaller(), tron.WithRealtimeConfig())
	if err != nil {
		log.Fatalf("can't create app: %s", err)
	}

	defer app.Close()

	// Init db.
	sqliteDB, err := database.New(config.SQLiteConfig(), database.WithDefaultOptions(app.Tracer()))
	if err != nil {
		app.Logger().Fatalf("init db failed: %v", err)
	}

	if err := repositories.Migrate(sqliteDB); err != nil {
		app.Logger().Fatalf("migrate failed: %v", err)
	}

	// Init repository.
	repository := repositories.NewRepository(sqliteDB, app.TraceLogger())

	// Init services.
	var (
		imageService = images.NewService(config.ImgPath(), app.TraceLogger())
		badgeService = badges.NewService(repository, app.TraceLogger())
		userService  = users.NewService(repository, app.TraceLogger())
		authService  = auth.NewService(repository, app.TraceLogger())
	)

	// Init controllers.
	var (
		badgesV1Handler = badgesV1.NewImplementation(badgeService, app.TraceLogger())
		usersV1Handler  = usersV1.NewImplementation(userService, app.TraceLogger())
		authV1Handlers  = authV1.NewImplementation(authService, app.TraceLogger())
		imageHandlers   = httpV1.NewImagesHandlers(imageService, app.TraceLogger())
	)

	// Init routes.
	app.Router().Use(authV1Handlers.Middleware)

	app.Router().HandleFunc("/api/v1/image/upload", imageHandlers.UploadPhotoHandler)
	app.Router().Mount("/images/", http.StripPrefix("/images/", imageHandlers.FileServer()))

	if err := app.WithRunOptions().Run(badgesV1Handler, usersV1Handler, authV1Handlers); err != nil {
		app.Logger().Fatalf("can't run app: %v", err)
	}

	if err := sqliteDB.Close(); err != nil {
		app.Logger().Errorf("failed to close database: %v", err)
	}
}

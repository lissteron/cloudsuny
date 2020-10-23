package main

import (
	"log"
	"net/http"

	"github.com/lissteron/cloudsuny/config"
	badgesV1 "github.com/lissteron/cloudsuny/internal/app/controllers/badges/v1"
	httpV1 "github.com/lissteron/cloudsuny/internal/app/controllers/http"
	usersV1 "github.com/lissteron/cloudsuny/internal/app/controllers/users/v1"
	"github.com/lissteron/cloudsuny/internal/app/repositories"
	"github.com/lissteron/cloudsuny/internal/app/services/images"
	"github.com/lissteron/cloudsuny/internal/app/usecases"
	"github.com/lissteron/cloudsuny/internal/pkg/sqlite3"

	"github.com/loghole/tron"
)

func main() {
	app, err := tron.New(tron.AddLogCaller())
	if err != nil {
		log.Fatalf("can't create app: %s", err)
	}

	defer app.Close()

	// Init db.
	sqliteDB, err := sqlite3.NewClient(config.DBPath())
	if err != nil {
		app.Logger().Fatalf("init db failed: %v", err)
	}

	// Init repository.
	repository := repositories.NewRepository(sqliteDB.DB(), app.Logger())

	// Init services.
	imageService := images.NewService(config.ImgPath(), app.Logger())

	// Init use cases.
	var (
		userCreate = usecases.NewCreateUser(repository, app.Logger())
		userList   = usecases.NewViewUser(repository, app.Logger())

		badgeCreate = usecases.NewCreateBadge(repository, app.Logger())
		badgeUpdate = usecases.NewUpdateBadge(repository, app.Logger())
	)

	// Init handlers.
	var (
		badgesV1Handler = badgesV1.NewBadges(badgeCreate, badgeUpdate, app.TraceLogger())
		usersV1Handler  = usersV1.NewUsers(userCreate, userList, app.TraceLogger())
		imageHandlers   = httpV1.NewImagesHandlers(imageService, app.TraceLogger())
	)

	// Init routes.
	app.Router().HandleFunc("/api/v1/image/upload", imageHandlers.UploadPhotoHandler)
	app.Router().Mount("/images/", http.StripPrefix("/images/", imageHandlers.FileServer()))

	if err := app.WithRunOptions().Run(badgesV1Handler, usersV1Handler); err != nil {
		app.Logger().Fatalf("can't run app: %v", err)
	}

	if err := sqliteDB.Close(); err != nil {
		app.Logger().Errorf("failed to close database: %v", err)
	}
}

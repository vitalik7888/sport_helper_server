package main

import (
	"fmt"
	"net"
	"net/http"
	"sport_helper/internal/config"
	"sport_helper/internal/handlers"
	"sport_helper/internal/service"
	"sport_helper/pkg/logger"
	db "sport_helper/pkg/sqldb"

	"github.com/julienschmidt/httprouter"
)

type appHandlers struct {
	router           *httprouter.Router
	person           handlers.Handler
	exercise         handlers.Handler
	trainingSession  handlers.Handler
	trainingExercise handlers.Handler
}

func newHandlers(repositories *db.DbClient) *appHandlers {
	return &appHandlers{
		router: httprouter.New(),
		person: handlers.NewPersonHandler(
			service.NewPersonService(repositories.Persons()),
		),
		exercise: handlers.NewExerciseHandler(
			service.NewExerciseService(repositories.Exercise()),
		),
		trainingSession: handlers.NewTrainingSessionHandler(
			service.NewTrainingSessionService(
				repositories.TrainingSession(),
			),
		),
		trainingExercise: handlers.NewTrainingExerciseHandler(
			service.NewTrainingExerciseService(
				repositories.TrainingExercise(),
			),
		),
	}
}

func (h *appHandlers) register() {
	h.person.Register(h.router)
	h.exercise.Register(h.router)
	h.trainingSession.Register(h.router)
	h.trainingExercise.Register(h.router)
}

type App struct {
	dbClient *db.DbClient
	handlers *appHandlers
}

func NewApp() *App {
	logger := logger.GetLogger()
	config := config.GetConfig()

	repositories, err := db.NewSqliteClient(config.Storage.DatabaseName)
	if err != nil {
		logger.Fatalf("App error: %v", err)
	}
	h := newHandlers(repositories)
	h.register()

	return &App{
		dbClient: repositories,
		handlers: h,
	}
}

func (a *App) Start() {
	logger := logger.GetLogger()
	config := config.GetConfig()

	logger.Info("Starting server")
	address := fmt.Sprintf("%s:%s", config.Server.BindAddress, config.Server.BindPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatalf("Server listening error: ", err)
	}
	logger.Debugf("Server is listening on %s", address)
	logger.Fatalln(http.Serve(listener, a.handlers.router))
}

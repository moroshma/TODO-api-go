package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	toDoRepo "github.com/moroshma/TODO-api-go/internal/todo/repository"
	toDoService "github.com/moroshma/TODO-api-go/internal/todo/service"
	"github.com/sirupsen/logrus"
)

type service struct {
	logger      *logrus.Logger
	router      *mux.Router
	toDoService toDoService.Service
}

func newHandler(lg *logrus.Logger, db *sqlx.DB) service {
	return service{
		logger:      lg,
		toDoService: toDoService.NewService(toDoRepo.NewRepository(db)),
	}
}

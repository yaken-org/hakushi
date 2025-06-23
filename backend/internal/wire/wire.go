//go:build wireinject

package wire

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/yaken-org/hakushi/internal/adapter/handler"
	"github.com/yaken-org/hakushi/internal/adapter/repository"
	"github.com/yaken-org/hakushi/internal/config"
	"github.com/yaken-org/hakushi/internal/infrastructure/database"
	"github.com/yaken-org/hakushi/internal/infrastructure/server"
	"github.com/yaken-org/hakushi/internal/usecase"
)

var repositorySet = wire.NewSet(
	repository.NewPostRepository,
	repository.NewUserAccountRepository,
	repository.NewAnnotationRepository,
	repository.NewTagRepository,
	repository.NewPostTagRepository,
)

var usecaseSet = wire.NewSet(
	usecase.NewPostUsecase,
	usecase.NewUserAccountUsecase,
	database.NewTransactionManager,
)

var handlerSet = wire.NewSet(
	handler.NewPostHandler,
	handler.NewUserAccountHandler,
)

func InitializeServer(config *config.Config, db *sql.DB) (*server.Server, error) {
	wire.Build(
		repositorySet,
		usecaseSet,
		handlerSet,
		server.New,
	)
	return nil, nil
}
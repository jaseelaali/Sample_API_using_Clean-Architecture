package dep

import (
	"net/http"

	"github.com/google/wire"
	db "github.com/jaseelaali/Sample_API_using_Clean-Architecture/database"
	config "github.com/jaseelaali/Sample_API_using_Clean-Architecture/pkg/config"
)	

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}

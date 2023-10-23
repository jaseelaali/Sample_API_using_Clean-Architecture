package dep

import (
	"net/http"

	"github.com/google/wire"
	"github.com/jaseelaali/Sample_API_using_Clean-Architecture/database"
	c "github.com/jaseelaali/Sample_API_using_Clean-Architecture/pkg/config"
)

func InitializeApi(config c.Config) (*http.ServerHTTP, error) {
	wire.Build(
		database.ConnectDatabase,
		repository.NewRepository,
		usecase.NewUsecase,
		handler.NewHandler,
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil

}

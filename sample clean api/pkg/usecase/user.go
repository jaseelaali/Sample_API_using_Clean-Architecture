package usecase

import interfaces "github.com/jaseelaali/Sample_API_using_Clean-Architecture/pkg/repository/interface"

type UserUsecase struct {
	repo interfaces.UserRepository
}

func NewUsecase(repos interfaces.UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: repos,
	}
}

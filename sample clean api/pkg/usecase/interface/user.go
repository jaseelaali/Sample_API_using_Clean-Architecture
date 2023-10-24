package interfaces

import "context"

type UserUsecase interface {
	Wish(ctx context.Context)(string,error)
}

package interfaces

import "context"

type UserRepository interface {
	Wish(ctx context.Context)(string,error)
}

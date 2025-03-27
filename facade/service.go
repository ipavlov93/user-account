package facade

import (
	"context"
	"event-calendar/smodel"
)

//type AuthService interface {
//	VerifyToken(token string) (claims auth.Claims, error)
//}

type UserService interface {
	GetUserByID(ctx context.Context, id int64) (userDto smodel.User, err error)
	GetUserByUUID(ctx context.Context, uuid string) (smodel.User, error)
	CreateUser(ctx context.Context, user smodel.User) (int64, error)
}

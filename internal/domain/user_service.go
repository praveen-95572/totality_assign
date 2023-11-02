package domain

import "context"

type Service interface {
	GetUserByID(ctx context.Context, userID int64) (*User, error)
}

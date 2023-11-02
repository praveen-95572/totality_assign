package user

import (
	"context"
	"fmt"
	"userservice/internal/domain"
)

func (s *userservice) GetUserByID(ctx context.Context, userID int64) (*domain.User, error) {
	for _, user := range mockUsers {
		if user.Id == userID {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

package user

import "userservice/internal/domain"

var mockUsers = []*domain.User{
	{
		Id:      1,
		Fname:   "Praveen Bisht",
		City:    "Noida",
		Phone:   1234567890,
		Height:  5.9,
		Married: false,
	},
	{
		Id:      1,
		Fname:   "Rahul Thapliyal",
		City:    "delhi",
		Phone:   1234567899,
		Height:  5.6,
		Married: false,
	},
}

type userservice struct {
}

func NewService() *userservice {
	return &userservice{}
}

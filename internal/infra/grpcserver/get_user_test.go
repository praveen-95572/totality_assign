package grpcserver

import (
	"context"
	"errors"
	"testing"

	"userservice/internal/domain"
	pb "userservice/internal/infra/grpcserver/.gen"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Create a mock for the UserService interface
type UserServiceMock struct {
	mock.Mock
}

func (srv *UserServiceMock) GetUserByID(ctx context.Context, userID int64) (*domain.User, error) {
	if userID == 1 {
		return &domain.User{
			Id:      1,
			Fname:   "Praveen Bisht",
			City:    "Noida",
			Phone:   1234567890,
			Height:  5.9,
			Married: false,
		}, nil
	} else {
		return nil, errors.New("user not found")
	}
}

func TestGetUserById(t *testing.T) {
	// Create a mock for the UserService interface
	userServiceMock := new(UserServiceMock)

	// Create a Server instance with the mock UserService
	server := &Server{
		userservice: userServiceMock,
	}

	tests := []struct {
		testName      string
		userID        int64
		expectedFname string
		expectedErr   error
	}{
		{
			testName:      "User Found",
			userID:        1,
			expectedFname: "Praveen Bisht",
		},
		{
			testName:    "User not Found",
			userID:      10,
			expectedErr: errors.New("user not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			// Create a UserRequest with the desired user ID
			userRequest := &pb.UserRequest{
				Id: tt.userID,
			}

			// Call the GetUserById method
			ctx := context.Background()
			userResponse, err := server.GetUserById(ctx, userRequest)

			if err != nil {
				assert.ErrorContains(t, err, tt.expectedErr.Error(), "Error should match")
			} else {
				assert.Equal(t, tt.expectedFname, userResponse.Fname, "Should match the Full Name")
			}
		})
	}

}

// Create a mock for the gRPC server stream
type MockUserStream struct {
	mock.Mock
}

// Implement the Send method of the gRPC server stream
func (m *MockUserStream) Send(user *pb.User) error {
	args := m.Called(user)
	return args.Error(0)
}

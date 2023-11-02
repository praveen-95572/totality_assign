package user

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByID_UserFound(t *testing.T) {
	// Create a test userservice with a known list of users
	userService := &userservice{}

	// Define the test case inputs and expected output
	ctx := context.TODO()

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
			// Call the GetUserByID method
			user, err := userService.GetUserByID(ctx, tt.userID)
			// Check if the user is found and the error is nil
			if err != nil {
				assert.ErrorContains(t, err, tt.expectedErr.Error(), "Error should match")
			} else {
				assert.Equal(t, tt.expectedFname, user.Fname, "Should match the Full Name")
			}
		})
	}

}

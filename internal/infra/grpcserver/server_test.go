package grpcserver

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	// Create a mock implementation of the UserService interface
	userServiceMock := &UserServiceMock{} // You'll need to define your own UserServiceMock

	// Call the NewServer function to create a new Server instance
	server := NewServer(userServiceMock)

	if server == nil {
		t.Error("Expected a non-nil Server, but got nil")
	}
}

package user

import "testing"

func TestNewService(t *testing.T) {
	service := NewService()
	if service == nil {
		t.Error("Expected a non-nil service, but got nil")
	}
}

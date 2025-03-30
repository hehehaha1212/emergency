package validation

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	"emergency-app/backend/internal/models"
)

func TestValidateRequest(t *testing.T) {
	tests := []struct {
		name    string
		request models.Request
		wantErr bool
		errMsg  string
	}{
		{
			name: "Valid blood request",
			request: models.Request{
				Type:      "blood",
				BloodType: "O+",
			},
			wantErr: false,
		},
		{
			name: "Missing blood type",
			request: models.Request{
				Type: "blood",
			},
			wantErr: true,
			errMsg:  "blood type is required for blood requests",
		},
		{
			name: "Invalid blood type",
			request: models.Request{
				Type:      "blood",
				BloodType: "XYZ",
			},
			wantErr: true,
			errMsg:  "invalid blood type",
		},
		{
			name: "Valid oxygen request",
			request: models.Request{
				Type:        "oxygen",
				OxygenUnits: 5,
			},
			wantErr: false,
		},
		{
			name: "Invalid oxygen units",
			request: models.Request{
				Type:        "oxygen",
				OxygenUnits: 0,
			},
			wantErr: true,
			errMsg:  "oxygen units must be greater than 0",
		},
		{
			name: "Valid medicine request",
			request: models.Request{
				Type:     "medicine",
				Medicine: "Paracetamol",
			},
			wantErr: false,
		},
		{
			name: "Missing medicine name",
			request: models.Request{
				Type: "medicine",
			},
			wantErr: true,
			errMsg:  "medicine name is required",
		},
		{
			name: "Invalid request type",
			request: models.Request{
				Type: "invalid",
			},
			wantErr: true,
			errMsg:  "invalid request type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRequest(&tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateUser(t *testing.T) {
	tests := []struct {
		name    string
		user    models.User
		wantErr bool
		errMsg  string
	}{
		{
			name: "Valid user",
			user: models.User{
				Name:     "Test User",
				Email:    "test@example.com",
				Password: "password123",
			},
			wantErr: false,
		},
		{
			name: "Missing name",
			user: models.User{
				Email:    "test@example.com",
				Password: "password123",
			},
			wantErr: true,
			errMsg:  "name is required",
		},
		{
			name: "Missing email",
			user: models.User{
				Name:     "Test User",
				Password: "password123",
			},
			wantErr: true,
			errMsg:  "email is required",
		},
		{
			name: "Invalid email format",
			user: models.User{
				Name:     "Test User",
				Email:    "invalid-email",
				Password: "password123",
			},
			wantErr: true,
			errMsg:  "invalid email format",
		},
		{
			name: "Password too short",
			user: models.User{
				Name:     "Test User",
				Email:    "test@example.com",
				Password: "pass",
			},
			wantErr: true,
			errMsg:  "password must be at least 8 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUser(&tt.user)
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
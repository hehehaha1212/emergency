package validation

import (
	"errors"
	"strings"
	
	"emergency-app/backend/internal/models"
)


func ValidateRequest(request *models.Request) error {

	if request.Type == "" {
		return errors.New("request type is required")
	}

	switch strings.ToLower(request.Type) {
	case "blood":
		if request.BloodType == "" {
			return errors.New("blood type is required for blood requests")
		}
		validBloodTypes := map[string]bool{
			"a+": true, "a-": true, "b+": true, "b-": true,
			"ab+": true, "ab-": true, "o+": true, "o-": true,
		}
		if !validBloodTypes[strings.ToLower(request.BloodType)] {
			return errors.New("invalid blood type, must be one of: A+, A-, B+, B-, AB+, AB-, O+, O-")
		}
	case "oxygen":
		if request.OxygenUnits <= 0 {
			return errors.New("oxygen units must be greater than 0")
		}
	case "medicine":
		if request.Medicine == "" {
			return errors.New("medicine name is required for medicine requests")
		}
	default:
		return errors.New("invalid request type: must be 'blood', 'oxygen', or 'medicine'")
	}
	
	return nil
}

func ValidateUser(user *models.User) error {

	if user.Name == "" {
		return errors.New("name is required")
	}
	
	if user.Email == "" {
		return errors.New("email is required")
	}

	if !strings.Contains(user.Email, "@") || !strings.Contains(user.Email, ".") {
		return errors.New("invalid email format")
	}
	
	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	
	return nil
}
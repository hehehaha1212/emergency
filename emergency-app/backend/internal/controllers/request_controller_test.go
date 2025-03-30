package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"emergency-app/backend/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.User{}, &models.Hospital{}, &models.Request{})
	models.SetDB(db)
	return db
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func createTestRequest(t *testing.T, router *gin.Engine, userID uint) models.Request {
	// Create test request
	testRequest := models.Request{
		UserID:      userID,
		Type:        "blood",
		BloodType:   "O+",
		Status:      "Pending",
		HospitalID:  1,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("userID", userID)
	c.Request = httptest.NewRequest("POST", "/api/request", nil)
	
	reqBody, _ := json.Marshal(testRequest)
	c.Request.Body = http.NoBody
	c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	c.Request.Header.Set("Content-Type", "application/json")
	
	CreateRequest(c)
	

	var response struct {
		Data models.Request `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &response)
	
	return response.Data
}

func TestCreateRequest(t *testing.T) {
	db := setupTestDB()
	router := setupTestRouter()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpassword"), bcrypt.DefaultCost)
	testUser := models.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: string(hashedPassword),
	}
	db.Create(&testUser)
	

	reqBody := models.Request{
		Type:      "blood",
		BloodType: "A+",
	}

	w := httptest.NewRecorder()
	reqJSON, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/request", bytes.NewBuffer(reqJSON))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("userID", testUser.ID)
	
	
	CreateRequest(c)
	

	assert.Equal(t, http.StatusCreated, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Request created successfully", response["message"])
}

func TestGetRequest(t *testing.T) {
	db := setupTestDB()
	router := setupTestRouter()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpassword"), bcrypt.DefaultCost)
	testUser := models.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: string(hashedPassword),
	}
	db.Create(&testUser)

	testRequest := createTestRequest(t, router, testUser.ID)
	
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/request/"+strconv.Itoa(int(testRequest.ID)), nil)
	
	// Set up context with user ID
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("userID", testUser.ID)
	c.Params = []gin.Param{{Key: "id", Value: strconv.Itoa(int(testRequest.ID))}}
	
	// Call handler
	GetRequest(c)
	
	// Assert response
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	
	// Check returned data
	data, ok := response["data"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, float64(testRequest.ID), data["id"])
	assert.Equal(t, "blood", data["type"])
	assert.Equal(t, "O+", data["blood_type"])
}
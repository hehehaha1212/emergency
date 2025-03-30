package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"emergency-app/internal/models"
	"emergency-app/internal/services"
	"emergency-app/pkg/validation"
)

func CreateRequest(c *gin.Context) {
	var request models.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}

	if err := validation.ValidateRequest(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}
	
	request.UserID = userID.(uint)
	request.Status = "Pending" 

	if err := services.CreateRequest(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Request created successfully",
		"data": request,
	})
}

func GetRequest(c *gin.Context) {
	id := c.Param("id")
	var request models.Request
	
	if err := models.DB.First(&request, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}
	
	userID, _ := c.Get("userID")
	if request.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view this request"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": request})
}

func UpdateRequest(c *gin.Context) {
	id := c.Param("id")
	var request models.Request
	
	if err := models.DB.First(&request, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}
	
	
	userID, _ := c.Get("userID")
	if request.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this request"})
		return
	}
	
	// Only allow updating certain fields
	var input models.Request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}
	
	// Validate the updated request
	input.ID = request.ID
	input.UserID = request.UserID
	if err := validation.ValidateRequest(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	models.DB.Model(&request).Updates(input)
	c.JSON(http.StatusOK, gin.H{
		"message": "Request updated successfully",
		"data": request,
	})
}

func DeleteRequest(c *gin.Context) {
	id := c.Param("id")
	var request models.Request
	
	if err := models.DB.First(&request, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}
	
	// Check if user has access to this request
	userID, _ := c.Get("userID")
	if request.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this request"})
		return
	}
	
	if err := models.DB.Delete(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete request: " + err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Request deleted successfully"})
}
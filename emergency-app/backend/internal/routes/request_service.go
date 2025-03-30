package services

import(
	"github.com/gin-gonic/gin"
    "emergency-app/internal/models"
	"gorm.io/gorm"
)

func CreateRequest(request *model.request) error{

	return models.DB.create(request).error
}

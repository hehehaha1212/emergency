package routes

import(
	"github.com/gin-gonic/gin"
    "emergency-app/backend/internal/models"
	"gorm.io/gorm"
)

func CreateRequest(request *model.Request) error{

	return models.DB.Create(request).error
}

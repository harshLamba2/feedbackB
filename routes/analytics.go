package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harshLamba2/feedbackF/controllers/analytics"
	"github.com/harshLamba2/feedbackF/middleware"
	"github.com/harshLamba2/feedbackF/utils/constants"
	"gorm.io/gorm"
)

func AnalyticsRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.POST(constants.ExecuteSelectQuery, middleware.IsAdmin() ,analytics.ExecuteSelectQuery(db))
}

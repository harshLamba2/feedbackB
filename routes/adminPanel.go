package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harshLamba2/feedbackF/controllers/adminPanel"
	"github.com/harshLamba2/feedbackF/middleware"
	"github.com/harshLamba2/feedbackF/utils/constants"
	"gorm.io/gorm"
)

func AdminPanelRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.POST(constants.GetAllUsers, middleware.IsAdmin(), adminPanel.GetAllUsers(db))
	router.PUT(constants.PermissionAccess, middleware.IsAdmin(), adminPanel.PermissionAccess(db))
}

package routes

import (
	"github.com/harshLamba2/feedbackF/controllers/ticketing"
	"github.com/harshLamba2/feedbackF/utils/constants"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TicketRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.POST(constants.AddTicket, ticketing.AddTicket(db))
	router.POST(constants.SaveBusColors, ticketing.SaveBusColors(db))
	router.POST(constants.SaveBusInitials, ticketing.SaveBusInitials(db))
	router.POST(constants.SaveBusStops, ticketing.SaveBusStops(db))
	router.POST(constants.SaveBusRoutes, ticketing.SaveBusRoutes(db))
}

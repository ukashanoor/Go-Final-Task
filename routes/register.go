package routes

import (
	"net/http"
	"strconv"
	"ukashanoor/event-booking/models"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userid := context.GetInt64("userid")
	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}

	event, err := models.GetEventByID(eventid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event ID"})
		return
	}

	err = event.Register(userid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Register"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Register for event sucessfully"})
}

func cancelRegistration(context *gin.Context) {
	userid := context.GetInt64("userid")
	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}
	var event models.Event
	event.ID = eventid
	err = event.CancelRegistration(userid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel Register"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled for event sucessfully"})

}

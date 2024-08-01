package handler

import (
	"api_gate_way/internal/entity"
	alarm_usecase "api_gate_way/internal/usecase/alarm"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Alarm struct {
	alarm alarm_usecase.Alarm
}

func NewAlarm(alarm alarm_usecase.Alarm) *Alarm {
	return &Alarm{alarm: alarm}
}

// AddAlarm godoc
// @Summary AddAlarm
// @Description Add a new smart alarm
// @Tags alarm
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body entity.AddSmartAlarmReq true "Add Smart Alarm Request"
// @Success 200 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /alarm/register [post]
func (a *Alarm) AddAlarm(c *gin.Context) {
	var req entity.AddSmartAlarmReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user_id not found"})
		return
	}
	req.UserID = id.(string)
	message, err := a.alarm.AddSmartAlarm(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// OpenCurtain godoc
// @Summary Open Curtain
// @Description Open the curtain
// @Tags alarm
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param device_name query string true "Name of the device"
// @Param open query string true "Open command"
// @Success 200 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /alarm/open/curtain [put]
func (a *Alarm) OpenCurtain(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user_id not found"})
		return
	}
	deviceName := c.Query("device_name")
	open := c.Query("open")
	if deviceName == "" || open == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "device_name or open is empty"})
		return
	}
	message, err := a.alarm.OpenCurtain(c.Request.Context(), id.(string), deviceName, open)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}

// OpenDoor godoc
// @Summary Open OpenDoor
// @Description Open the alarm
// @Tags alarm
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param device_name query string true "Name of the device"
// @Param open query string true "Open command"
// @Success 200 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /alarm/open/door [put]
func (a *Alarm) OpenDoor(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user_id not found"})
		return
	}
	deviceName := c.Query("device_name")
	open := c.Query("open")
	if deviceName == "" || open == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "device_name or open is empty"})
		return
	}
	message, err := a.alarm.OpenDoor(c.Request.Context(), id.(string), deviceName, open)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// CreateAlarmClock godoc
// @Summary Create Alarm Clock
// @Description Create a new alarm clock
// @Tags alarm
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body entity.CreateAlarmClockReq true "Create Alarm Clock Request"
// @Success 200 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /alarm/clock [post]
func (a *Alarm) CreateAlarmClock(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user_id not found"})
		return
	}
	var createAlarmClockReq entity.CreateAlarmClockReq
	if err := c.ShouldBindJSON(&createAlarmClockReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	createAlarmClockReq.UserID = id.(string)
	message, err := a.alarm.CreateAlarmClock(c.Request.Context(), &createAlarmClockReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}

// GetRemainingTime godoc
// @Summary Get Remaining Time
// @Description Get the remaining time for the alarm
// @Tags alarm
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param device_name query string true "Name of the device"
// @Success 200 {object} entity.RemainingTimRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /alarm/clock [get]
func (a *Alarm) GetRemainingTime(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user_id not found"})
		return
	}
	deviceName := c.Query("device_name")
	if deviceName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "device_name is empty"})
		return
	}
	alarms, err := a.alarm.GetRemainingTime(c.Request.Context(), id.(string), deviceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"alarms": alarms})
}

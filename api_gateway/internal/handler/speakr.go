package handler

import (
	"api_gate_way/internal/entity"
	speaker_use_case "api_gate_way/internal/usecase/speaker"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Speaker struct {
	speaker speaker_use_case.SpeakerUseCase
}

func NewSpeaker(tv speaker_use_case.SpeakerUseCase) *Speaker {
	return &Speaker{
		speaker: tv,
	}
}

// AddSpeaker godoc
// @Summary AddSpeaker
// @Description AddSpeaker add speaker
// @Tags speaker
// @Accept json
// @Produce json
// @Param body body entity.AddSpeakerReq true "tv create information"
// @Security ApiKeyAuth
// @Success 200 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /speaker/register [post]
// @Security BearerAuth
func (t *Speaker) AddSpeaker(c *gin.Context) {
	var addTvReq entity.AddSpeakerReq
	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}
	if err := c.ShouldBindJSON(&addTvReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message, err := t.speaker.AddSpeaker(c.Request.Context(), &entity.AddSpeakerReq{
		ModelName: addTvReq.ModelName,
		UserID:    id.(string),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// OpenSpeaker godoc
// @Summary OpenSpeaker
// @Description Add tv to home
// @Tags television
// @Accept json
// @Produce json
// @Param open query string false "true or false"
// @Security ApiKeyAuth
// @Success 200 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /speaker/open [put]
// @Security BearerAuth
func (t *Speaker) OpenSpeaker(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}
	open := c.Query("open")
	message, err := t.speaker.OpenSpeaker(c.Request.Context(), open, id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// GetChannels godoc
// @Summary GetChannels
// @Description Add tv to home
// @Tags television
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} entity.GetAllUserRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /speaker/user/channels [get]
// @Security BearerAuth
func (t *Speaker) GetChannels(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}
	channels, err := t.speaker.GetChannels(c.Request.Context(), id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"channels": channels})
}

// PreviousAndNext godoc
// @Summary PreviousAndNext
// @Description Add tv to home
// @Tags television
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param next query string false "true or false"
// @Success 200 {object} entity.PreviousAndNextRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /speaker/channel/cursor [post]
// @Security BearerAuth
func (u *Speaker) PreviousAndNext(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}
	next := c.Query("next")
	message, err := u.speaker.NextChannel(c.Request.Context(), id.(string), next)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// AddChannel godoc
// @Summary AddChannel
// @Description Add tv to home
// @Tags television
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body entity.AddChannelReqChannel true "channel create information"
// @Success 200 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /speaker/channel [post]
// @Security BearerAuth
func (t *Speaker) AddChannel(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}
	var addChannelReq entity.AddChannelReqChannel
	if err := c.ShouldBindJSON(&addChannelReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message, err := t.speaker.AddChannel(c.Request.Context(), &entity.AddChannelReqChannel{
		UserID:      id.(string),
		ChannelName: addChannelReq.ChannelName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// DeleteChannel godoc
// @Summary DeleteChannel
// @Description Add tv to home
// @Tags television
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body entity.DeleteChannelReq true "channel create information"
// @Success 200 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /speaker/channel [delete]
// @Security BearerAuth
func (t *Speaker) DeleteChannel(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}
	var deleteChannelReq entity.DeleteChannelReq
	if err := c.ShouldBindJSON(&deleteChannelReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	message, err := t.speaker.DeleteChannel(c.Request.Context(), &entity.DeleteChannelReqChannel{
		UserID:      id.(string),
		ChannelName: deleteChannelReq.ChannelName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// ControlVoice godoc
// @Summary ControlVoice
// @Description Add tv to home
// @Tags television
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param up query string false "true or false"
// @Success 200 {object} entity.DownOrUpVoiceTvRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /speaker/voice [post]
// @Security BearerAuth
func (t *Speaker) ControlVoice(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
	}
	up := c.Query("up")
	message, err := t.speaker.UpVoice(c.Request.Context(), id.(string), up)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"sound": message})
}

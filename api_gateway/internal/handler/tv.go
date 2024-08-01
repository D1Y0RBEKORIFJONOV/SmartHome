package handler

import (
	"api_gate_way/internal/entity"
	"api_gate_way/internal/usecase/tv_use_case"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TV struct {
	tv tv_use_case.TvUseCaseImpl
}

func NewTV(tv tv_use_case.TvUseCaseImpl) *TV {
	return &TV{
		tv: tv,
	}
}

// AddTvHome godoc
// @Summary AddTvHome
// @Description Add tv to home
// @Tags television
// @Accept json
// @Produce json
// @Param body body entity.AddTVReq true "tv create information"
// @Security ApiKeyAuth
// @Success 200 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /tv/register [post]
// @Security BearerAuth
func (t *TV) AddTvHome(c *gin.Context) {
	var addTvReq entity.AddTVReq
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
	message, err := t.tv.AddTV(c.Request.Context(), &entity.AddTVReq{
		ModelName: addTvReq.ModelName,
		UserID:    id.(string),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// OpenTv godoc
// @Summary OpenTv
// @Description Add tv to home
// @Tags television
// @Accept json
// @Produce json
// @Param open query string false "true or false"
// @Security ApiKeyAuth
// @Success 200 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /tv/open [put]
// @Security BearerAuth
func (t *TV) OpenTv(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}
	open := c.Query("open")
	message, err := t.tv.OpenTV(c.Request.Context(), open, id.(string))
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
// @Router /tv/user/channels [get]
// @Security BearerAuth
func (t *TV) GetChannels(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}
	channels, err := t.tv.GetChannels(c.Request.Context(), id.(string))
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
// @Router /tv/channel/cursor [post]
// @Security BearerAuth
func (u *TV) PreviousAndNext(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}
	next := c.Query("next")
	message, err := u.tv.NextChannel(c.Request.Context(), id.(string), next)
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
// @Param body body entity.AddChannelReq true "channel create information"
// @Success 200 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /tv/channel [post]
// @Security BearerAuth
func (t *TV) AddChannel(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}
	var addChannelReq entity.AddChannelReq
	if err := c.ShouldBindJSON(&addChannelReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message, err := t.tv.AddChannel(c.Request.Context(), &entity.AddChannelReq{
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
// @Router /tv/channel [delete]
// @Security BearerAuth
func (t *TV) DeleteChannel(c *gin.Context) {
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
	message, err := t.tv.DeleteChannel(c.Request.Context(), &entity.DeleteChannelReq{
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
// @Router /tv/voice [post]
// @Security BearerAuth
func (t *TV) ControlVoice(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
	}
	up := c.Query("up")
	message, err := t.tv.UpVoice(c.Request.Context(), id.(string), up)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"sound": message})
}

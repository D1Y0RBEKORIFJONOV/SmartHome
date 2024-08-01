package router

import (
	"api_gate_way/internal/handler"
	"api_gate_way/internal/middleware"
	speaker_use_case "api_gate_way/internal/usecase/speaker"
	"api_gate_way/internal/usecase/tv_use_case"
	"api_gate_way/internal/usecase/user_usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRouter(user user_usecase.User, tv tv_use_case.TvUseCaseImpl, speaker speaker_use_case.SpeakerUseCaseImpl) *gin.Engine {
	userHandler := handler.NewUserServer(user)
	tvHandler := handler.NewTV(tv)
	spHandler := handler.NewSpeaker(&speaker)

	router := gin.Default()

	router.Use(middleware.Middleware)
	router.Use(middleware.TimingMiddleware)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", userHandler.Login)
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/verify", userHandler.VerifyUser)
		userGroup.PUT("/update", userHandler.UpdateUser)
		userGroup.PUT("/update/password", userHandler.UpdatePassword)
		userGroup.PUT("/update/email", userHandler.UpdateEmail)
		userGroup.DELETE("/delete", userHandler.DeleteUser)
		userGroup.GET("", userHandler.GetUser)
		userGroup.GET("/all", userHandler.GetAllUser)
	}

	tvGroup := router.Group("/tv")
	{
		tvGroup.POST("/register", tvHandler.AddTvHome)
		tvGroup.PUT("/open", tvHandler.OpenTv)
		tvGroup.GET("/user/channels", tvHandler.GetChannels)
		tvGroup.POST("/channel/cursor", tvHandler.PreviousAndNext)
		tvGroup.POST("/channel", tvHandler.AddChannel)
		tvGroup.DELETE("/channel", tvHandler.DeleteChannel)
		tvGroup.POST("/voice", tvHandler.ControlVoice)
	}

	speakerGroup := router.Group("/speaker")
	{
		speakerGroup.POST("/register", spHandler.AddSpeaker)
		speakerGroup.PUT("/open", spHandler.OpenSpeaker)
		speakerGroup.GET("/user/channels", spHandler.GetChannels)
		speakerGroup.POST("/channel/cursor", spHandler.PreviousAndNext)
		speakerGroup.POST("/channel", spHandler.AddChannel)
		speakerGroup.DELETE("/channel", spHandler.DeleteChannel)
		speakerGroup.POST("/voice", spHandler.ControlVoice)
	}

	return router
}

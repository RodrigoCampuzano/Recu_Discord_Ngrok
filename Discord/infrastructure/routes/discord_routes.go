package routes



import (
    "github.com/gin-gonic/gin"
    "RecuNgrok/Discord/infrastructure/controllers"
)

func SetupWebhookRoutes(router *gin.Engine, handler *controllers.WebhookHandler) {
    router.POST("/webhook", handler.ProcessEvent)
}

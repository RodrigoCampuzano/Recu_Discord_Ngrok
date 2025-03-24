package controllers

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "RecuNgrok/Discord/application/usecases"
    "RecuNgrok/Discord/domain/entities"
)

type WebhookHandler struct {
    eventProcessor *usecases.EventProcessor
}

func NewWebhookHandler(processor *usecases.EventProcessor) *WebhookHandler {
    return &WebhookHandler{eventProcessor: processor}
}

func (wh *WebhookHandler) ProcessEvent(ctx *gin.Context) {
    var event entities.EventData

    if err := ctx.ShouldBindJSON(&event); err != nil {
        log.Println("Error parsing JSON:", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    log.Println("Webhook received:", event)

    channel, message := wh.eventProcessor.HandleEvent(event)
    if channel != "" {
        ctx.JSON(http.StatusOK, gin.H{"channel": channel, "message": message})
    } else {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process event"})
    }
}
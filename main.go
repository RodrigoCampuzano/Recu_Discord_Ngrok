package main

import (
    "os"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "webhook_multi/cmd/app/usecases"
    "webhook_multi/cmd/infra/adapters"
    "webhook_multi/cmd/infra/controllers"
    "webhook_multi/cmd/infra/routes"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("No se pudo cargar el archivo .env")
    }

    devWebhook := os.Getenv("WEBHOOK_DEV")
    testWebhook := os.Getenv("WEBHOOK_TEST")

    if devWebhook == "" || testWebhook == "" {
        log.Fatal("Las variables de entorno WEBHOOK_DEV y/o WEBHOOK_TEST no están definidas")
    }

    router := gin.Default()

    discordClient := adapters.NewDiscordClient(devWebhook, testWebhook)
    eventProcessor := usecases.NewEventProcessor(discordClient)
    webhookHandler := controllers.NewWebhookHandler(eventProcessor)

    routes.SetupWebhookRoutes(router, webhookHandler)

    router.Run(":8080")
}
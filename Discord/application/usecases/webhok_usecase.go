package usecases

import (
    "fmt"
    "RecuNgrok/Discord/domain/entities"
    "RecuNgrok/Discord/infrastructure/repositories"
)

type EventProcessor struct {
    discordClient *repositories.DiscordClient
}

func NewEventProcessor(client *repositories.DiscordClient) *EventProcessor {
    return &EventProcessor{discordClient: client}
}

func (ep *EventProcessor) HandleEvent(event entities.EventData) (string, string) {
    channel, message := ep.determineChannelAndMessage(event)
    if channel != "" {
        ep.discordClient.PostMessage(channel, message)
    }
    return channel, message
}

func (ep *EventProcessor) determineChannelAndMessage(event entities.EventData) (string, string) {
    switch event.Action {
    case "opened":
        return "Development", createPRMessage(event)
    case "ready_for_review":
        return "Development", createReviewMessage(event)
    case "reopened":
        return "Development", createReopenedMessage(event)
    case "synchronize":
        return "Development", createPushMessage(event)
    default:
        if event.Workflow != nil {
            return "Testing", createWorkflowMessage(event)
        }
        return "", ""
    }
}

func createPRMessage(event entities.EventData) string {
    return fmt.Sprintf("🔔 **[Pull Request] Nueva actividad**\n\n📄 **Título:** %s\n🔗 [Ver PR](%s)",
        event.PullRequest.Title, event.PullRequest.URL)
}

func createPushMessage(event entities.EventData) string {
    return fmt.Sprintf("🚀 **[Push] Nuevos cambios subidos**\n\n🔗 [Ver cambios](%s)", event.PullRequest.URL)
}

func createReviewMessage(event entities.EventData) string {
    return fmt.Sprintf("👀 **[Review] Pull Request listo para revisión**\n\n📄 **Título:** %s\n🔗 [Revisar PR](%s)",
        event.PullRequest.Title, event.PullRequest.URL)
}

func createReopenedMessage(event entities.EventData) string {
    return fmt.Sprintf("🔄 **[Reabierto] Pull Request reabierto**\n\n📄 **Título:** %s\n🔗 [Abrir PR](%s)",
        event.PullRequest.Title, event.PullRequest.URL)
}

func createWorkflowMessage(event entities.EventData) string {
    conclusion := event.Workflow.Conclusion
    if conclusion == "" {
        conclusion = "Sin conclusión"
    }

    return fmt.Sprintf("⚡ **[GitHub Actions] Workflow ejecutado**\n\n✅ **Estado:** %s\n📌 **Conclusión:** %s\n🔗 [Ver detalles](%s)",
        event.Workflow.Status, conclusion, event.Workflow.URL)
}
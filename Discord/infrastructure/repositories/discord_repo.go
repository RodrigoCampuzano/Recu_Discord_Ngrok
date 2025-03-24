package repositories
import (
    "bytes"
    "encoding/json"
    "net/http"
    "log"
)

type DiscordClient struct {
    DevWebhook   string
    TestWebhook  string
}

func NewDiscordClient(devWebhook, testWebhook string) *DiscordClient {
    return &DiscordClient{
        DevWebhook:   devWebhook,
        TestWebhook:  testWebhook,
    }
}

func (dc *DiscordClient) PostMessage(channel string, content string) error {
    var webhook string
    switch channel {
    case "Development":
        webhook = dc.DevWebhook
    case "Testing":
        webhook = dc.TestWebhook
    default:
        log.Println("Canal no reconocido:", channel)
        return nil
    }

    payload := map[string]string{"content": content}
    body, _ := json.Marshal(payload)

    resp, err := http.Post(webhook, "application/json", bytes.NewBuffer(body))
    if err != nil {
        log.Printf("Error enviando mensaje a Discord: %v", err)
        return err
    }
    defer resp.Body.Close()

    log.Printf("Mensaje enviado a Discord en %s: %s", channel, content)
    return nil
}

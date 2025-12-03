package http

import (
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/thomasfsr/fitness-app/internal/domain/message"
    "github.com/thomasfsr/fitness-app/internal/usecase"
)

// Simple WhatsApp webhook handler placeholder.
// This supports:
// - GET /webhook?hub.mode=... for verification (WhatsApp Cloud API style)
// - POST /webhook to receive messages (you should adapt JSON parsing to your provider)
func NewWhatsAppHandler(r *gin.Engine, messageUC *usecase.MessageUseCase) {
    r.GET("/webhook", func(c *gin.Context) {
        mode := c.Query("hub.mode")
        token := c.Query("hub.verify_token")
        challenge := c.Query("hub.challenge")
        verifyToken := os.Getenv("WHATSAPP_VERIFY_TOKEN")
        if mode == "subscribe" && token == verifyToken {
            c.String(http.StatusOK, challenge)
            return
        }
        c.String(http.StatusForbidden, "forbidden")
    })

    r.POST("/webhook", func(c *gin.Context) {
        // NOTE: Different providers send different payloads.
        // Adapt this parsing to your provider (WhatsApp Cloud API / Twilio).
        var payload map[string]interface{}
        if err := c.BindJSON(&payload); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Very naive parsing: attempt to extract phone and text if present.
        // For production, parse exact provider schema.
        // Example: store a message with role "user"
        m := &message.Message{
            UserId: 0, // map phone -> user in your code
            Role:   "user",
            Message: "received raw webhook (see payload)",
        }
        _ = messageUC.CreateMessage(m)

        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })
}

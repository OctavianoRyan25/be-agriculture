package chatbot

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"github.com/OctavianoRyan25/be-agriculture/modules/chatbot/request"

	"github.com/labstack/echo/v4"
)

type ChatAI struct{}

// Model chatbot start conversation
var AiPayload = map[string]interface{}{
	"model": "gpt-3.5-turbo",
	"messages": []map[string]string{
		{"role": "system", "content": "Anda seorang ahli dalam bidang pertanian dan perkebunan"},
	},
}

func NewChatAI() *ChatAI {
	return &ChatAI{}
}

func (c *ChatAI) HandleChatCompletion(ctx echo.Context) error {
	if len(AiPayload["messages"].([]map[string]string)) == 1 {
		AiPayload["messages"] = append(AiPayload["messages"].([]map[string]string), map[string]string{"role": "system", "content": "Anda seorang ahli dalam bidang pertanian dan perkebunan"})
	}

	// Parse request body
	var request request.Request
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to parse request body"})
	}

	// add user messages
	AiPayload["messages"] = append(AiPayload["messages"].([]map[string]string), map[string]string{"role": "user", "content": request.Messages[0].Content})

	// payload sended to endpoint
	payload := map[string]interface{}{
		"model":    "gpt-3.5-turbo",
		"messages": []map[string]string{AiPayload["messages"].([]map[string]string)[len(AiPayload["messages"].([]map[string]string))-2], AiPayload["messages"].([]map[string]string)[len(AiPayload["messages"].([]map[string]string))-1]},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to marshal JSON payload"})
	}

	// Request to endpoint
	resp, err := http.Post("https://wgpt-production.up.railway.app/v1/chat/completions", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to connect to chatbot endpoint"})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to read response body"})
	}

	return ctx.JSONBlob(resp.StatusCode, body)
}
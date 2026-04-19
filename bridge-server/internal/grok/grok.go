package grok

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

type Client struct {
	client *openai.Client
}

func NewClient(apiKey string) *Client {
	config := openai.DefaultConfig(apiKey)
	// Point to xAI's base URL as requested
	config.BaseURL = "https://api.x.ai/v1"
	
	client := openai.NewClientWithConfig(config)
	return &Client{
		client: client,
	}
}

func (c *Client) QueryGrok(ctx context.Context, prompt string) (string, error) {
	if c.client == nil {
		return "", fmt.Errorf("grok client is not initialized")
	}

	req := openai.ChatCompletionRequest{
		Model: "grok-beta", // Using grok-beta/grok-2 for the standard API model. Adjust based on exact alias.
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}

	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response choices returned from Grok")
}

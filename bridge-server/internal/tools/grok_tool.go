package tools

import (
	"context"
	"fmt"

	"bridge-server/internal/grok"
	localmcp "bridge-server/internal/mcp"

	"github.com/jackc/pgx/v5/pgxpool"
	mark3labs "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type GrokTool struct {
	client *grok.Client
	dbPool *pgxpool.Pool
}

func NewGrokTool(client *grok.Client, dbPool *pgxpool.Pool) localmcp.ToolProvider {
	return &GrokTool{
		client: client,
		dbPool: dbPool,
	}
}

func (t *GrokTool) Tool() mark3labs.Tool {
	return mark3labs.NewTool("query_grok",
		mark3labs.WithDescription("Sends a prompt to the grok-4.20-reasoning model directly and returns its response."),
		mark3labs.WithString("prompt",
			mark3labs.Required(),
			mark3labs.Description("The exact prompt to be forwarded to Grok."),
		),
	)
}

func (t *GrokTool) Handler() server.ToolHandlerFunc {
	return func(ctx context.Context, request mark3labs.CallToolRequest) (*mark3labs.CallToolResult, error) {
		prompt, ok := request.Params.Arguments["prompt"].(string)
		if !ok {
			return nil, fmt.Errorf("prompt argument is required and must be a string")
		}

		response, err := t.client.QueryGrok(ctx, prompt)
		if err != nil {
			return mark3labs.NewToolResultError(fmt.Sprintf("Grok Error: %v", err)), nil
		}

		return mark3labs.NewToolResultText(response), nil
	}
}

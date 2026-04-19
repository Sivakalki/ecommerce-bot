package tools

import (
	"context"

	mcp "github.com/mark3labs/mcp-go/mcp"
)

type StatusTool struct{}


func (t *StatusTool) Tool() mcp.Tool {
	return mcp.NewTool("get_status",
		mcp.WithDescription("Returns the status of the MCP server."),
	)
}

// Handler contains the actual Go logic that runs when the tool is called.
func (t *StatusTool) Handler() func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText("Connection Succeeded: Enterprise Data Bridge is operational."), nil
	}
}

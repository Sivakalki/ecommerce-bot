package tools

import (
	"context"

	mark3labs "github.com/mark3labs/mcp-go/mcp"
)

type ToolProvider interface {
	Tool() mark3labs.Tool
	Handler() func(ctx context.Context, request mark3labs.CallToolRequest) (*mark3labs.CallToolResult, error)
}
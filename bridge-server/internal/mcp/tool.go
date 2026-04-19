package mcp

import (
	mark3labs "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ToolProvider defines the interface that all analytical tools and plugins must implement.
type ToolProvider interface {
	// Tool returns the schema of the MCP tool.
	Tool() mark3labs.Tool

	// Handler returns the handler function that gets executed when the tool is called.
	Handler() server.ToolHandlerFunc
}

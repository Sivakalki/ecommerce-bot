package tools

import (
	"bridge-server/internal/grok"
	localmcp "bridge-server/internal/mcp"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Registry contains the initialized dependencies required by tools.
type Registry struct {
	GrokClient *grok.Client
	DBPool     *pgxpool.Pool
}

// GetTools returns a slice of all registered ToolProviders.
// 
// Add new analytics tools here by appending them to the slice.
func (r *Registry) GetTools() []localmcp.ToolProvider {
	return []localmcp.ToolProvider{
		NewGrokTool(r.GrokClient, r.DBPool),
		// Add future analytics tools here!
		// e.g. NewAnalyticsTool(r.DBPool),
	}
}

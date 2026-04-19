package resources

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

type LogsResource struct {}

func NewLogsResource() *LogsResource {
	return &LogsResource{}
}

func (r *LogsResource) Resource() mcp.Resource {
	return mcp.NewResource("enterprise://logs","enterprise logs",
		mcp.WithResourceDescription("The last 5 system logs"),
		mcp.WithMIMEType("application/json"),
	)
}

func (r *LogsResource) Handler() func(ctx context.Context, request mcp.ReadResourceRequest) ([]interface{}, error) {
	return func(ctx context.Context, request mcp.ReadResourceRequest) ([]interface{}, error) {
		return []interface{}{
			mcp.TextResourceContents{
				ResourceContents: mcp.ResourceContents{
					URI:      request.Params.URI,
					MIMEType: "application/json",
				},
				Text: `{"logs": ["Log 1: System Boot", "Log 2: DB Connected..."]}`,
			},
		}, nil
	}
}
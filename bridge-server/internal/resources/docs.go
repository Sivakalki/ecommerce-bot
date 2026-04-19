package resources

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

type DocsResource struct {}

func NewDocsResource() *DocsResource {
	return &DocsResource{}
}

func (r *DocsResource) Resource() mcp.Resource {
	return mcp.NewResource("enterprise://docs","enterprise schema",
		mcp.WithResourceDescription("The database schema for the e-commerce bridge"),
		mcp.WithMIMEType("application/json"),
	)
}

func (r *DocsResource) Handler() func(ctx context.Context, request mcp.ReadResourceRequest) ([]interface{}, error) {
	return func(ctx context.Context, request mcp.ReadResourceRequest) ([]interface{}, error) {
		return []interface{}{	
			mcp.TextResourceContents{
				ResourceContents: mcp.ResourceContents{
					URI:      "enterprise://docs",
					MIMEType: "application/json",
				},
				Text:     `{"tables": ["customers", "orders", "products"]}`,
			},
		}, nil
	}
}

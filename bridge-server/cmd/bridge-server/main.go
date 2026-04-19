package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"bridge-server/internal/config"
	"bridge-server/internal/db"
	// "bridge-server/internal/grok"
	"bridge-server/internal/tools"
	"bridge-server/internal/resources"
	"bridge-server/internal/prompts"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	log.SetOutput(os.Stderr)
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	actualDBPool, err := db.NewPostgresPool(ctx, cfg.DatabaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Database connection failed. Analytics tools requiring DB may crash. Error: %v\n", err)
	} else {
		defer actualDBPool.Close()
		fmt.Fprintf(os.Stderr, "Successfully connected to pgvector database.\n")
	}

	// Create MCP Server
	mcpServer := server.NewMCPServer(
		"Enterprise Data Bridge MCP",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithPromptCapabilities(true),
	)

	statusTool := tools.StatusTool{}
	mcpServer.AddTool(statusTool.Tool(), statusTool.Handler())

	if actualDBPool != nil {
		orderDetailsTool := tools.NewOrderDetailsTool(actualDBPool)
		latestOrderTool := tools.NewLatestOrderTool(actualDBPool)
		mostOrderedTool := tools.NewMostOrderedProductTool(actualDBPool)

		mcpServer.AddTool(orderDetailsTool.Tool(), orderDetailsTool.Handler())
		mcpServer.AddTool(latestOrderTool.Tool(), latestOrderTool.Handler())
		mcpServer.AddTool(mostOrderedTool.Tool(), mostOrderedTool.Handler())
	}

	docsResource := resources.DocsResource{}
	logsResource := resources.LogsResource{}
	mcpServer.AddResource(docsResource.Resource(), docsResource.Handler())
	mcpServer.AddResource(logsResource.Resource(), logsResource.Handler())

	explainSQL := prompts.ExplainSQL{}
	mcpServer.AddPrompt(explainSQL.Prompt(), explainSQL.Handler())

	if err := server.ServeStdio(mcpServer); err != nil {
		log.Fatalf("Server closed abruptly: %v\n", err)
	}
}
	
package tools

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	mcp "github.com/mark3labs/mcp-go/mcp"
)

type LatestOrderTool struct {
	dbPool *pgxpool.Pool
}

func NewLatestOrderTool(dbPool *pgxpool.Pool) *LatestOrderTool {
	return &LatestOrderTool{dbPool: dbPool}
}

func (t *LatestOrderTool) Tool() mcp.Tool {
	return mcp.NewTool("latest_order",
		mcp.WithDescription("Get the most recent order from the database"),
	)
}

func (t *LatestOrderTool) Handler() func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query := `SELECT id, user_id, total_amount, status, created_at FROM orders ORDER BY created_at DESC LIMIT 1`
		
		var id, userID int
		var totalAmount float64
		var status string
		var createdAt time.Time

		err := t.dbPool.QueryRow(ctx, query).Scan(&id, &userID, &totalAmount, &status, &createdAt)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to query latest order: %v", err)), nil
		}

		resultStr := fmt.Sprintf("Latest Order ID: %d\nUser ID: %d\nTotal Amount: $%.2f\nStatus: %s\nCreated At: %s",
			id, userID, totalAmount, status, createdAt.Format(time.RFC3339))
		return mcp.NewToolResultText(resultStr), nil
	}
}


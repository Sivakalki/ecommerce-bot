package tools

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	mcp "github.com/mark3labs/mcp-go/mcp"
)

type OrderDetailsTool struct {
	dbPool *pgxpool.Pool
}

func NewOrderDetailsTool(dbPool *pgxpool.Pool) *OrderDetailsTool {
	return &OrderDetailsTool{dbPool: dbPool}
}

func (t *OrderDetailsTool) Tool() mcp.Tool {
	return mcp.NewTool("order_details",
		mcp.WithDescription("Get order details"),
		mcp.WithString("order_id",
			mcp.Required(),
			mcp.Description("Order ID"),
		),
	)
}

func (t *OrderDetailsTool) Handler() func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		orderIDStr, ok := request.Params.Arguments["order_id"].(string)
		if !ok {
			return nil, fmt.Errorf("order_id argument is required and must be a string")
		}

		query := `SELECT id, user_id, total_amount, status, created_at FROM orders WHERE id = $1`
		var id, userID int
		var totalAmount float64
		var status string
		var createdAt time.Time

		err := t.dbPool.QueryRow(ctx, query, orderIDStr).Scan(&id, &userID, &totalAmount, &status, &createdAt)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to query order: %v", err)), nil
		}

		resultStr := fmt.Sprintf("Order ID: %d\nUser ID: %d\nTotal Amount: $%.2f\nStatus: %s\nCreated At: %s",
			id, userID, totalAmount, status, createdAt.Format(time.RFC3339))
		return mcp.NewToolResultText(resultStr), nil
	}
}
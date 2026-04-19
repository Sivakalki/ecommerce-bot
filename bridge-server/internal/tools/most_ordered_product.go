package tools

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	mcp "github.com/mark3labs/mcp-go/mcp"
)

type MostOrderedProductTool struct {
	dbPool *pgxpool.Pool
}

func NewMostOrderedProductTool(dbPool *pgxpool.Pool) *MostOrderedProductTool {
	return &MostOrderedProductTool{dbPool: dbPool}
}

func (t *MostOrderedProductTool) Tool() mcp.Tool {
	return mcp.NewTool("most_ordered_product",
		mcp.WithDescription("Get the product with the highest total ordered quantity"),
	)
}

func (t *MostOrderedProductTool) Handler() func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query := `
			SELECT p.id, p.name, SUM(oi.quantity) as total_quantity
			FROM order_items oi
			JOIN products p ON oi.product_id = p.id
			GROUP BY p.id, p.name
			ORDER BY total_quantity DESC
			LIMIT 1
		`
		var id int
		var name string
		var totalQuantity int

		err := t.dbPool.QueryRow(ctx, query).Scan(&id, &name, &totalQuantity)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to query most ordered product: %v", err)), nil
		}

		resultStr := fmt.Sprintf("Most Ordered Product:\nID: %d\nName: %s\nTotal Quantity Sold: %d", id, name, totalQuantity)
		return mcp.NewToolResultText(resultStr), nil
	}
}
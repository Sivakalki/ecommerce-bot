package prompts

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

type ExplainSQL struct{}

const explainSQLPrompt = `
	You are a SQL expert. Your task is to analyze the provided SQL query and explain it in simple, clear English.

**Input SQL Query:**
{{SQL_QUERY}}

**Your Output Must Include:**
1.  **Purpose:** A one-sentence summary of what the query does.
2.  **Tables Involved:** List the tables being queried.
3.  **Logic Breakdown:** Explain the joins, filters (WHERE clause), and aggregations (GROUP BY) step-by-step.
4.  **Result:** Describe what data the user will see when this query is run.

Keep the explanation concise but thorough.
`

func (e *ExplainSQL) Prompt() mcp.Prompt {
	prompt := mcp.NewPrompt("explain_sql",
		mcp.WithPromptDescription("Ask the AI to explain a complex SQL query"),
	)

	prompt.Arguments = []mcp.PromptArgument{
		 {
            Name:        "query",
            Description: "The SQL query to explain",
            Required:    true,
        },
	}

	return prompt
}

func (e *ExplainSQL) Handler() func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	return func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		query, ok := request.Params.Arguments["query"]
		if !ok || query == "" {
			return nil, fmt.Errorf("argument 'query' is required")
		}

		finalPrompt := strings.Replace(explainSQLPrompt, "{{SQL_QUERY}}", query, 1)

		return &mcp.GetPromptResult{
			Description: "SQL explanation prompt",
			Messages: []mcp.PromptMessage{
				mcp.NewPromptMessage(mcp.RoleUser, mcp.NewTextContent(
					finalPrompt,
				)),
			},
		}, nil
		
	}
}
	
import { Router, Request, Response } from 'express';
import OpenAI from 'openai';
import { getMcpClient } from '../mcp/client.js';
import { groqClient, GROQ_MODEL, SYSTEM_PROMPT } from '../llm/groq.js';
import { mapMcpToolsToOpenAI } from '../llm/tools.js';

const router = Router();

router.post('/', async (req: Request, res: Response): Promise<void> => {
  const { prompt } = req.body as { prompt?: string };

  if (!prompt?.trim()) {
    res.status(400).json({ error: 'Request body must contain a non-empty "prompt" string.' });
    return;
  }

  const mcpClient = getMcpClient();
  if (!mcpClient) {
    res.status(503).json({ error: 'MCP server is not connected. Please restart the middleman.' });
    return;
  }

  try {
    // 1. Retrieve available tools from the Go MCP server
    const { tools: mcpTools } = await mcpClient.listTools();
    const openAITools = mapMcpToolsToOpenAI(mcpTools);

    console.log(`[CHAT] Prompt: "${prompt}" | Tools available: ${openAITools.length}`);

    // 2. First LLM call — let Groq decide which tool to use
    const messages: OpenAI.Chat.ChatCompletionMessageParam[] = [
      { role: 'system', content: SYSTEM_PROMPT },
      { role: 'user', content: prompt },
    ];

    const firstResponse = await groqClient.chat.completions.create({
      model: GROQ_MODEL,
      messages,
      tools: openAITools,
      tool_choice: 'auto',
    });

    const firstChoice = firstResponse.choices[0];
    const assistantMessage = firstChoice.message;

    // 3. If the LLM wants to call a tool
    if (firstChoice.finish_reason === 'tool_calls' && assistantMessage.tool_calls?.length) {
      const toolCall = assistantMessage.tool_calls[0];

      if (toolCall.type !== 'function') {
        res.status(500).json({ error: `Unsupported tool_call type: ${toolCall.type}` });
        return;
      }

      const toolName = toolCall.function.name;
      const toolArgs = JSON.parse(toolCall.function.arguments || '{}') as Record<string, unknown>;

      console.log(`[CHAT] Tool selected: "${toolName}"`, toolArgs);

      // 4. Execute the tool on the Go MCP server
      const toolResult = await mcpClient.callTool({ name: toolName, arguments: toolArgs });

      const resultContent = (toolResult.content as Array<{ type: string; text?: string }>)
        .filter((c) => c.type === 'text')
        .map((c) => c.text ?? '')
        .join('\n');

      console.log(`[CHAT] Tool result: ${resultContent}`);

      // 5. Second LLM call — summarize the tool result
      const secondResponse = await groqClient.chat.completions.create({
        model: GROQ_MODEL,
        messages: [
          ...messages,
          assistantMessage,
          {
            role: 'tool',
            tool_call_id: toolCall.id,
            content: resultContent,
          },
        ],
      });

      const finalText = secondResponse.choices[0].message.content ?? 'No response generated.';
      res.json({ response: finalText, tool_used: toolName });
    } else {
      // LLM answered directly without needing a tool
      const directText = assistantMessage.content ?? 'No response generated.';
      res.json({ response: directText, tool_used: null });
    }
  } catch (err) {
    console.error('[CHAT] Error:', err);
    res.status(500).json({ error: 'Internal server error during chat processing.' });
  }
});

export default router;

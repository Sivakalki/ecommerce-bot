import { Client } from '@modelcontextprotocol/sdk/client/index.js';
import OpenAI from 'openai';
import type { FunctionParameters } from 'openai/resources/shared.js';

type McpTools = Awaited<ReturnType<Client['listTools']>>['tools'];

export function mapMcpToolsToOpenAI(mcpTools: McpTools): OpenAI.Chat.ChatCompletionTool[] {
  return mcpTools.map((tool) => ({
    type: 'function' as const,
    function: {
      name: tool.name,
      description: tool.description ?? '',
      parameters: (tool.inputSchema ?? {
        type: 'object',
        properties: {},
      }) as FunctionParameters,
    },
  }));
}

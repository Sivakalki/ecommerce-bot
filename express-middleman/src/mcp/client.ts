import { Client } from '@modelcontextprotocol/sdk/client/index.js';
import { StdioClientTransport } from '@modelcontextprotocol/sdk/client/stdio.js';
import { env } from '../config/env.js';

let mcpClient: Client | null = null;
let mcpTransport: StdioClientTransport | null = null;

export function getMcpClient(): Client | null {
  return mcpClient;
}

export async function connectMCP(): Promise<void> {
  console.log(`[MCP] Spawning Go binary: ${env.GO_BINARY_PATH}`);

  mcpTransport = new StdioClientTransport({
    command: env.GO_BINARY_PATH,
    args: [],
  });

  // Gracefully handle binary crashes — never hang the server
  mcpTransport.onerror = (err) => {
    console.error('[MCP] Transport error:', err.message);
    mcpClient = null;
  };

  mcpTransport.onclose = () => {
    console.warn('[MCP] Transport closed. Go binary may have exited.');
    mcpClient = null;
    mcpTransport = null;
  };

  mcpClient = new Client(
    { name: 'express-middleman', version: '1.0.0' },
    { capabilities: {} }
  );

  await mcpClient.connect(mcpTransport);
  console.log('[MCP] Connected to Go MCP server successfully.');
}

export async function disconnectMCP(): Promise<void> {
  if (mcpTransport) {
    await mcpTransport.close();
    mcpClient = null;
    mcpTransport = null;
    console.log('[MCP] Disconnected cleanly.');
  }
}

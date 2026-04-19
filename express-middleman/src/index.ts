import { createApp } from './app.js';
import { env } from './config/env.js';
import { connectMCP } from './mcp/client.js';

async function setup(): Promise<void> {
  // 1. Connect to Go MCP server
  try {
    await connectMCP();
  } catch (err) {
    console.error('[STARTUP] Failed to connect to MCP server:', err);
    console.warn('[STARTUP] Starting Express without MCP. /chat will return 503.');
  }

  // 2. Boot Express
  const app = createApp();

  app.listen(env.PORT, () => {
    console.log(`[SERVER] Express Middleman listening on http://localhost:${env.PORT}`);
  });
}

setup();

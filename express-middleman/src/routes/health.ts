import { Router, Request, Response } from 'express';
import { getMcpClient } from '../mcp/client.js';

const router = Router();

router.get('/', (_req: Request, res: Response) => {
  const client = getMcpClient();
  res.json({
    status: 'ok',
    mcpConnected: client !== null,
    timestamp: new Date().toISOString(),
  });
});

export default router;

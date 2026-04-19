import express from 'express';
import cors from 'cors';
import { env } from './config/env.js';
import healthRouter from './routes/health.js';
import chatRouter from './routes/chat.js';

export function createApp(): express.Express {
  const app = express();

  // Middleware
  app.use(cors({ origin: env.CORS_ORIGIN }));
  app.use(express.json());

  // Routes
  app.use('/health', healthRouter);
  app.use('/chat', chatRouter);

  return app;
}

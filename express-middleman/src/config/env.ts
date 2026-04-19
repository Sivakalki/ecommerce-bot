import path from 'path';
import 'dotenv/config';

function require_env(key: string): string {
  const val = process.env[key];
  if (!val) {
    console.error(`[CONFIG] Fatal: environment variable "${key}" is not set.`);
    process.exit(1);
  }
  return val;
}

export const env = {
  GROQ_API_KEY: require_env('GROQ_API_KEY'),
  GO_BINARY_PATH: path.resolve(
    process.env.GO_BINARY_PATH ?? 'Your_path_here'
  ),
  PORT: parseInt(process.env.PORT ?? '3001', 10),
  CORS_ORIGIN: process.env.CORS_ORIGIN ?? 'http://localhost:3000',
} as const;

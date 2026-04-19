# 🚀 MCP-Enterprise-Bridge

**An AI-Native E-commerce Intelligence Layer built on the Model Context Protocol (MCP).**

This repository provides a complete end-to-end implementation of an AI agent capable of querying an e-commerce database using natural language. It leverages **Go** for performance, **Node.js** for AI orchestration, and **PostgreSQL/pgvector** for intelligent data retrieval.

---

## 🏗️ Architecture
The project follows a decoupled, three-tier architecture:
1.  **Go MCP Server**: Interacts directly with PostgreSQL and exposes database tools.
2.  **Express Middleman (MCP Client)**: Orchestrates LLM (Groq/OpenAI) logic and spawns the MCP binary.
3.  **Dashboard UI**: A React-based frontend for the end-user chat experience.

---

## 🛠️ Getting Started

### 1. Database Setup (Docker)
Ensure Docker is running, then initialize the database and pgvector extension:
```bash
docker-compose up -d
```

### 2. Configure Environment Variables
Create a `.env` file in the `express-middleman/` directory:

**`express-middleman/.env`**
```env
# Groq API Key (Get it free: https://console.groq.com)
GROQ_API_KEY=your_gsk_key_here

# Path to the compiled Go MCP binary (Absolute path recommended)
GO_BINARY_PATH=D:\MyPro\ecomerce-mcp\bridge-server\bridge.exe

# Database Connection String (Matches docker-compose config)
DATABASE_URL=postgres://admin:password@localhost:5433/ecommerce
```

### 3. Build & Verify the MCP Server
Move to the server folder, compile the binary, and optionally verify tools using the MCP Inspector:
```bash
cd bridge-server
go build -o bridge.exe ./cmd/bridge/main.go

# Optional: Verify tools are registered correctly
npx @modelcontextprotocol/inspector ./bridge.exe
```

### 4. Start the Middleman (MCP Client)
Open a new terminal and start the backend bridge:
```bash
cd express-middleman
npm install
npm run dev
```

### 5. Launch the Dashboard
Open a new terminal and start the frontend:
```bash
cd dashboard-ui
npm install
npm run dev
```
Visit `http://localhost:3000` to start chatting with your database!

---

## 💎 Use Cases
* **E-commerce Support**: Automate order status checks and product availability queries.
* **Analytical Platforms**: Connect complex SQL databases to an LLM for instant data visualization and reporting.
* **Semantic Search**: Use pgvector to find products based on "vibes" or descriptions rather than just keywords.

---

## 🛠️ Current Status & Roadmap
Currently, I have implemented **3 core tools** for database interaction. I am actively updating this repo daily with:
- [ ] Advanced vector similarity search tools.
- [ ] Multi-turn conversation memory.
- [ ] Support for more complex relational queries.

**I am building this to help developers bridge the gap between static databases and dynamic AI agents.**

---

## 🤝 Connect With Me
If you're interested in building MCP-native applications for your business or have any questions about this architecture, let's connect!

📧 **Email:** [sivakalkipusarla6@gmail.com](mailto:sivakalkipusarla6@gmail.com)

---

### ⭐ Show your support
If you find this project helpful for your learning or work, please **give it a star!** It keeps me motivated to keep updating the tools daily.

# Aurelian Server

Aurelian Server provides a chat interface over any LLM inference backend. It is meant for larger organizations with many concurrent users. The hope is that organizations can maintain custody over their data while taking advantage of LLMs.

If you are an individual hoping to use LLMs privately, check out [Aurelian-Desktop](github.com/aurelianai/Aurelian-Desktop). It can run on a laptop without a GPU.

# Development

1. `/frontend` - [SvelteKit](kit.svelte.dev) SPA
2. `/backend` - Golang API backend: persists chats in a PostgreSQL database.
3. `/backend_tests` - Python testing suite that probes API routes for expected behavior

To bring up a local instance of the backend and database, run

`docker compose up -d`

To spawn a development server for the frontend, run

`npm run dev`

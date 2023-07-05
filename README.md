# Aurelian Enterprise Language Server

Welcome to the AELS monorepo. We use sveltekit to build an SPA frontend and then go-fiber for a backend/server.


# Development

I separate the Svelte Kit and go into two separate containers. There are three Dockerfiles in this repo
1. `Dockerfile.production`, a three stage build that builds the fiber api, builds the svelte SPA, then has a final alpine stage for deployment
2. `/backend/Dockerfile`, a development container for the go api. It uses `air` to watch for changes and rebuild. It proxies traffic to the node server.

   **Important!!** In `backend/main.go`, proxies for all routes of the SPA must be registered. Even though they all match to the same html, this is required for the traffic to make it to the svelte dev server.

3. `/frontend/Dockerfile`, a development container that watches for changes with `npm run dev`. Recieves traffic through go dev container 

The development system can be launched with `docker compose up -d`

The production container is built with `docker build -t aurelian-enterprise-language-server -f Dockerfile.production`

It should be noted that the fiber api is the only one exposed at `localhost:2140` and proxies requests for `index.html` to the frontend through 
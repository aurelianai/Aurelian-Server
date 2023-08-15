# Aurelian Enterprise Language Server

Welcome to the AELS monorepo. It's a [sveltekit](kit.svelte.dev) SPA and a [fiber](gofiber.io) + [gorm](gorm.io) REST API/webserver.

# Development

I separate the frontend and backend into two separate containers. There are three Dockerfiles in this repo
1. `Dockerfile.production`, a three stage build that builds the api, builds the svelte SPA, then has a final alpine stage for deployment.
2. `/backend/Dockerfile`, a development container for the go api. It uses `air` to watch for changes and rebuild.
3. `/frontend/Dockerfile`, a development container that watches for changes with `npm run dev`.

- Nginx is used to split traffic as would be expected between the static files and other routes in production. `nginx.conf` specifies the configuration.

The development system can be launched with `docker compose up -d` and the dev service is available at [localhost:8080](localhost:8080)

Although it's a bit complex, it allows for a fast pace of development due to near immediate loading of changes both the backend and frontend.

# Production
To bring up a local instance of the production containers 

```docker compose up -f docker-compose.production.yml```

It will be available at `localhost:2140`
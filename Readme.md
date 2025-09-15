# This is the API for the unitycode web app

## How the routes are laid out
- /status -> show current status of app
- /register -> register a user
- /login -> login with existing user data

## How to start docker container (with their respectable env file)
localhost development:
- docker compose --env-file .env.dev up
docker environment:
- docker compose --env-file .env.prod up

### Rebuild only one container from the docker compose
- docker compose up -d --build <service-name>

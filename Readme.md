# Realtime Chat App with React and Go

## Develop

### Backend:

1. Setup enviroment variables.

```bash
cp .env.example .env

```

2. Start docker services.

```bash
docker run --name pg-chat -e POSTGRES_PASSWORD=example -p 5432:5432 -d postgres

docker run --name some-redis -p 6379:6379 -d redis
```

3. Start Go App.

```bash
go run ./cmd/chat/main.go
```

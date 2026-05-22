# Caddy Reverse Proxy & Load Balancer with Redis Rate Limiting

A hands-on project exploring how to use **Caddy** as a reverse proxy that load balances across multiple server instances with centralized rate limiting backed by **Redis**.

---

## Architecture

```
Internet → Caddy → [ Instance 1 | Instance 2 | Instance 3 | Frontend ]
                         ↕
                       Redis
                  (Rate Limit Store)
```

---

## Prerequisites

- [Docker](https://www.docker.com/)
- [Go](https://golang.org/)

---

## Getting Started

### 1. Start Caddy and Redis

```bash
docker compose up -d
```

### 2. Start the Server Instances

Run each of the following in a separate terminal:

```powershell
# Instance 1
$env:NAME="app1"; $env:PORT="8080"; go run main.go

# Instance 2
$env:NAME="app2"; $env:PORT="8081"; go run main.go

# Instance 3
$env:NAME="app3"; $env:PORT="8082"; go run main.go

# Frontend
$env:NAME="frontend"; $env:PORT="3000"; go run main.go
```

> **Note:** The server instances will be added to the `docker-compose.yml` soon so everything can be started with a single command.

---

## Features

- Caddy as a reverse proxy and load balancer
- Traffic distributed across multiple backend instances
- Centralized rate limiting with Redis (shared across all instances)

---

## Troubleshooting

### Caddy not routing traffic to backend instances

- Verify all server instances are running and listening on the correct ports
- Check Caddy's logs: `docker compose logs caddy`
- Ensure the `Caddyfile` has the correct upstream addresses

### Rate limiting not working

- Confirm Redis is running: `docker compose ps redis`
- Check Redis connectivity from Caddy container: `docker exec caddy caddy list-modules`
- Verify the `rate_limit` directive in `Caddyfile` is correctly configured

### Connection refused errors

- Ensure no firewall is blocking the configured ports
- Verify port availability: `netstat -tlnp | grep <port>`
- Check that Docker containers are using the correct port mappings

### Docker containers failing to start

- Check Docker logs: `docker compose logs`
- Ensure Docker and Docker Compose are installed and running
- Try restarting: `docker compose down && docker compose up -d`


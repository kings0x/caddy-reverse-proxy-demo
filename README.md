# Caddy Reverse Proxy & Load Balancer with Redis Rate Limiting

A hands-on project exploring how to use **Caddy** as a reverse proxy that load balances across multiple server instances with centralized rate limiting backed by **Redis**.

---

## Architecture

``` mermaid
sequenceDiagram
    participant Client as Client
    participant Caddy as Caddy Reverse Proxy
    participant Redis as Redis Rate Limiter
    participant App1 as App Instance 1
    participant App2 as App Instance 2
    participant App3 as App Instance 3

    Client->>Caddy: HTTP Request
    Caddy->>Redis: Check Rate Limit
    Redis-->>Caddy: Allow/Deny
   alt Allowed
        Caddy->>App1: Forward (LB: Round Robin)
        Caddy->>App2: Forward
        Caddy->>App3: Forward
        App1-->>Client: Response
    else Denied
        Caddy-->>Client: 429 Too Many Requests
    end
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


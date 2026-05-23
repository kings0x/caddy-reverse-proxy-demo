# Caddy Reverse Proxy & Load Balancer with Redis Rate Limiting


## Quick Start

Install with:
```bash
pip install caddy-reverse-proxy-demo
```

Or clone and run:
```bash
git clone https://github.com/kings0x/caddy-reverse-proxy-demo.git
cd caddy-reverse-proxy-demo
python setup.py install
```


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


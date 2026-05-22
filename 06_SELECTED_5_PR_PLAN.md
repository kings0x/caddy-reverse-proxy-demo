# 06_SELECTED_5_PR_PLAN.md — caddy-reverse-proxy-demo

## Selected 5 PRs

### PR-1: Add Go server instances to docker-compose.yml
**Target**: `docker-compose.yml`
**Priority**: P0 (promised feature, major usability gap)
**Summary**: Add app1, app2, app3, and frontend as services so entire stack starts with `docker compose up -d`.

**Changes**:
```yaml
services:
  app1:
    image: golang:alpine
    working_dir: /app
    command: go run main.go
    environment:
      - NAME=app1
      - PORT=8080
      - HOST=0.0.0.0
    volumes:
      - .:/app
    network_mode: host
    depends_on:
      - redis

  app2:
    image: golang:alpine
    working_dir: /app
    command: go run main.go
    environment:
      - NAME=app2
      - PORT=8081
      - HOST=0.0.0.0
    volumes:
      - .:/app
    network_mode: host
    depends_on:
      - redis

  app3:
    image: golang:alpine
    working_dir: /app
    command: go run main.go
    environment:
      - NAME=app3
      - PORT=8082
      - HOST=0.0.0.0
    volumes:
      - .:/app
    network_mode: host
    depends_on:
      - redis

  frontend:
    image: golang:alpine
    working_dir: /app
    command: go run main.go
    environment:
      - NAME=frontend
      - PORT=3000
      - HOST=0.0.0.0
    volumes:
      - .:/app
    network_mode: host
    depends_on:
      - redis
```

**Verification**: `docker compose up -d` → all 5 services running → `curl http://localhost/health` returns healthy response.

---

### PR-2: Add MIT LICENSE
**Target**: `LICENSE`
**Priority**: P1 (legal requirement for open source)
**Summary**: Add MIT License so project is legally reusable.

**Changes**: Create `LICENSE` with MIT License text.

**Verification**: File exists, content is valid MIT License.

---

### PR-3: Fix Caddyfile rate_limit syntax
**Target**: `caddy/Caddyfile`
**Priority**: P1 (rate limiting may be broken due to syntax issues)
**Summary**: Fix the `order rate_limit before basicauth` directive and nested rate_limit blocks to use valid Caddy 2 syntax.

**Changes**: Review Caddyfile — the `order rate_limit` global option is unusual; the nested `rate_limit { zone ... }` blocks inside `handle @auth` and `handle @normal` need verification against current Caddy 2 + caddy-ratelimit plugin syntax. May need restructuring to proper directive placement.

**Verification**: Build custom Caddy image → `docker compose up` → rate limiting works (test with >100 requests in window).

---

### PR-4: Fix targets.txt formatting
**Target**: `targets.txt`
**Priority**: P2 (broken developer tool)
**Summary**: Properly format requests on separate lines.

**Changes**:
```
GET http://localhost:80/api/register

GET http://localhost:80/api/buy

POST http://localhost:80/health
Content-Type: application/json
```

**Verification**: File readable and each request is distinguishable.

---

### PR-5: Add Go .gitignore
**Target**: `.gitignore`
**Priority**: P2 (standard practice)
**Summary**: Add standard Go project gitignore.

**Changes**:
```
# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test artifacts
*.test
*.out

# IDE
.idea/
.vscode/

# Docker
.docker/
```

**Verification**: `git status` after `go build` doesn't show binary files.

---

## Implementation Order
1. LICENSE (independent, P1)
2. .gitignore (independent, P2)
3. targets.txt fix (independent, P2)
4. docker-compose.yml server instances (largest change, P0)
5. Caddyfile rate_limit fix (requires testing knowledge, P1)

## Out of Scope
- PR-6 (Redis healthcheck) — deferred
- PR-7 (README improvements) — deferred
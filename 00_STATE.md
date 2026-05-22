# 00_STATE.md — caddy-reverse-proxy-demo

## Repository Identity
- **Repo**: kings0x/caddy-reverse-proxy-demo
- **Fork**: okwn/caddy-reverse-proxy-demo (forked 2026-05-22)
- **Language**: Go
- **Archived**: false
- **License**: None (no license set)
- **Default branch**: main

## Fork Status
- Fork created: ✅ (2026-05-22)
- Upstream sync: ✅ (origin/main == upstream/main — identical commits)
- 0 open issues, 0 open PRs

## Local Working Copy
- **Path**: `/root/oss-pr-campaign/repos/caddy-reverse-proxy-demo`
- **Current branch**: main (on origin/okwn fork)
- **Upstream remote**: `upstream` → https://github.com/kings0x/caddy-reverse-proxy-demo

## Project Type
- Demo/tutorial project
- Infrastructure-as-Code (Docker + Caddy)
- Go backend server instances

## Architecture
```
Internet → Caddy (reverse proxy + load balancer) → [ app1:8080 | app2:8081 | app3:8082 | frontend:3000 ]
                                    ↕
                                  Redis (rate limit storage)
```

## Key Files
| File | Purpose |
|------|---------|
| `main.go` | Go HTTP server (configurable via NAME/PORT/HOST env vars) |
| `docker-compose.yml` | Redis + Caddy services |
| `caddy/Caddyfile` | Caddy reverse proxy config with rate limiting |
| `Dockerfile.caddy` | Custom Caddy build with rate_limit + redis plugins |
| `go.mod` | Go module (go 1.25.3) |
| `targets.txt` | HTTP test targets (single line, looks corrupted) |

## Sync Timeline
- Created fork at 16:12 UTC
- Both repos show identical 2 commits: `0e2f383`, `5efac28`
- No divergent branches

## Notes
- No license — default copyright applies, no reusable open source license
- No CI/CD workflows present
- targets.txt appears to have all content on one line (formatting issue)
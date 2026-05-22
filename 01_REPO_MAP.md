# 01_REPO_MAP.md — caddy-reverse-proxy-demo

## Upstream (kings0x/caddy-reverse-proxy-demo)
```
https://github.com/kings0x/caddy-reverse-proxy-demo
├── README.md              # Project overview, architecture, getting started
├── go.mod                 # Module: github.com/kings0x/caddy-reverse-proxy-demo, go 1.25.3
├── main.go                # Go HTTP server (env: NAME, PORT, HOST)
├── docker-compose.yml     # Services: redis:7.2, caddy (custom build)
├── Dockerfile.caddy       # Custom Caddy build with rate_limit + redis plugins
├── caddy/
│   └── Caddyfile          # Reverse proxy config, rate limiting rules, load balancing
└── targets.txt            # HTTP test targets (appears malformed)
```

## Fork (okwn/caddy-reverse-proxy-demo)
- Same as upstream (forked 2026-05-22, no new commits)
- Remote: `origin` → https://github.com/okwn/caddy-reverse-proxy-demo

## Git Remotes
| Name | URL |
|------|-----|
| origin | https://github.com/okwn/caddy-reverse-proxy-demo (fork) |
| upstream | https://github.com/kings0x/caddy-reverse-proxy-demo (parent) |

## Commit History
| Commit | Message |
|--------|---------|
| `5efac28` | fixed the readme file |
| `0e2f383` | caddy reverse proxy setup with redis rate limiting and load balancing |

## Branches
| Branch | Tracking | Ahead | Behind |
|--------|----------|-------|--------|
| main | origin/main | 0 | 0 |
| upstream/main | upstream/main | 0 | 0 |

## Code Locations
- **Go server**: `main.go` — simple mux with `/`, `/api/auth/*`, `/api/buy`, `/health`
- **Caddy config**: `caddy/Caddyfile` — routes `/api/auth/*`, `/api/*` (non-frontend) through rate limiter → reverse proxy to 3 backends
- **Docker compose**: `docker-compose.yml` — defines redis + caddy services
- **Custom Caddy build**: `Dockerfile.caddy` — adds `caddy-ratelimit` + `caddy-storage-redis` plugins

## Known Issues
- No license file
- targets.txt has formatting issue (all content on single line)
- README note says server instances will be added to docker-compose.yml "soon" — still not done
- rate_limit directive in Caddyfile may have deprecated syntax (order + nested blocks)
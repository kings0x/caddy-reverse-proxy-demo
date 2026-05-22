# 05_PR_CANDIDATES.md — caddy-reverse-proxy-demo

## Quality Audit Findings

### Issues Identified

| # | Category | Issue | Severity |
|---|----------|-------|----------|
| 1 | Documentation | README promises docker-compose server instances "soon" but never delivered | Medium |
| 2 | Configuration | Caddyfile rate_limit syntax appears deprecated/invalid for Caddy 2 | High |
| 3 | DevX | targets.txt is malformed (all content on single line) | Low |
| 4 | Legal | No LICENSE file — default copyright applies, not reusable | Medium |
| 5 | DevX | No .gitignore — compiled binaries, IDE files not excluded | Low |
| 6 | Infrastructure | docker-compose.yml doesn't include Go server instances (incomplete) | Medium |
| 7 | Infrastructure | No healthcheck on Redis service | Low |

---

## PR Candidate 1: Add Go server instances to docker-compose.yml
**File**: `docker-compose.yml`
**Problem**: README explicitly promises this feature ("The server instances will be added to the `docker-compose.yml` soon so everything can be started with a single command.") but it's still missing. Users must run 4 separate terminal commands.
**Fix**: Add app1, app2, app3, and frontend as services in docker-compose.yml so `docker compose up -d` starts everything.
**Impact**: Major usability improvement — single command to start entire stack.

---

## PR Candidate 2: Fix Caddyfile rate_limit directive syntax
**File**: `caddy/Caddyfile`
**Problem**: The `order rate_limit before basicauth` global option and nested `rate_limit { zone ... }` blocks inside `handle` blocks may use deprecated Caddy 2 syntax. The correct Caddy 2 syntax for rate limiting uses `@matchers` and directives differently.
**Fix**: Review and update rate_limit directive to use correct Caddy 2 syntax (likely needs `handle @auth { rate_limit ... }` pattern adjustment or middleware ordering fix).
**Impact**: Ensures rate limiting actually works as documented.

---

## PR Candidate 3: Fix targets.txt formatting
**File**: `targets.txt`
**Problem**: All content is on a single line (`GET http://localhost:80/api/registerGET http://localhost:80/api/buyPOST http://localhost:80/healthContent-Type: application/json`) making it unusable for testing.
**Fix**: Properly format each request on its own line with blank lines between requests.
```
GET http://localhost:80/api/register

GET http://localhost:80/api/buy

POST http://localhost:80/health
Content-Type: application/json
```
**Impact**: Makes the file actually useful for manual/API testing.

---

## PR Candidate 4: Add LICENSE file (MIT)
**File**: `LICENSE`
**Problem**: No license specified. Default copyright applies — others cannot legally use this code.
**Fix**: Add MIT License file.
**Impact**: Makes the project legally reusable by others.

---

## PR Candidate 5: Add .gitignore for Go project
**File**: `.gitignore`
**Problem**: No .gitignore — compiled binaries, vendor directories, IDE files would be tracked.
**Fix**: Add standard Go .gitignore:
```
*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out
.idea/
.vscode/
```
**Impact**: Keeps repo clean, prevents accidental binary commits.

---

## PR Candidate 6: Add Redis healthcheck
**File**: `docker-compose.yml`
**Problem**: Redis service has no healthcheck to ensure it's ready before Caddy starts.
**Fix**: Add `healthcheck` to redis service.
**Impact**: More robust container orchestration.

---

## PR Candidate 7: Improve README with troubleshooting and better structure
**File**: `README.md`
**Problem**: README is minimal — missing troubleshooting, tested commands, environment variable docs.
**Fix**: Add sections: Prerequisites with versions, Troubleshooting (common issues), Environment variables reference, Testing guide using targets.txt.
**Impact**: Better onboarding for new contributors/users.

---

## Evaluation Summary

| Priority | Candidate | Reason |
|----------|-----------|--------|
| 1 | #1 (docker-compose servers) | Promised feature, high impact |
| 2 | #2 (Caddyfile syntax) | Functional bug, rate limiting may not work |
| 3 | #4 (LICENSE) | Legal requirement for OSS reuse |
| 4 | #7 (README improvements) | Low effort, high value |
| 5 | #3 (targets.txt) | Simple formatting fix |
| 6 | #5 (.gitignore) | Standard practice |
| 7 | #6 (Redis healthcheck) | Nice to have |
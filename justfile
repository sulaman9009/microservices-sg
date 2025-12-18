# justfile: developer tasks and commands

# Run a single Go service by name (e.g. `just service posts`).
service target:
	@sh -lc 'if [ -z "{{target}}" ]; then echo "usage: just service <name>"; exit 1; fi; SVC_DIR="services/{{target}}/cmd/server"; if [ -d "$SVC_DIR" ]; then echo "Starting service: {{target}} in $SVC_DIR"; (cd "$SVC_DIR" && if command -v air >/dev/null 2>&1; then air; else go run .; fi); else echo "Service not found: services/{{target}}"; exit 1; fi'

# Run the web app in `web` using pnpm.
web:
	@echo "Starting web app...";
	pnpm --prefix web run dev

# Dev task: start all services (with air or fallback to go run) in the background
# and start the frontend dev server in the foreground.
dev:
	@echo "Starting all Go services in services/* using 'air' (or 'go run' fallback)";
	@sh -lc 'for d in services/*; do if [ -d "$d/cmd/server" ]; then echo "-> $d"; (cd "$d/cmd/server" && if command -v air >/dev/null 2>&1; then air & else go run . & fi); fi; done'
	@echo "Starting frontend (pnpm) in ./web";
	pnpm --prefix web run dev

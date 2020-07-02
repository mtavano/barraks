APP=main
BIN=$(PWD)/$(APP)
GO ?= go

run r:
	@echo "[RUN] Running [barraks]"
	@$(GO) run cmd/main.go

run-bin rb:
	@echo "[RUN] Running bin [barraks]"
	@./main

build b: clean
	@echo "[BUILD] Building [barraks]..."
	@cd cmd && $(GO) build -o $(BIN)

seed s:
	@echo "[SEED] Seeding [barraks]..."
	@cd cmd/seed && $(GO) run main.go

migrate-up m:
	@echo "[MIGRATE] Migrating up [barraks]..."
	@cd database/migrations && goose sqlite3 ../../cmd/barraks.db up


clean c:
	@echo "[CLEAN] Cleaning files..."
	@rm -f $(BIN)


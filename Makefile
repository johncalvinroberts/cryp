FE_DIR=internal/ui
BE_CLI_ENTRYPOINT=cmd/server/main.go
AIR_BIN=./bin/air
BIN=./bin/cryp


build: build-fe build-be

install: install-be install-fe install-air

dev: 
	make -j 2 dev-fe dev-be

install-air: 
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

install-be: 
	go mod tidy

build-be:
	go build -o $(BIN) $(BE_ENTRYPOINT)

dev-be:
	$(AIR_BIN) $(BE_ENTRYPOINT)

build-fe:
	cd $(FE_DIR); pnpm run build;

install-fe:
	cd $(FE_DIR); pnpm install;

dev-fe:
	cd $(FE_DIR); pnpm run dev;


run:
	$(BIN)

clean:
	rm -rf $(BIN)
	rm -rf bin/main
	rm -rf $(FE_DIR)/build

docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-restart: docker-down docker-up

fmt:
	@gofmt -l -w internal cmd

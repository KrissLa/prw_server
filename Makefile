.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: test
test:
	go test ./tests/...


.PHONY: run
run:
	go run ./app/src/main.go


.PHONY: build
build:
	go build -o ./build/builded_app ./app/src/main.go


.PHONY: fmt
fmt:
	go fmt ./...


.PHONY: up
up:
	docker-compose up --build --remove-orphans --abort-on-container-exit -t 0 app


.PHONY: down
down:
	docker-compose down --remove-orphans -t 0
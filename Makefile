.PHONY: help build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## build docker image to deploy
	docker build -t ysaito8015/gotodo:${DOCKER_TAG} \
		--target deploy ./


build-local: ## build docker image to local development
	docker compose build --no-cache

up: ## start docker compose
	docker compose up -d

down: ## stop docker compose
	docker compose down

logs: ## show logs
	docker compose logs -f

ps: ## show docker compose status
	docker compose ps

test: ## run test
	go test -race -shuffle=on ./...

help: ## show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'


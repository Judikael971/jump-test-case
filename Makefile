COMPOSE_BASE_COMMAND := docker compose --env-file .env -f build/package/compose.yaml -p jump-technical-case
COMPOSE_COMMAND := $(COMPOSE_BASE_COMMAND)

usage:
	@echo "\033[34mUsage:\033[0m make <target>"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m \r\t\t%s\n", $$1, $$2}'

up: ## Build, (re)create and start a container, by default, the container is exposed on http://localhost:8999/
	$(COMPOSE_COMMAND) --profile development up --force-recreate --remove-orphans --build

down: ## Stops containers and removes containers, networks, volumes, and images created by up
	$(COMPOSE_COMMAND) down --remove-orphans --rmi local --volumes

sh: ## Runs a new /bin/sh in the running container
	$(COMPOSE_COMMAND) --profile development exec -it local /bin/sh
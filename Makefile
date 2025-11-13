.PHONY: help build up down restart logs clean dev

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build all services
	docker-compose build

up: ## Start all services
	docker-compose up -d

dev: ## Start all services with logs
	docker-compose up

down: ## Stop all services
	docker-compose down

restart: ## Restart all services
	docker-compose restart

logs: ## Show logs for all services
	docker-compose logs -f

logs-backend: ## Show backend logs
	docker-compose logs -f backend

logs-frontend: ## Show frontend logs
	docker-compose logs -f frontend

logs-db: ## Show database logs
	docker-compose logs -f db

clean: ## Remove containers, networks, and volumes
	docker-compose down -v --remove-orphans
	docker system prune -f

rebuild: ## Rebuild and restart all services
	$(MAKE) down
	$(MAKE) build
	$(MAKE) up

backend-shell: ## Open shell in backend container
	docker-compose exec backend sh

frontend-shell: ## Open shell in frontend container
	docker-compose exec frontend sh

db-shell: ## Open MySQL shell
	docker-compose exec db mysql -u app_user -p app_db
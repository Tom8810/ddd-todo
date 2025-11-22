
include backend/.env
export

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

backend-generate:
	cd backend && gqlgen generate

# Migration commands
migrate-create: ## Create new migration file (usage: make migrate-create name=create_users)
	cd backend && migrate create -ext sql -dir db/migrations -seq $(name)

migrate-up: ## Run all up migrations
	cd backend && migrate -path db/migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" up

migrate-down: ## Run one down migration
	cd backend && migrate -path db/migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" down 1

migrate-force: ## Force migration version (usage: make migrate-force version=1)
	cd backend && migrate -path db/migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" force $(version)

# GORM Gen commands
db-generate: ## Generate GORM models from database
	cd backend && go run db/generator/main.go

backend-shell: ## Open shell in backend container
	docker-compose exec backend sh

frontend-shell: ## Open shell in frontend container
	docker-compose exec frontend sh

db-shell: ## Open MySQL shell
	docker-compose exec db mysql -u app_user -p app_db
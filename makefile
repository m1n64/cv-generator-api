# Variables
SERVICES = shared-network user-service cv-service gateway-service

.PHONY: help
help:
	@echo "Usage:"
	@echo "  make up - Start all services"
	@echo "  make stop - Stop all services"
	@echo "  make restart - Restart all services"
	@echo "  make network - Create docker network"
	@echo "  make up-service SERVICE= - Start service"
	@echo "  make stop-service SERVICE= - Stop service"
	@echo "  make restart-service SERVICE= - Restart service"

.PHONY: init
init:
	@echo "🛠️  Initialization all services and networks..."
	$(MAKE) network
	$(MAKE) up

# Инициализация всех сервисов
#.PHONY: init
#init:
#	@echo "🛠️  Initialization all services..."
#	@for service in $(SERVICES); do \
#		echo "🛠️  Initialization $$service..."; \
#		(cd $$service && docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d || echo "❌ Error initialization $$service"); \
#	done

# Запуск всех сервисов
.PHONY: up
up:
	@echo "🚀 Start all serviceS..."
	@for service in $(SERVICES); do \
    		echo "🟢  Starting $$service..."; \
    		(cd $$service && docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d || echo "❌ Error starting $$service"); \
    done

.PHONY: stop
stop:
	@echo "📴 Stop all services..."
	@for service in $(SERVICES); do \
		echo "🔴 Stopping $$service..."; \
		(cd $$service && docker-compose stop || echo "❌ Error stop $$service"); \
	done

.PHONY: restart
restart:
	@echo "🔄 Reload all services..."
	@for service in $(SERVICES); do \
		echo "🔄 Reload $$service..."; \
		(cd $$service && docker-compose -f docker-compose.yml -f docker-compose.dev.yml restart || echo "❌ Error restarting $$service"); \
	done

.PHONY: restart-service
restart-service:
	@if [ -z "$(SERVICE)" ]; then \
		echo "❌ Please specify a service to restart, e.g., 'make restart-service SERVICE=cv-service'"; \
		exit 1; \
	fi
	@echo "🔄 Restarting service: $(SERVICE)..."
	@(cd $(SERVICE) && docker-compose -f docker-compose.yml -f docker-compose.dev.yml restart || echo "❌ Error restarting $(SERVICE)")

.PHONY: up-service
up-service:
	@if [ -z "$(SERVICE)" ]; then \
		echo "❌ Please specify a service to start, e.g., 'make up-service SERVICE=cv-service'"; \
		exit 1; \
	fi
	@echo "🚀 Starting service: $(SERVICE)..."
	@(cd $(SERVICE) && docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d || echo "❌ Error starting $(SERVICE)")

.PHONY: stop-service
stop-service:
	@if [ -z "$(SERVICE)" ]; then \
		echo "❌ Please specify a service to stop, e.g., 'make stop-service SERVICE=cv-service'"; \
		exit 1; \
	fi
	@echo "📴 Stopping service: $(SERVICE)..."
	@(cd $(SERVICE) && docker-compose stop || echo "❌ Error stopping $(SERVICE)")

.PHONY: network
network:
	@echo "🌐 Create network..."
	docker network create cv-generator-network || echo "❌ Network \"cv-generator-network\" already exists"
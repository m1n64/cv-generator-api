# Variables
SERVICES = shared-network user-service cv-service information-service templates-service generator-service ai-service gateway-service healthcheck-service
SHARED_NETWORK_DELAY = 15

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
	$(MAKE) copy-envs
	$(MAKE) up

# Init all services
#.PHONY: init
#init:
#	@echo "🛠️  Initialization all services..."
#	@for service in $(SERVICES); do \
#		echo "🛠️  Initialization $$service..."; \
#		(cd $$service && docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d || echo "❌ Error initialization $$service"); \
#	done

# Start all services
.PHONY: up
up:
	@echo "🚀 Start all serviceS..."
	@for service in $(SERVICES); do \
    	if [ "$$service" = "shared-network" ]; then \
    		echo "🟢 Starting $$service with default compose file..."; \
    		(cd $$service && docker-compose up -d || echo "❌ Error starting $$service"); \
    		echo "⏳ Waiting for shared-network to be ready for $(SHARED_NETWORK_DELAY) seconds..."; \
            sleep $(SHARED_NETWORK_DELAY); \
    	else \
    		echo "🟢 Starting $$service with dev compose file..."; \
    		(cd $$service && docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d || echo "❌ Error starting $$service"); \
    	fi; \
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
  		if [ "$$service" = "shared-network" ]; then \
      		echo "🔄 Reload $$service with default compose file..."; \
      		(cd $$service && docker-compose restart || echo "❌ Error restarting $$service"); \
      		echo "⏳ Waiting for shared-network to be ready for $(SHARED_NETWORK_DELAY) seconds..."; \
            sleep $(SHARED_NETWORK_DELAY); \
      	else \
      		echo "🔄 Reload $$service with dev compose file..."; \
      		(cd $$service && docker-compose -f docker-compose.yml -f docker-compose.dev.yml restart || echo "❌ Error restarting $$service"); \
      	fi; \
	done

.PHONY: restart-service
restart-service:
	@if [ -z "$(SERVICE)" ]; then \
		echo "❌ Please specify a service to restart, e.g., 'make restart-service SERVICE=cv-service'"; \
		exit 1; \
	fi
	@echo "🔄 Restarting service: $(SERVICE)..."
	@if [ "$(SERVICE)" = "shared-network" ]; then \
    	(cd $(SERVICE) && docker-compose up -d || echo "❌ Error starting $(SERVICE)"); \
    else \
    	(cd $(SERVICE) && docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d || echo "❌ Error restarting $(SERVICE)"); \
    fi

.PHONY: up-service
up-service:
	@if [ -z "$(SERVICE)" ]; then \
		echo "❌ Please specify a service to start, e.g., 'make up-service SERVICE=cv-service'"; \
		exit 1; \
	fi
	@echo "🚀 Starting service: $(SERVICE)..."
	@if [ "$(SERVICE)" = "shared-network" ]; then \
    	(cd $(SERVICE) && docker-compose up -d || echo "❌ Error starting $(SERVICE)"); \
    else \
    	(cd $(SERVICE) && docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d || echo "❌ Error starting $(SERVICE)"); \
    fi

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

.PHONY: copy-envs
copy-envs:
	@echo "📄 Copy .env.example to .env, if file exists and .env not exists..."
	@for service in $(SERVICES); do \
		if [ -f $$service/.env.example ]; then \
			if [ ! -f $$service/.env ]; then \
				cp $$service/.env.example $$service/.env; \
				echo "✅ $$service: .env created"; \
			else \
				echo "ℹ️  $$service: .env already created, skipping"; \
			fi; \
		else \
			echo "⚠️  $$service: .env.example not found"; \
		fi; \
	done
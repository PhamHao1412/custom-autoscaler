# ===== Custom Autoscaler Makefile =====

APP_NAME = autoscaler
GO_FILES = ./cmd/main.go

# Build Go binary
build:
	@echo "🔨 Building $(APP_NAME)..."
	go mod tidy
	go build -o bin/$(APP_NAME) $(GO_FILES)

# Run autoscaler locally (non-docker)
run:
	@echo "🚀 Running autoscaler locally..."
	go run $(GO_FILES)

# Build Docker image for autoscaler
docker-build:
	@echo "🐳 Building Docker image for $(APP_NAME)..."
	docker build -t custom-autoscaler .

# Run full monitoring stack (autoscaler + prometheus + grafana)
up:
	@echo "📊 Starting Autoscaler + Prometheus + Grafana..."
	docker-compose up -d --build

# Stop all containers
down:
	@echo "🛑 Stopping all services..."
	docker-compose down

# View live logs from all containers
logs:
	@docker-compose logs -f

# Clean up everything (containers, volumes, binary)
clean:
	@echo "🧹 Cleaning up containers, volumes, and binaries..."
	docker-compose down -v
	rm -rf bin

.PHONY: build run docker-build up down logs clean

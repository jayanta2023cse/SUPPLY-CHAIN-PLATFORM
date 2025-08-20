#!/bin/bash

# Real-Time Supply Chain Visibility Platform
# Script to run all microservices for development

set -e

echo "🚚 Starting Supply Chain Platform Development Environment"
echo "========================================================="

# Check if Docker is running
# if ! docker info >/dev/null 2>&1; then
#     echo "❌ Docker is not running. Please start Docker first."
#     exit 1
# fi

# Create .env file if it doesn't exist
# if [ ! -f .env ]; then
#     echo "📝 Creating .env file from template..."
#     cp .env.example .env 2>/dev/null || cp .env .env.backup
# fi

# Start infrastructure services first
echo "🐘 Starting infrastructure services (PostgreSQL, Redis, Kafka)..."
docker-compose -f docker/docker-compose.yml up -d postgres redis zookeeper kafka

# Wait for services to be healthy
echo "⏳ Waiting for infrastructure services to be ready..."
sleep 30

# Check if Kafka is ready
echo "🔄 Checking Kafka readiness..."
docker-compose -f docker/docker-compose.yml exec kafka kafka-topics --bootstrap-server localhost:9092 --list >/dev/null 2>&1 || {
    echo "⚠️  Kafka not ready yet, waiting longer..."
    sleep 30
}

# Create Kafka topics
echo "📋 Creating Kafka topics..."
docker-compose -f docker/docker-compose.yml exec kafka kafka-topics \
    --bootstrap-server localhost:9092 \
    --create --if-not-exists \
    --topic shipment-events \
    --partitions 3 \
    --replication-factor 1

docker-compose -f docker/docker-compose.yml exec kafka kafka-topics \
    --bootstrap-server localhost:9092 \
    --create --if-not-exists \
    --topic alert-events \
    --partitions 3 \
    --replication-factor 1

docker-compose -f docker/docker-compose.yml exec kafka kafka-topics \
    --bootstrap-server localhost:9092 \
    --create --if-not-exists \
    --topic payment-events \
    --partitions 3 \
    --replication-factor 1

# Start monitoring services
echo "📊 Starting monitoring services..."
docker-compose -f docker/docker-compose.yml up -d prometheus grafana

# Build and start all microservices
echo "🏗️  Building and starting microservices..."
docker-compose -f docker/docker-compose.yml up --build -d user-service shipment-service alert-service payment-service analytics-service

# Wait for services to start
echo "⏳ Waiting for microservices to start..."
sleep 20

# Check service health
echo "🏥 Checking service health..."
services=("user-service:8001" "shipment-service:8002" "alert-service:8003" "payment-service:8004" "analytics-service:8005")

for service in "${services[@]}"; do
    IFS=':' read -r name port <<< "$service"
    if curl -f http://localhost:$port/health >/dev/null 2>&1; then
        echo "✅ $name is healthy"
    else
        echo "❌ $name is not responding"
    fi
done

echo ""
echo "🎉 Supply Chain Platform is now running!"
echo "========================================="
echo ""
echo "📚 API Documentation:"
echo "   User Service:      http://localhost:8001/swagger/index.html"
echo "   Shipment Service:  http://localhost:8002/swagger/index.html"
echo "   Alert Service:     http://localhost:8003/swagger/index.html"
echo "   Payment Service:   http://localhost:8004/swagger/index.html"
echo "   Analytics Service: http://localhost:8005/swagger/index.html"
echo ""
echo "🔍 Service Health Checks:"
echo "   User Service:      http://localhost:8001/health"
echo "   Shipment Service:  http://localhost:8002/health"
echo "   Alert Service:     http://localhost:8003/health"
echo "   Payment Service:   http://localhost:8004/health"
echo "   Analytics Service: http://localhost:8005/health"
echo ""
echo "📊 Monitoring:"
echo "   Prometheus:        http://localhost:9090"
echo "   Grafana:           http://localhost:3000 (admin/admin)"
echo ""
echo "💾 Infrastructure:"
echo "   PostgreSQL:        localhost:5432 (postgres/postgres)"
echo "   Redis:             localhost:6379"
echo "   Kafka:             localhost:9092"
echo ""
echo "🛑 To stop all services: docker-compose -f docker/docker-compose.yml down"
echo "🗑️  To reset data: docker-compose -f docker/docker-compose.yml down -v"
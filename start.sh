#!/bin/bash

# Sailing Backend API - Quick Start Script

set -e

echo "======================================"
echo "Sailing Backend API - Quick Start"
echo "======================================"
echo ""

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if .env exists
if [ ! -f .env ]; then
    echo -e "${YELLOW}Creating .env file from .env.example...${NC}"
    cp .env.example .env
    echo -e "${GREEN}.env file created successfully${NC}"
    echo ""
fi

# Check Docker
if ! command -v docker &> /dev/null; then
    echo -e "${RED}Docker is not installed. Please install Docker first.${NC}"
    exit 1
fi

# Check Docker Compose
if ! command -v docker-compose &> /dev/null; then
    echo -e "${RED}Docker Compose is not installed. Please install Docker Compose first.${NC}"
    exit 1
fi

echo "Starting services with Docker Compose..."
echo ""

# Start services
docker-compose up -d

echo ""
echo -e "${GREEN}Services started successfully!${NC}"
echo ""
echo "======================================"
echo "Service URLs:"
echo "======================================"
echo "API Server:    http://localhost:8080"
echo "Health Check:  http://localhost:8080/health"
echo "MongoDB:       mongodb://localhost:27017"
echo ""
echo "======================================"
echo "API Authentication:"
echo "======================================"
echo "API Key: sailing-api-key-12345"
echo "Header:  X-API-Key: sailing-api-key-12345"
echo ""
echo "======================================"
echo "Useful Commands:"
echo "======================================"
echo "View logs:        docker-compose logs -f"
echo "Stop services:    docker-compose down"
echo "Restart services: docker-compose restart"
echo ""
echo -e "${GREEN}Setup complete! You can now test the API.${NC}"
echo ""

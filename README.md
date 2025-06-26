# Thmanyah CMS

Golang-based content management system with gRPC/HTTP APIs for media and podcast management.

## Table of Contents

- [Features](#features)
- [Technology Stack](#technology-stack)
- [Architecture](#architecture)
- [Quick Start](#quick-start)
- [API Documentation](#api-documentation)
- [Project Structure](#project-structure)
- [Development](#development)
- [Deployment](#deployment)
- [Performance](#performance)
- [Appendix](#appendix)
  - [Full Journey Test](#full-journey-test)
  - [Future Improvements](#future-improvements)

## Features

- **File Upload**: S3-compatible storage for episode thumbnails and media files via MinIO
- **Content Management**: Full CRUD operations for programs, episodes, and categories
- **Search**: Full-text search across content with PostgreSQL
- **Authentication**: JWT-based user authentication and authorization

## Technology Stack

- **Language**: Go 1.24+
- **Database**: PostgreSQL 17+
- **Storage**: MinIO (S3-compatible)
- **Cache**: Ristretto in-memory
- **API**: gRPC + HTTP/JSON
- **Authentication**: JWT tokens
- **Documentation**: OpenAPI/Swagger
- **Testing**: Go testing
- **Infrastructure**: Docker + Docker Compose

## Architecture

Clean architecture with three distinct layers:

### Network Layer
- **Location**: `internal/modules/*/service/`
- **Purpose**: gRPC/HTTP handlers and protocol conversion
- **Components**: AuthService, CmsService, DiscoverService

### Business Layer
- **Location**: `internal/modules/*/biz/`
- **Purpose**: Business logic and use cases
- **Components**: UseCase implementations with domain rules

### Data Layer
- **Location**: `internal/modules/*/data/`
- **Purpose**: Database operations and external integrations
- **Components**: Repository implementations with database queries

### Dependency Injection
- **Tool**: Google Wire
- **Config**: `cmd/app/wire.go`
- **Generation**: `make generate`

## Quick Start

### Prerequisites

- Docker and Docker Compose
- Make utility
- Go 1.24+ (for development)

### Start Services

```bash
make start  # Starts API on :8000 (HTTP) and :8001 (gRPC)
```

### Stop Services

```bash
make stop   # Graceful shutdown
make down   # Remove volumes
```

## API Documentation

### Access Points
- **Swagger UI**: `http://localhost:8000/q/swagger-ui`
- **OpenAPI Spec**: `embeds/openapi.yaml`
- **Proto Definitions**: `api/v1/*.proto`

### Authentication
- **Method**: Bearer JWT tokens
- **Scope**: All CMS endpoints require authentication
- **Access Control**: Owner-based permissions for mutations

### API Versioning
- **Current Version**: v1
- **Base Path**: `/api/v1/`

### Endpoints
- **Auth**: `/auth/login`, `/auth/register`, `/auth/profile`
- **Categories**: CRUD operations for content categories
- **Programs**: CRUD operations for podcast/media programs
- **Episodes**: CRUD operations for individual episodes
- **Import**: Bulk data import functionality
- **Search**: Full-text search across content

## Project Structure

```
thmanyah/
├── api/                     # Protocol buffer definitions
│   ├── grpc/v1/             # Generated gRPC code
│   └── v1/                  # Proto source files
├── assets/                  # Static assets
├── cmd/app/                 # Application entry point
│   ├── main.go              # Main application
│   └── wire.go              # Dependency injection
├── configs/                 # Configuration files
├── embeds/                  # Embedded assets (OpenAPI specs)
├── internal/                # Private application code
│   ├── modules/             # Feature modules
│   │   ├── cms/             # Content management
│   │   │   ├── biz/         # Business logic
│   │   │   ├── data/        # Data repositories & S3
│   │   │   └── service/     # Network services
│   │   └── discover/        # Content discovery
│   │       ├── biz/         # Business logic
│   │       ├── data/        # Data repositories & cache
│   │       └── service/     # Network services
│   ├── conf/                # Configuration structs
│   ├── postgres/            # Database setup
│   ├── server/              # Server setup
│   └── utils/               # Shared utilities
├── keys/                    # Key management
├── platform/                # Infrastructure
│   ├── docker/              # Docker configurations
│   └── sql/                 # Database schemas
└── third_party/             # External dependencies
    ├── google/              # Google proto files
    ├── openapi/             # OpenAPI proto files
    ├── swaggerui/           # Swagger UI assets
    ├── validate/            # Validation proto files
    └── errors/              # Error proto files
```

## Development

### Prerequisites

- **Go 1.24+**: Download from [golang.org](https://golang.org/downloads/)
- **GOPATH/GOBIN**: Ensure Go binary directory is in your PATH
  ```bash
  export PATH=$PATH:$(go env GOPATH)/bin
  ```

### Setup Development Environment

```bash
# Install required development tools and binaries
make init
```

### Code Generation
```bash
make generate  # Generate Wire dependencies
make api      # Generate API code from protos
```

### Local Development with IDE

For development, you can run the application from your IDE while using Docker for dependencies:

```bash
# 1. Start required services (PostgreSQL + MinIO)
make start

# 2. Stop the API container (keep databases running)
docker compose -p thmanyah -f platform/docker/docker-compose.yaml stop thmanyah-api

# 3. Run from IDE or command line
go run cmd/app/main.go
```

#### Configuration

Configure the application using either:

**Environment Variables** (prefixed with `CONFIG_`):
```bash
export CONFIG_DATABASE_HOST=localhost
export CONFIG_DATABASE_PORT=5432
export CONFIG_S3_ENDPOINT=localhost:9000
```

**Configuration File**:
```bash
# Edit the configuration file directly
nano configs/config.yaml
```

### Testing

```bash
# Run all tests
go test ./... -v
```

### Build

```bash
make build        # Build binary to bin/
```

### Debugging

```bash
# View logs from all services
make debug-logs

# View API service logs only
make debug-api-logs

# View database logs
make debug-db-logs

# Check container status
make debug-ps

# Access API container shell
make debug-shell

# Access PostgreSQL shell
make debug-db-shell

# Access MinIO console (file storage)
make debug-minio

# Reset environment (clean restart)
make debug-reset
```

## Deployment

### Docker

```bash
make start # Build API Docker image and Start with docker-compose
```

Or 

```bash
make api-image    # Build API Docker image
make up          # Start with docker-compose
```

## Performance

### Benchmarks

#### Search Endpoint (`/api/v1/discover/search`)
- **Throughput**: 56,733 requests/second
- **Average Response Time**: 1.4ms
- **99th Percentile**: 8.7ms
- **Test**: 300,000 requests with 100 concurrent connections

#### Featured Endpoint (`/api/v1/discover/featured`)
- **Throughput**: 69,777 requests/second
- **Average Response Time**: 1.0ms
- **99th Percentile**: 7.0ms  
- **Test**: 300,000 requests with 100 concurrent connections

#### Load Testing Commands
```bash
make search-load-test    # Test search endpoint performance
make featured-load-test  # Test featured endpoint performance
```

**Tool**: `hey` CLI load testing  
**Database**: PostgreSQL with full-text search  
**Caching**: Ristretto in-memory cache (5min search, 15min featured)

#### Test Environment
- **CPU**: Intel i9-13950HX (32 cores) @ 5.3GHz
- **Memory**: 64GB RAM
- **OS**: Ubuntu 22.04.5 LTS
- **Storage**: Local development environment

> **Note**: Performance results may vary significantly based on hardware specifications, available system resources, network conditions, and concurrent system load. Your results may differ depending on CPU performance, memory capacity, and I/O capabilities.

### Optimizations

- **Search**: `ts_vector` with `pg_trgm` and `unaccent` extensions
- **Indexing**: GIN indexes on tags and search vectors
- **Connection Pooling**: pgx connection pool
- **Caching**: Memory cache for frequently accessed data

### Architecture Guidelines
- Keep business logic in `biz` layer
- Database operations only in `repo` layer
- Protocol handling only in `service` layer
- Use dependency injection via Wire

## Appendix

### Full Journey Test

Complete end-to-end workflow testing using `curl` and `jq` commands.

### Prerequisites

```bash
# Ensure the API is running
make start

# Set base URL
export API_BASE="http://localhost:8000/api/v1"
```

### 1. Register User

```bash
curl -X POST "$API_BASE/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "securepassword123",
    "confirm_password": "securepassword123",
    "name": "Test User"
  }'
```

### 2. Login and Store Access Token

```bash
# Login and extract access token
RESPONSE=$(curl -s -X POST "$API_BASE/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "securepassword123"
  }')

# Store access token in environment variable
export ACCESS_TOKEN=$(echo $RESPONSE | jq -r '.access_token')
```

### 3. Create Category

```bash
CATEGORY_RESPONSE=$(curl -s -X POST "$API_BASE/cms/categories" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -d '{
    "name": "Technology Podcast",
    "description": "Latest trends in technology and innovation",
    "type": "CATEGORY_TYPE_PODCAST"
  }')

# Store category ID
export CATEGORY_ID=$(echo $CATEGORY_RESPONSE | jq -r '.category.id')
```

### 4. Create Program

```bash
PROGRAM_RESPONSE=$(curl -s -X POST "$API_BASE/cms/programs" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -d '{
    "title": "Tech Talk Weekly",
    "description": "Weekly discussions on the latest technology trends",
    "category_id": "'$CATEGORY_ID'",
    "tags": ["weekly", "technology", "discussion"],
    "is_featured": true
  }')

# Store program ID
export PROGRAM_ID=$(echo $PROGRAM_RESPONSE | jq -r '.program.id')
```

### 5. Create Episode

```bash
EPISODE_RESPONSE=$(curl -s -X POST "$API_BASE/cms/episodes" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -d '{
    "program_id": "'$PROGRAM_ID'",
    "title": "AI Revolution in 2024",
    "description": "Exploring the latest AI developments and their impact",
    "duration_seconds": 3600,
    "episode_number": 1,
    "season_number": 1,
    "tags": ["AI", "2024", "revolution"]
  }')

# Store episode ID
export EPISODE_ID=$(echo $EPISODE_RESPONSE | jq -r '.episode.id')
```

### 6. Search Content

```bash
curl -X POST "$API_BASE/discover/search" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -d '{
    "query": "AI",
    "page": 1,
    "page_size": 10
  }'
```

### 7. Get Featured Programs

```bash
curl -X POST "$API_BASE/discover/featured" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -d '{}'
```

### Future Improvements

Scalability and performance enhancements for high-load production environments.

#### Search Engine Integration
- **Elasticsearch/Solr**: Replace PostgreSQL full-text search with dedicated search engines for high-intensive search loads
- **Data Synchronization**: Implement PostgreSQL logical replication to sync data between CMS and search engine in real-time
- **Search Features**: Enable advanced search capabilities like faceted search, auto-complete, and search analytics

#### Distributed Architecture
- **Kubernetes Deployment**: Convert Docker Compose setup to Kubernetes deployment with:
  - Multiple API replicas for horizontal scaling
  - Load balancers for traffic distribution
  - Auto-scaling based on CPU/memory usage
  - Rolling updates with zero downtime

#### Centralized Caching
- **Redis Integration**: Replace in-memory Ristretto cache with Redis for:
  - Shared cache across multiple application instances
  - Cache persistence and high availability
  - Advanced caching strategies (LRU, TTL, cache warming)
  - Session storage for distributed authentication

#### Additional Enhancements
- **CDN Integration**: CloudFront/CloudFlare for media file delivery
- **Message Queues**: Redis/RabbitMQ for asynchronous processing
- **Monitoring**: Prometheus + Grafana for metrics and alerting
- **Distributed Tracing**: Jaeger/Zipkin for request tracking

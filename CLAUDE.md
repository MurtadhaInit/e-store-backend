# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Database Code Generation
- `sqlc generate` - Generate Go code from SQL queries in db/queries/
- `sqlc diff` - Check if generated code is up to date with schema/queries
- `sqlc vet` - Validate SQL queries

### Development Server
- `docker compose up` - Start development environment with hot reload (uses Air)
- `docker compose -f compose.yml -f compose.prod.yml up` - Start production environment
- Server runs on port 4210

### Code Quality
- `golangci-lint run` - Run Go linter (configured in GitHub Actions)
- Tests are in the `tests/` directory

## Architecture Overview

This is a Go e-commerce backend using:
- **Database**: MySQL 8.4 with SQLC for type-safe queries
- **HTTP Router**: Go 1.24 stdlib HTTP mux with pattern matching
- **Hot Reload**: Air for development (configured in .air.toml)
- **Containerization**: Multi-stage Docker build (dev/prod targets)

### Project Structure
- `cmd/web/` - HTTP server, handlers, and routes
- `internal/repository/` - SQLC-generated database models and queries
- `db/schema.sql` - Database schema definition
- `db/queries/` - SQL queries for SQLC generation

### Database Schema
Core entities: customers, products, product_categories, orders, order_items, carts, cart_items

### Application Structure
- `application` struct contains logger and database queries
- Handlers follow REST conventions with proper HTTP methods
- Uses structured logging with slog
- Database connection configured via DSN environment variable

### Key Configuration
- DSN environment variable required for database connection
- Server address configurable via -addr flag (default :4210)
- Air configuration builds from ./cmd/** directory
- SQLC generates code in internal/repository package with JSON tags

### Development Flow
1. Modify SQL schema in db/schema.sql
2. Add/modify queries in db/queries/
3. Run `sqlc generate` to update Go code
4. Implement handlers in cmd/web/
5. Use `docker compose up` for hot reload development
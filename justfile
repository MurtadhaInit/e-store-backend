# List available recipes
default:
    just --list

# Generate Go function from SQL queries
sqlc-gen:
    sqlc generate

# Show schema migration status
goose-status:
    goose status

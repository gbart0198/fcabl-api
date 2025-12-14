#!/bin/bash
# init-db.sh - Database initialization helper script

set -e

DB_CONTAINER="fcabl-db"
DB_USER="fcabl"
DB_NAME="fcabl_db"

echo "Waiting for database to be ready..."
until docker exec $DB_CONTAINER pg_isready -U $DB_USER -d $DB_NAME > /dev/null 2>&1; do
  echo "Database is unavailable - sleeping"
  sleep 2
done

echo "Database is ready!"
echo "Data is persisted in Docker volume 'fcabl_pgdata'"
echo ""
echo "Useful commands:"
echo "  Connect to DB:    docker exec -it $DB_CONTAINER psql -U $DB_USER -d $DB_NAME"
echo "  View tables:      docker exec -it $DB_CONTAINER psql -U $DB_USER -d $DB_NAME -c '\dt'"
echo "  Reset database:   docker-compose down -v && docker-compose up -d"
echo "  Keep data:        docker-compose down && docker-compose up -d"

#!/usr/bin/env bash
set -euo pipefail

APP_SLUG="__APP_SLUG__"
DEPLOY_DIR="/opt/${APP_SLUG}"
BACKUP_DIR="${DEPLOY_DIR}/backups"

mkdir -p "$BACKUP_DIR"
cd "$DEPLOY_DIR"

set -a
source .env
set +a

TIMESTAMP="$(date +%Y-%m-%d_%H-%M-%S)"
OUT_FILE="${BACKUP_DIR}/${APP_SLUG}_${TIMESTAMP}.sql.gz"

docker compose -f docker-compose.prod.yaml exec -T postgres \
  pg_dump -U "$POSTGRESQL_USERNAME" -d "$POSTGRESQL_DATABASE" \
  | gzip > "$OUT_FILE"

echo "backup written to $OUT_FILE"

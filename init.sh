#!/usr/bin/env bash
set -euo pipefail

MODULE=""
APP_SLUG=""
IMAGE_REPOSITORY=""

# Parse args
for arg in "$@"; do
  case $arg in
    --module=*)
      MODULE="${arg#*=}"
      shift
      ;;
    --app-slug=*)
      APP_SLUG="${arg#*=}"
      shift
      ;;
    --image-repo=*)
      IMAGE_REPOSITORY="${arg#*=}"
      shift
      ;;
    *)
      echo "Unknown argument: $arg"
      exit 1
      ;;
  esac
done

# Validate
if [ -z "$MODULE" ] || [ -z "$APP_SLUG" ] || [ -z "$IMAGE_REPOSITORY" ]; then
  echo "Usage:"
  echo "  ./init.sh --module=... --app-slug=... --image-repo=..."
  echo ""
  echo "Example:"
  echo "  ./init.sh \\"
  echo "    --module=github.com/alex/myapp \\"
  echo "    --app-slug=myapp \\"
  echo "    --image-repo=ghcr.io/alex/myapp"
  exit 1
fi

echo "Initializing project..."
echo "MODULE=$MODULE"
echo "APP_SLUG=$APP_SLUG"
echo "IMAGE_REPOSITORY=$IMAGE_REPOSITORY"

# rename cmd directory
if [ -d "cmd/__APP_SLUG__" ]; then
	mv cmd/__APP_SLUG__ "cmd/$APP_SLUG"
fi

# Replace placeholders
if [[ "$OSTYPE" == "darwin"* ]]; then
  find . -type f \
    \( \
      -name "*.go" -o \
      -name "go.mod" -o \
      -name "*.yaml" -o \
      -name "*.yml" -o \
      -name "*.json" -o \
      -name "*.sh" -o \
      -name "Dockerfile*" -o \
      -name "Makefile" -o \
      -name "Caddyfile" -o \
      -name "*.md" -o \
      -name "*.sql" \
    \) \
    -exec sed -i '' \
      "s|__MODULE__|$MODULE|g; \
       s|__APP_SLUG__|$APP_SLUG|g; \
       s|__IMAGE_REPOSITORY__|$IMAGE_REPOSITORY|g" \
    {} +
else
  find . -type f \
    \( \
      -name "*.go" -o \
      -name "go.mod" -o \
      -name "*.yaml" -o \
      -name "*.yml" -o \
      -name "*.json" -o \
      -name "*.sh" -o \
      -name "Dockerfile*" -o \
      -name "Makefile" -o \
      -name "Caddyfile" -o \
      -name "*.md" -o \
      -name "*.sql" \
    \) \
    -exec sed -i \
      "s|__MODULE__|$MODULE|g; \
       s|__APP_SLUG__|$APP_SLUG|g; \
       s|__IMAGE_REPOSITORY__|$IMAGE_REPOSITORY|g" \
    {} +
fi

# Go cleanup
go mod tidy

if [ -d "frontend" ]; then
  (cd frontend && npm ci)
fi

echo "✅ Project initialized successfully."

# Delete this script after successful init
echo "Removing init.sh..."
SCRIPT_PATH="$(realpath "$0")"
rm -- "$SCRIPT_PATH"

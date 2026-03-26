# syntax=docker/dockerfile:1.7

FROM --platform=$BUILDPLATFORM node:24-alpine AS frontend-builder

WORKDIR /app

COPY frontend/package*.json ./frontend/
RUN cd frontend && npm ci

COPY frontend ./frontend
COPY internal/http/templates ./internal/http/templates

RUN cd frontend && npm run build


FROM --platform=$BUILDPLATFORM golang:1.25-alpine AS backend-builder

WORKDIR /app

RUN apk add --no-cache git brotli gzip jq coreutils
RUN go install github.com/a-h/templ/cmd/templ@v0.3.1001

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN templ generate
COPY --from=frontend-builder /app/frontend/dist ./internal/http/static

RUN apk add --no-cache brotli gzip jq coreutils

COPY deploy/scripts/build_static.sh /usr/local/bin/build-static
RUN chmod +x /usr/local/bin/build-static \
 && /usr/local/bin/build-static /app/internal/http/static \
 && chmod -R a+rX /app/internal/http/static

ARG VERSION=dev
ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLED=0 \
  GOOS=$TARGETOS \
  GOARCH=$TARGETARCH \
  go build \
  -ldflags "-X main.Version=${VERSION}" \
  -o /app/__APP_SLUG__ \
  ./cmd/__APP_SLUG__


FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --chown=65532:65532 --from=backend-builder /app/__APP_SLUG__ /app/__APP_SLUG__
COPY --chown=65532:65532 --from=backend-builder /app/internal/http/static /app/internal/http/static
COPY --from=backend-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER nonroot:nonroot

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=10s \
  CMD ["/app/__APP_SLUG__", "healthcheck"]

ENTRYPOINT ["/app/__APP_SLUG__"]

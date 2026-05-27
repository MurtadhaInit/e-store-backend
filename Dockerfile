# syntax=docker/dockerfile:1

ARG GO_VERSION=1.26

# ---- base: shared module-download layer (always native arch for speed)
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-bookworm AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# ---- dev: hot-reload via Air, used by docker compose
FROM base AS dev
RUN groupadd --system app \
      && useradd --system --gid app --home-dir /home/app --create-home app \
      && chown -R app:app /go /app
USER app
RUN go install github.com/air-verse/air@latest
CMD ["air", "-c", ".air.toml"]

# ---- build: cross-compile via Go for the target arch (no QEMU on the Go build)
FROM base AS build
ARG TARGETOS
ARG TARGETARCH
COPY . .
ENV CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH}
RUN go build \
      -trimpath \
      -buildvcs=false \
      -ldflags="-s -w" \
      -o /out/e-store-backend \
      .

# ---- prod: distroless, non-root, ~10 MB
FROM gcr.io/distroless/static-debian12:nonroot AS prod
ARG VERSION=dev
ARG REVISION=unknown
LABEL org.opencontainers.image.source="https://github.com/MurtadhaInit/e-store-backend" \
      org.opencontainers.image.title="e-store-backend" \
      org.opencontainers.image.description="Go API for the e-store project" \
      org.opencontainers.image.version="${VERSION}" \
      org.opencontainers.image.revision="${REVISION}"
WORKDIR /app
COPY --from=build /out/e-store-backend /app/e-store-backend
EXPOSE 4210
USER nonroot:nonroot
ENTRYPOINT ["/app/e-store-backend"]

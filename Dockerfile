FROM golang:1.24-bookworm AS base
RUN groupadd -r app && useradd --no-log-init -r -g app -m app
USER app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM base AS dev
RUN go install github.com/air-verse/air@latest
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 CMD [ "curl", "-f", "http://localhost:4210/health" ]
CMD ["air", "-c", ".air.toml"]

FROM base AS build
COPY . .
# To produce static binaries
ENV CGO_ENABLED=0
RUN go build -buildvcs=false -o ./bin/ .

FROM scratch AS prod
WORKDIR /app
COPY --from=build /app/bin/ ./
CMD [ "./e-store-backend" ]
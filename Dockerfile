FROM golang:1.24-bookworm AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM base AS dev
RUN go install github.com/air-verse/air@latest
CMD ["air", "-c", ".air.toml"]

FROM base AS build
COPY . .
# To produce static binaries
ENV CGO_ENABLED=0
RUN go build -o ./bin/ .

FROM scratch AS prod
WORKDIR /app
COPY --from=build /app/bin/ ./
CMD [ "./e-store-backend" ]
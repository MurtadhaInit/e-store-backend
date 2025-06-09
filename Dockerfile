FROM golang:1.24-bookworm AS build
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# To produce static binaries
ENV CGO_ENABLED=0
# TODO: why can't `go build` recursively search directories for Go files?
RUN go build -o ./bin ./cmd/**

FROM scratch
WORKDIR /app
COPY --from=build /build/bin/ ./
CMD [ "./bin" ]
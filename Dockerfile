FROM golang:1.24-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o bin/ .

FROM scratch
WORKDIR /app
COPY --from=build /app/bin/ ./
EXPOSE 8080
CMD [ "./bin/main" ]
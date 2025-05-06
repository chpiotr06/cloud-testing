FROM golang:1.24.2-alpine AS build_base

WORKDIR /tmp/ram-load-app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/ram-load-app .

FROM alpine:3.21.3

COPY --from=build_base ./tmp/ram-load-app/out/ram-load-app /app/ram-load-app
COPY --from=build_base ./tmp/ram-load-app/config.json /app/config.json

EXPOSE 3000

CMD ["/app/ram-load-app", "-config=/app/config.json"]

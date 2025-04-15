FROM golang:1.24.2-alpine AS build_base

WORKDIR /tmp/cpu-load-app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/cpu-load-app .

FROM alpine:3.21.3

COPY --from=build_base ./tmp/cpu-load-app/out/cpu-load-app /app/cpu-load-app 

EXPOSE 3000

CMD ["/app/cpu-load-app"]
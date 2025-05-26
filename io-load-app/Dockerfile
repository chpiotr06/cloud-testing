FROM golang:1.24.2-alpine AS build_base

WORKDIR /tmp/io-load-app

COPY . .

RUN go build -o ./out/io-load-app main.go

FROM alpine:3.21.3

COPY --from=build_base ./tmp/io-load-app/out/io-load-app /app/io-load-app
COPY --from=build_base ./tmp/io-load-app/1mb.txt /app/1mb.txt
COPY --from=build_base ./tmp/io-load-app/5mb.txt /app/5mb.txt
COPY --from=build_base ./tmp/io-load-app/10mb.txt /app/10mb.txt
COPY --from=build_base ./tmp/io-load-app/25mb.txt /app/25mb.txt
COPY --from=build_base ./tmp/io-load-app/50mb.txt /app/50mb.txt

EXPOSE 3000

WORKDIR /app

CMD ["./io-load-app"]

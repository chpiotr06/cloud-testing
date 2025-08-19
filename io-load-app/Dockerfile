FROM golang:1.24.2-alpine AS build_base

WORKDIR /tmp/io-load-app

COPY . .

RUN go build -o ./out/io-load-app main.go

FROM alpine:3.21.3

COPY --from=build_base ./tmp/io-load-app/out/io-load-app /app/io-load-app
COPY --from=build_base ./tmp/io-load-app/5kb.txt /app/5kb.txt
COPY --from=build_base ./tmp/io-load-app/10kb.txt /app/10kb.txt
COPY --from=build_base ./tmp/io-load-app/25kb.txt /app/25kb.txt
COPY --from=build_base ./tmp/io-load-app/50kb.txt /app/50kb.txt
COPY --from=build_base ./tmp/io-load-app/100kb.txt /app/100kb.txt

EXPOSE 3000

WORKDIR /app

CMD ["./io-load-app"]

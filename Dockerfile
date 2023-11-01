# build flutter-web stage
FROM neosu/flutter-web:3.13.9 AS flutter-web-builder

COPY ui /ui
WORKDIR /ui

RUN flutter pub get
RUN flutter build web --dart-define=BUILD_TYPE=prod --build-name=0.0.3 --tree-shake-icons --pwa-strategy none --web-renderer html --release

# build golang stage
FROM golang:1.21.3 AS go-builder

WORKDIR /app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
COPY --from=flutter-web-builder /ui/build/web /app/web
RUN go build -a -ldflags="-w -s" -o /app/cmd/kubebadges/main ./cmd/kubebadges

# production stage
FROM alpine:latest
WORKDIR /app
COPY --from=go-builder /app/cmd/kubebadges/main /app/main
RUN chmod +x /app/main
CMD ["/app/main"]
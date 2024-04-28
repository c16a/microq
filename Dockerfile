FROM docker.io/library/golang:1.22.2 AS builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o microq github.com/c16a/microq

FROM scratch

LABEL org.opencontainers.image.source=https://github.com/c16a/microq
LABEL org.opencontainers.image.description="A tiny event broker"
LABEL org.opencontainers.image.licenses=MIT

WORKDIR /app
COPY --from=builder /app/microq ./
CMD ["/app/microq"]

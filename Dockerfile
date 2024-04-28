FROM docker.io/library/golang:1.22.2 AS builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o microq github.com/c16a/microq

FROM scratch

WORKDIR /app
COPY --from=builder /app/microq ./
CMD ["/app/microq"]

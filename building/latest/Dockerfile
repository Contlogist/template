# Step 1: Modules caching
FROM golang:latest as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:latest as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -tags migrate -o /bin/app ./cmd/app

# Step 3: Final
FROM debian:buster
COPY --from=builder /app/config /config
COPY --from=builder /bin/app /app
COPY --from=builder /app/pkg/cert/ /etc/ssl/certs/
COPY --from=builder /app/pkg/cert/ /usr/share/ca-certificates

RUN apt-get update
RUN apt-get install ca-certificates -y
RUN update-ca-certificates

EXPOSE 1100

CMD ["/app"]
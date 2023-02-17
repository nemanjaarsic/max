FROM golang:alpine

WORKDIR /web-svc

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /web-svc ./...

EXPOSE 8000

ENV API_PORT="8000"
ENV DATABASE_SERVICE_HOST="db-svc:9000"
ENV IDEMPOTENCY_SERVICE_HOST="idempotency-svc:7000"

ENTRYPOINT ["/web-svc/max-web-svc"]
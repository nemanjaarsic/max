FROM golang:alpine

WORKDIR /idempotency-svc

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /idempotency-svc ./...

EXPOSE 7000

ENV HOST="0.0.0.0:7000"
ENV REDIS_HOST="redis_db:6379"
ENV REDIS_PASSWORD=""

ENTRYPOINT ["/idempotency-svc/max-idempotency-svc"]
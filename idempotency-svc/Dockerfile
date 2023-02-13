FROM golang:alpine

WORKDIR /idempotency-svc

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /idempotency-svc ./...

EXPOSE 7000

ENTRYPOINT ["/idempotency-svc/max-idempotency-svc"]
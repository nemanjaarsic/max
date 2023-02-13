FROM golang:alpine

WORKDIR /db-svc

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /db-svc ./...

EXPOSE 9000

ENTRYPOINT ["/db-svc/max-db-svc"]
FROM golang:alpine

WORKDIR /web-svc

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /web-svc ./...

EXPOSE 8000

ENTRYPOINT ["/web-svc/max-web-svc"]
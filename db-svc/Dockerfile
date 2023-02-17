FROM golang:alpine

WORKDIR /db-svc

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /db-svc ./...

EXPOSE 9000

ENV HOST="0.0.0.0:9000"
ENV POSTGRES_CONNECTION_STRING="postgres://postgres:maximilian@postgres_image:5432/maxDB?sslmode=disable"


ENTRYPOINT ["/db-svc/max-db-svc"]
# Build stage
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build server
RUN go build -o server ./cmd/server

# Build client
RUN go build -o client ./cmd/client

# Final stage
FROM swaggerapi/swagger-ui AS swagger

COPY ./protos/rusprofile_lookup.swagger.json /usr/share/nginx/html/swagger.json

FROM golang:latest

WORKDIR /app

COPY --from=builder /app/server /app/client ./

EXPOSE 8080

CMD ["./server"]

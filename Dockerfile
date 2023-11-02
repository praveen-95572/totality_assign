
FROM golang:latest
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o userservice ./cmd/server/

# Expose the port
EXPOSE 8088

# Command to run the application
CMD ["./userservice"]

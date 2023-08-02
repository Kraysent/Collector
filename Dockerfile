FROM golang:1.20

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./collector
RUN chmod +x ./collector
CMD ["/app/collector"]
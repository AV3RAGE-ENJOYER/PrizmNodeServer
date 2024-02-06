# syntax=docker/dockerfile:1

FROM golang:1.21.4
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN export CGO_ENABLED=1
RUN export GOOS=linux 
RUN go build -o /app
EXPOSE 8080
CMD ["/app"]

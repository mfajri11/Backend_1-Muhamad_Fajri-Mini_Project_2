FROM golang:1.20-alpine as pkg
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM golang:1.20-alpine as builder
COPY --from=pkg /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o mp ./main.go

FROM alpine:latest
COPY --from=builder /app/mp /app/
WORKDIR /app
CMD ["./mp"]
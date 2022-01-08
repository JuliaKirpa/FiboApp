FROM golang:latest as builder
WORKDIR /app
COPY ./ /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./testfbs ./cmd/main.go

FROM scratch
COPY --from=builder /app/testfbs /usr/bin/testfbs


ENTRYPOINT ["/usr/bin/testfbs"]
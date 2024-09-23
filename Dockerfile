FROM golang:1.18-alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
COPY .env ./

RUN go mod download
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o dev-user .

FROM scratch
COPY --from=builder ["/build/app", "/"]
EXPOSE 8888
ENTRYPOINT ["/app"]
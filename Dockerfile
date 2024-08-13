FROM golang:1.22.5-alpine3.20 AS build

WORKDIR /go/src/api-gateway

COPY . .

RUN go mod download

ARG CGO_ENABLED=0 GOOS=linux

RUN go build -o /out/api-gateway cmd/main.go

# FROM alpine:3.20
FROM scratch

WORKDIR /app

COPY --from=build /out/api-gateway ./

EXPOSE 8080

ENTRYPOINT [ "/app/api-gateway" ]
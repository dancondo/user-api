FROM golang:1.20.2-alpine3.17 AS builder

RUN apk add --no-cache pkgconfig make gcc musl-dev tzdata git && \
        cp /usr/share/zoneinfo/America/Recife /etc/localtime && \
        echo "America/Recife" >  /etc/timezone

WORKDIR /users-api

COPY . .

RUN go mod download && go mod verify

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g api/api.go -o /docs/swagger

RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --ldflags='-w -s -extldflags "-static"' -v -a -o /go/bin/users-api .

FROM alpine:3.17

RUN apk add --no-cache ca-certificates

COPY --from=builder /etc/localtime /etc/localtime
COPY --from=builder /etc/timezone /etc/timezone
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /go/bin/users-api /usr/bin/users-api
COPY --from=builder /docs/swagger/swagger.json /docs/swagger/swagger.json

EXPOSE 8080

ENTRYPOINT ["/usr/bin/users-api"]

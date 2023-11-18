# Builder
FROM golang:1.21.1-alpine as development
WORKDIR /go/src/app
RUN apk update && apk add --no-cache bash tzdata
COPY go.* ./
RUN go mod download -x
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -o main cmd/main.go
CMD ["go", "run", "cmd/main.go"]

# Distribution
FROM alpine:latest as production
WORKDIR /app
RUN apk update && apk add --no-cache bash
COPY --from=development /go/src/app/main .
COPY --from=development /usr/share/zoneinfo /usr/share/zoneinfo
CMD ["./main"]
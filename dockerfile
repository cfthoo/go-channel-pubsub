FROM golang:1.17-alpine AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o pubsub

FROM alpine:latest
COPY --from=build /app/pubsub .

ENTRYPOINT ["./pubsub"]
CMD ["run"]

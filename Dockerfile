FROM golang:1.17-alpine as build-img

WORKDIR /go/src/app
RUN apk update && apk add --update gcc musl-dev && rm -rf /var/cache/apk/*

COPY ./ .

RUN go mod tidy
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

RUN go test ./...

RUN go build -o tagging

# final stage
FROM alpine
EXPOSE 8080
WORKDIR /app
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=build-img /go/src/app/tagging /app/
ENTRYPOINT ["/app/tagging"]
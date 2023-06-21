FROM golang:1.20-alpine AS build-env
ENV GO111MODULE on

WORKDIR $GOPATH/src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine

WORKDIR /app
RUN apk add tzdata
COPY --from=build-env  go/src/main /app/
#APM Test End

EXPOSE 80

ENTRYPOINT ./main
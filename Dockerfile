FROM golang:1.17-alpine

WORKDIR /opt/code
ADD ./ /opt/code/

RUN apk update && apk upgrade && apk add --no-cache git

RUN go mod download

RUN go build -o ./build/builded_app ./app/src/main.go

ENTRYPOINT ["./build/builded_app"]
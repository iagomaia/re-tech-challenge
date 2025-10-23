FROM golang:1.24-alpine as build
RUN apk add --no-cache make

WORKDIR $GOPATH/re-tech-challenge
COPY . .

RUN mkdir -p build/static
RUN go build -mod=mod -v -o ./build ./...
RUN cp ./docs/swagger.yaml ./build/static/swagger.yaml
RUN cp -r ./static/swagger-ui ./build/static

FROM alpine:latest

RUN apk update
RUN apk upgrade
RUN apk add --no-cache bash
RUN apk add --upgrade --no-cache coreutils

WORKDIR /

COPY --from=build /go/re-tech-challenge/build /usr/local/bin

ARG PORT=3000
ENV SV_PORT=${PORT}

EXPOSE ${SV_PORT}

CMD ["api"]
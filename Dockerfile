FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/../../jenkins-x/test/myliles
COPY . /usr/local/go/src/../../jenkins-x/test/myliles
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/myliles ./myliles


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/../../jenkins-x/test/myliles/build/myliles /bin/myliles
CMD ["myliles", "up"]

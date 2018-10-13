FROM golang:1.11.1 as builder

WORKDIR ${GOPATH}/src/github.com/sivaramalingamk/wim-api
COPY . .
RUN set -x && \
    go get -d -v . && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o wim-api .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder go/src/github.com/sivaramalingamk/wim-api .
CMD ["./wim-api"]


# Expose the application on port 8080.
# This should be the same as in the app.conf file
EXPOSE 8080

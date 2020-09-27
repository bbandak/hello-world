FROM golang:alpine as builder
ADD . /usr/local/go/src/server
WORKDIR /usr/local/go/src/server
RUN go get server .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /usr/local/go/src/server /app/
WORKDIR /app
EXPOSE 11130
ENTRYPOINT ["./server"]

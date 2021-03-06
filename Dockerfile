FROM golang:alpine as builder
RUN mkdir /usr/local/go/src/server
ADD . /usr/local/go/src/server/
WORKDIR /usr/local/go/src/server
RUN go get server .
RUN go build -o server .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /usr/local/go/src/server /app/
WORKDIR /app
EXPOSE 11130
ENTRYPOINT ["./server"]




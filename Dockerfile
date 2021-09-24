FROM golang:1.17.1-alpine3.14 as builder
ADD . /app
WORKDIR /app
## Add this go mod download command to pull in any dependencies
RUN go mod download
## Our project will now successfully build with the necessary go libraries included.
RUN go build -o main .
## Our start command which kicks off
## our newly created binary executable
FROM alpine
WORKDIR /app
COPY --from=builder /app/ /app/
CMD ["/app/main"]

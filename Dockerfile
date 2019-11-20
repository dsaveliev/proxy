FROM golang:1.13.4-alpine AS builder
WORKDIR /go/src/github.com/hellofreshdevtests/dsaveliev-golang-test
ADD . .
# Don't forget to clean up the binary from the unnecessary debug data
RUN go build -v -ldflags '-w -s' -o app

# This is not very useful for the test, 
# but usually, I keep build image and runtime image separately
FROM alpine:latest AS web
WORKDIR /root/
COPY --from=builder /go/src/github.com/hellofreshdevtests/dsaveliev-golang-test .
CMD ["./app"]
EXPOSE 8080

# build stage
FROM golang:latest as build
WORKDIR /go/src/basegowebserver
COPY . .
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
RUN go install -v ./...

# run stage
FROM busybox
WORKDIR /go/bin/
COPY --from=build /go/bin/basegowebserver /go/bin/basegowebserver 
EXPOSE 11789
CMD ["/go/bin/basegowebserver"]

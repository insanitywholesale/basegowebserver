# build stage
FROM golang:1.17 as build

ENV CGO_ENABLED 0
ENV GO111MODULE on

WORKDIR /go/src/basegowebserver
COPY . .

RUN go install -v

# run stage
FROM busybox
COPY --from=build /go/bin/basegowebserver /go/bin/basegowebserver 
EXPOSE 11789
CMD ["/go/bin/basegowebserver"]

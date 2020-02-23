FROM golang:1.13 as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/prbuilder ./cmd/prbuilder/

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/prbuilder /
CMD ["/prbuilder"]
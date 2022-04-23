FROM golang:alpine as build
WORKDIR /go/src/app
ADD . /go/src/app
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -o /go/bin/app cmd/main.go

FROM gcr.io/distroless/static
COPY --from=build /go/bin/app /
EXPOSE 8080
CMD ["/app"]
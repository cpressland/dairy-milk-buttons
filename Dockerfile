FROM golang:latest AS build

ADD main.go .
RUN go build .

FROM scratch
COPY --from=build /go/go /
ENTRYPOINT [ "/go" ]

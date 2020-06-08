FROM golang as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN go build -o /go/bin/httpExtractor

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/httpExtractor /
EXPOSE 8080
CMD ["/httpExtractor"]
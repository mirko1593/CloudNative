FROM golang:1.17 as service_builder
ENV CGO_ENABLED 0

COPY . /service
WORKDIR /service
RUN go build 

FROM alpine:3.14
COPY --from=service_builder /service/. /service
WORKDIR /service
CMD ["./service"]


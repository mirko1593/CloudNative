FROM golang:1.17 as service_builder
ENV CGO_ENABLED 0

COPY . /service
WORKDIR /service
RUN go build 

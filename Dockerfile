FROM golang:alpine as builder
WORKDIR /src
COPY . /src/
RUN cd cmd \
    && rm main \
    && go build main.go

FROM alpine:latest as exe
WORKDIR /target
COPY --from=builder /src/cmd/main /target/
CMD ./main

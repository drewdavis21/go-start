FROM golang:1.16 as build
ENV GO111MODULE=on
COPY . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/server main.go

FROM alpine:latest as final
WORKDIR /app
COPY --from=build /build/bin .
ENTRYPOINT ["./server"]

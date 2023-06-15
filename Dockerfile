FROM golang:1.20.3-alpine AS builder

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY . /go/src/exporter-demo

WORKDIR /go/src/exporter-demo

RUN go build -o exporter-demo .

FROM alpine

COPY --from=builder /go/src/exporter-demo/exporter-demo /bin/exporter-demo

EXPOSE 18080

ENTRYPOINT [ "/bin/exporter-demo" ]
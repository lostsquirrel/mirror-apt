FROM registry.lisong.pub:5000/golang:1.13-buster AS builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
ENV CGO_ENABLED=0
ADD . /dist
WORKDIR /dist
RUN go get -v all
RUN go build \
        -a -installsuffix cgo \
        -o bootstrap apt.go

FROM scratch

COPY --from=builder /dist/bootstrap /
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENV TZ=Asia/Shanghai
ENV LANG=C.UTF-8
ENTRYPOINT ["/bootstrap"]

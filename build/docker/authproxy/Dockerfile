FROM golang:1.19.1 as builder

WORKDIR /opt/build
COPY . .

ENV GOPROXY https://goproxy.cn


RUN go mod vendor && go build -o authproxy ./cmd/authproxy/app.go

FROM ubuntu:20.04

WORKDIR /opt/app

COPY --from=builder /opt/build/authproxy ./authproxy
COPY scripts/entrypoint/authproxy/entrypoint.sh ./entrypoint.sh

RUN chmod +x ./entrypoint.sh

VOLUME /config

SHELL ["/bin/bash", "-c"]
ENTRYPOINT ["./entrypoint.sh"]
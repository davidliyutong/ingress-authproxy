FROM golang:1.19.1 as builder

WORKDIR /opt/build
COPY . .

ENV GOPROXY https://goproxy.cn


RUN go mod vendor && go build -ldflags "-X ingress-authproxy/pkg/version.GitVersion=$(git describe --tags --always --dirty)" -o authproxy ./cmd/authproxy/app.go

FROM ubuntu:20.04

WORKDIR /opt/app

COPY --from=builder /opt/build/authproxy ./authproxy
COPY manifests/authproxy-backend/entrypoint/entrypoint.sh ./entrypoint.sh

RUN chmod +x ./entrypoint.sh

VOLUME /config

SHELL ["/bin/bash", "-c"]
ENTRYPOINT ["./entrypoint.sh"]
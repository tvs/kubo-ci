FROM alpine:latest
MAINTAINER pcf-kubo@pivotal.io

RUN apk update && apk add \
      bash \
      curl \
      less \
      groff \
      jq \
      python \
      py-pip \
      py2-pip && \
      pip install --upgrade pip awscli && \
      mkdir /root/.aws

# BOSH CLI
RUN curl https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-2.0.44-linux-amd64 -o bosh && \
  install bosh /usr/local/bin && \
  ln -s /usr/local/bin/bosh /usr/local/bin/bosh-cli

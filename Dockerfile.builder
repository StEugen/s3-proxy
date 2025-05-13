FROM golang:1.23.4-bullseye AS builder

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
       curl git make ca-certificates \
    && rm -rf /var/lib/apt/lists/*

ENV GOPATH=/go \
    PATH=/go/bin:/usr/local/go/bin:$PATH

ARG GOLANGCI_LINT_VERSION="v2.0.2"
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
    | sh -s -- -b /go/bin ${GOLANGCI_LINT_VERSION}

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN mkdir -p bin && \
    go build \
      -o ./bin/s3-proxy \
      -ldflags="-w -s" \
      ./cmd/s3-proxy

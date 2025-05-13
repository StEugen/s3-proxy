###############################################################################
#                                  Builder                                    #
###############################################################################
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
ARG VERSION="steugen@4.0.4"
RUN mkdir -p bin \
    && CGO_ENABLED=0 go build \
    -o bin/s3-proxy \
    -ldflags "\
    -X github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/version.Version=${VERSION} \
    -w -s" \
    ./cmd/s3-proxy

###############################################################################
#                                   Runtime                                   #
###############################################################################
FROM debian:bookworm-slim

RUN groupadd --system s3proxy \
    && useradd --system --gid s3proxy --home-dir /nonexistent --shell /usr/sbin/nologin s3proxy

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /src/bin/s3-proxy /usr/local/bin/s3-proxy
COPY --from=builder --chown=s3proxy:s3proxy /src/templates /templates

USER s3proxy:s3proxy

ENTRYPOINT ["s3-proxy"]

# Build PhoenixChain in a stock Go builder container
FROM golang:1.16-alpine3.13 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers g++ llvm bash cmake git gmp-dev openssl-dev

ADD . /Phoenix-Chain-Core
RUN cd /Phoenix-Chain-Core && make clean && make phoenixchain

# Pull PhoenixChain into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates libstdc++ bash tzdata gmp-dev
COPY --from=builder /Phoenix-Chain-Core/build/bin/phoenixchain /usr/local/bin/

VOLUME /data/phoenixchain
EXPOSE 6060 6789 6790 16789 16789/udp
CMD ["phoenixchain"]
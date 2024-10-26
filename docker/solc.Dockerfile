FROM ethereum/solc:stable-alpine

WORKDIR /nft/solidity

RUN apk update && apk upgrade && \
    apk add --no-cache alpine-sdk bash gcc git \
            libc6-compat libressl-dev musl-dev ca-certificates

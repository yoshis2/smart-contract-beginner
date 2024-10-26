FROM ethereum/client-go:alltools-latest

COPY --from=golang:1-alpine /usr/local/go/ /usr/local/go/
 
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH /go
ENV GOBIN /go/bin

RUN apk update

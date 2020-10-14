FROM golang:1.13 AS builder

RUN apt-get update && apt-get -y install upx 

#RUN mkdir -p $GOPATH/src/github.com/DTherHtun/google-chat-action
#ADD . $GOPATH/src/github.com/DTherHtun/google-chat-action
RUN go get -u github.com/sethvargo/go-githubactions/...
ENV CGO_ENABLED=0
COPY . .

RUN go build \
  -a \
  -trimpath \
  -ldflags "-s -w -extldflags '-static'" \
  -installsuffix cgo \
  -tags netgo \
  -o /bin/google-chat-action \
  . 

RUN strip /bin/google-chat-action

RUN upx -q -9 /bin/google-chat-action

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /bin/google-chat-action /bin/google-chat-action

ENTRYPOINT ["/bin/google-chat-action"]

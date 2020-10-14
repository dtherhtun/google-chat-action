FROM golang:1.13 AS builder

RUN apt-get update && apt-get -y install upx

ENV GO111MODULE=on CGO_ENABLED=0

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

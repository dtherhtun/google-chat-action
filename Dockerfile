FROM golang:1.18 AS builder

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

RUN apt-get -qq update && \
  apt-get -yqq install upx

WORKDIR /src
COPY . .

RUN go build \
  -trimpath \
  -ldflags "-s -w -extldflags '-static'" \
  -o /bin/app \
  .

# RUN strip /bin/app
RUN upx -q -9 /bin/app

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/app /app

ENTRYPOINT ["/app"]

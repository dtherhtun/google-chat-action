FROM golang:1.13

WORKDIR /src
COPY . .

RUN go get -u github.com/sethvargo/go-githubactions/... && go build -o /bin/google-chat-action

ENTRYPOINT ["/bin/google-chat-action"]

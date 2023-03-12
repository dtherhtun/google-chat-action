docker build -t tmp .

docker tag tmp:latest mattnotarangelo/google-chat-action-github-pr-review:latest

docker push mattnotarangelo/google-chat-action-github-pr-review:latest

# Build image
docker build -t mattnotarangelo/google-chat-action-github-pr-review .

# Push to docker hub. Make sure you're logged in
docker push mattnotarangelo/google-chat-action-github-pr-review:latest

# Clean up
docker image rm mattnotarangelo/google-chat-action-github-pr-review

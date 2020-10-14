# google-chat-action

Sample

```yaml
name: Sample Testing
on: [push]

jobs:
  my_job:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout
      uses: actions/checkout@v2
    - name: Chat Setup
      uses: DTherHtun/google-chat-action@v0.1
      with:
        msg: ${GITHUB_REF}
        webhook: "https://chat.googleapis.co....."
```

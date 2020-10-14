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
    - run: echo ::set-output name=action_msg::Someone $GITHUB_EVENT_NAME to $GITHUB_REPOSITORY - $GITHUB_REF - commitid $GITHUB_SHA
      id: txt
    - name: Chat Setup
      uses: DTherHtun/google-chat-action@v0.4
      with:
        msg: ${{ steps.txt.outputs.action_msg }}
        webhook: "https://chat.googleapis.co....."
```
OR

for faster.
```yaml
name: Sample Testing
on: [push]

jobs:
  my_job:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout
      uses: actions/checkout@v2
    - run: echo ::set-output name=action_msg::Someone $GITHUB_EVENT_NAME to $GITHUB_REPOSITORY - $GITHUB_REF - commitid $GITHUB_SHA
      id: txt
    - name: Chat Setup
      uses: docker://dther/google-chat-action:latest
      with:
        msg: ${{ steps.txt.outputs.action_msg }}
        webhook: "https://chat.googleapis.co....."
```

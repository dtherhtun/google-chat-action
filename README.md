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
      uses: DTherHtun/google-chat-action@v0.6
      with:
        project: ${{ github.repository }}
        commit: ${{ github.sha }}
        branch: ${{ github.ref }}
        status: ${{ job.status }}
        actionid: ${{ github.action }}
        webhook: "https://chat.googleapis.com/v1/spaces/AAAAzPcAy4s/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=MmdzluicdrdkyUAV_QwB6BzlLcIhbfrwNzxVrEllaec%3D&threadKey=git-commit"



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

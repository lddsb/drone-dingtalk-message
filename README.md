# Drone CI DingTalk Message Plugin

### Drone CI Plugin Config
`0.8.x`
```yaml
pipeline:
  ...
  notification:
    image: lddsb/drone-dingtalk-message
    token: your-group-bot-token
    type: markdown
```

`1.0.x`
```yaml
kind: pipeline
name: default

steps:
...
- name: notification
  image: lddsb/drone-dingtalk-message
  settings:
    token: your-groupbot-token
    type: markdown

```

### Screen Shot
- Send Success

![send-success](https://i.imgur.com/cECppkW.jpg)

- Missing Access Token

![missing-access-token](https://i.imgur.com/Su7iiyw.jpg)

- Missing Message Type Or Not Support Message Type

![message-type-error](https://i.imgur.com/qtJ4DsA.jpg)

-  Markdown DingTalk Message

![markdown](https://i.imgur.com/LhenKf5.jpg)

- Markdown DingTalk Message(beta tag)

![markdown-massage-beta-tag](https://i.imgur.com/zYuc8hc.jpg)

### Todo

- Multi-Type
- Multi-Lang
- More User Customization


### Development

- First get this repo
```shell
go get github.com/lddsb/drone-dingtalk-message
```
- get dependent lib
```shell
dep ensure
```
- build
```shell
cd $GOPATH/src/github.com/lddsb/drone-dingtalk-message && go build .
```
- run
```shell
./drone-dingtalk-message -h
```
# Drone CI DingTalk Message Plugin
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/lddsb/drone-dingtalk-message)](https://hub.docker.com/r/lddsb/drone-dingtalk-message) [![Go Report Card](https://goreportcard.com/badge/github.com/lddsb/drone-dingtalk-message)](https://goreportcard.com/report/github.com/lddsb/drone-dingtalk-message) [![codecov](https://codecov.io/gh/lddsb/drone-dingtalk-message/branch/master/graph/badge.svg)](https://codecov.io/gh/lddsb/drone-dingtalk-message) [![LICENSE: MIT](https://img.shields.io/github/license/lddsb/drone-dingtalk-message.svg?style=flat-square)](LICENSE)

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

### Plugin Parameter Reference
`token`(required)

String. Access token for group bot. (you can get the access token when you add a bot in a group)

`type`(required)

String. Message type, plan support text, markdown, link and action card, but due to time issue, it's only support `markdown` and `text` now, and you can get the best experience by use markdown.

`message_color`(when `type=markdown`)

Boolean value. This option can change the title and commit message color if turn on.

`success_color`(when `message_color=true`)

String. You can customize the color for the `build success` message by this option, you should input a hex color, example: `008000`.

`failure_color`(when `message_color=true`)

String. You can customize the color for the `build success` message by this option, you should input a hex color, example: `FF0000`.

`sha_link`(when `type=markdown`)

Boolean value. This option can link the sha to your source page when it turn on.

`message_pic`(when `type=markdown`)

Boolean value. If this option turn on,  it will embed a image into the message.

`success_pic`(when `message_pic=true`)

String. You can customize the picture for the `build success` message by this option.

`failure_pic`(when `message_pic=true`)

String. You can customize the picture for the `build failure` message by this option.

### Screen Shot
- Send Success

![send-success](https://i.imgur.com/cECppkW.jpg)

- Missing Access Token

![missing-access-token](https://i.imgur.com/Su7iiyw.jpg)

- Missing Message Type Or Not Support Message Type

![message-type-error](https://i.imgur.com/qtJ4DsA.jpg)

- Markdown DingTalk Message(default)

![markdown-message-default](https://i.imgur.com/Bl7cT1y.jpg)

- Markdown DingTalk Message(color and sha link)

![markdown-massage-customize](https://i.imgur.com/pzdFzIw.jpg)

- Markdown DingTalk Message(color, pic and sha link)

![markdown-massage-customize](https://i.imgur.com/xFrCTZp.jpg)

### Todo

- Multi-Type
- Multi-Lang
- More User Customization


### Development
We use `go mod` to manage dependencies, so it's easy to build.

- get this repo
```shell
$ git clone https://github.com/lddsb/drone-dingtalk-message.git /path/to/you/want
```
- build
```shell
$ cd /path/to/you/want && GO111MODULE=on go build .
```
- run
```shell
$ ./drone-dingtalk-message -h
```

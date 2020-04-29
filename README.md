# Drone CI DingTalk Message Plugin
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/lddsb/drone-dingtalk-message)](https://hub.docker.com/r/lddsb/drone-dingtalk-message) [![Go Report Card](https://goreportcard.com/badge/github.com/lddsb/drone-dingtalk-message)](https://goreportcard.com/report/github.com/lddsb/drone-dingtalk-message) [![codecov](https://codecov.io/gh/lddsb/drone-dingtalk-message/branch/master/graph/badge.svg)](https://codecov.io/gh/lddsb/drone-dingtalk-message) [![Dependabot](https://api.dependabot.com/badges/status?host=github&repo=lddsb/drone-dingtalk-message&identifier=159822771)](https://app.dependabot.com/accounts/lddsb/repos/159822771) [![LICENSE: MIT](https://img.shields.io/github/license/lddsb/drone-dingtalk-message.svg?style=flat-square)](LICENSE)

[中文说明](README_ZH.md)

<!-- toc -->

- [Drone CI Plugin Config](#drone-ci-plugin-config)
- [Plugin Parameter Reference](#plugin-parameter-reference)
- [TPL](#tpl)
- [Screen Shot](#screen-shot)
- [Development](#development)
- [Todo](#todo)
- [Kubernetes Users](#kubernetes-users)

<!-- tocstop -->

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

`1.x`
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
    secret: your-secret-for-generate-sign

```

### Plugin Parameter Reference
`token`(required)

String. Access token for group bot. (you can get the access token when you add a bot in a group)

`type`(required)

String. Message type, plan support text, markdown, link and action card, but due to time issue, it's only support `markdown` and `text` now, and you can get the best experience by use markdown.

`secret`

String. Secret for generate sign.

`tpl`

String. Your custom `tpl`, it can be a local path, or a remote http link.

`tips_title`

String. You can customize the title for the message tips, just work when message type is markdown.

`success_color`

String. You can customize the color for the `build success` message by this option, you should input a hex color, example: `008000`.

`failure_color`

String. You can customize the color for the `build success` message by this option, you should input a hex color, example: `FF0000`.

`success_pic`

String. You can customize the picture for the `build success` message by this option.

`failure_pic`

String. You can customize the picture for the `build failure` message by this option.

`tpl_commit_branch_name`

String. You can customize the [TPL_COMMIT_BRANCH] by this configuration item.

`tpl_repo_short_name`

String. You can customize the [TPL_REPO_SHORT_NAME] by this configuration item.

`tpl_repo_full_name`

String. You can customize the [TPL_REPO_FULL_NAME] by this configuration item.

`tpl_build_status_success`

String. You can customize the [TPL_BUILD_STATUS] (when status=`success`) by this configuration item.

`tpl_build_status_failure`

String. You can customize the [TPL_BUILD_STATUS] (when status=`failure`) by this configuration item.

### TPL
> `tpl` won't work with message type `link` !!!

That's a good news, we support `tpl` now.This is an example for `markdown` message:

	# [TPL_REPO_FULL_NAME] build [TPL_BUILD_STATUS], takes [TPL_BUILD_CONSUMING]s
	[TPL_COMMIT_MSG]

	[TPL_COMMIT_SHA]([TPL_COMMIT_LINK])

	[[TPL_AUTHOR_NAME]([TPL_AUTHOR_EMAIL])](mailto:[TPL_AUTHOR_EMAIL])

	[Click To The Build Detail Page [TPL_STATUS_EMOTICON)]]([TPL_BUILD_LINK])
You can write your own `tpl` what you want. The syntax of	`tpl` is very simple, you can fill `tpl` with preset variables. It's a list of currently supported preset variables:

|       Variable        |                        Value                        |
| :-------------------: | :-------------------------------------------------: |
| [TPL_REPO_SHORT_NAME] |            current repo name(bare name)             |
| [TPL_REPO_FULL_NAME]  |   the full name(with group name) of current repo    |
| [TPL_REPO_GROUP_NAME] |           the group name of current repo            |
| [TPL_REPO_OWNER_NAME] |           the owner name of current repo            |
| [TPL_REPO_REMOTE_URL] |           the remote url of current repo            |
|  [TPL_BUILD_STATUS]   |    current build status(e.g., success, failure)     |
|   [TPL_BUILD_LINK]    |                 current build link                  |
|   [TPL_BUILD_EVENT]   | current build event(e.g., push, pull request, etc.) |
|  [TPL_BUILD_CONSUMING]  |    current build consuming, second     |
|   [TPL_COMMIT_SHA]    |                 current commit sha                  |
|   [TPL_COMMIT_REF]    |  current commit ref(e.g., refs/heads/master, etc.)  |
|   [TPL_COMMIT_LINK]   |           current commit remote url link            |
|  [TPL_COMMIT_BRANCH]  |         current branch name(e.g., dev, etc)         |
|   [TPL_COMMIT_MSG]    |               current commit message                |
|   [TPL_AUTHOR_NAME]   |             current commit author name              |
|  [TPL_AUTHOR_EMAIL]   |             current commit author email             |
| [TPL_AUTHOR_USERNAME] |           current commit author username            |
|  [TPL_AUTHOR_AVATAR]  |            current commit author avatar             |
|   [TPL_STATUS_PIC]    |             custom pic for build status             |
|  [TPL_STATUS_COLOR]   |            custom color for build status            |
| [TPL_STATUS_EMOTICON] |          custom emoticon for build status           |



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

### Todo
It's sad, just support `text`, `markdown` and `link` type now.
- implement all message type

### Kubernetes Users
Attention kubernetes users, [CHANGELOG](CHANGELOG.md#124---2020-04-28).It's the available versions:

- `1.1`(always latest for `1.1.x`)
- `>=1.1.4`
- `1.2`(always latest for `1.2.x`)
- `>=1.2.4`
- latest(always latest)

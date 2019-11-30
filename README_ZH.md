# Drone CI的钉钉群组机器人通知插件
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/lddsb/drone-dingtalk-message)](https://hub.docker.com/r/lddsb/drone-dingtalk-message) [![Go Report Card](https://goreportcard.com/badge/github.com/lddsb/drone-dingtalk-message)](https://goreportcard.com/report/github.com/lddsb/drone-dingtalk-message) [![codecov](https://codecov.io/gh/lddsb/drone-dingtalk-message/branch/master/graph/badge.svg)](https://codecov.io/gh/lddsb/drone-dingtalk-message) [![Dependabot](https://api.dependabot.com/badges/status?host=github&repo=lddsb/drone-dingtalk-message&identifier=159822771)](https://app.dependabot.com/accounts/lddsb/repos/159822771) [![LICENSE: MIT](https://img.shields.io/github/license/lddsb/drone-dingtalk-message.svg?style=flat-square)](LICENSE)

目前仅支持 `text`, `markdown` 以及 `link` 类型的消息，建议使用`markdown`类型。
### 怎么使用本插件
添加一个`step`到你的`.drone.yml`中，下面是简单的例子：

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
    secret: your-secret-for-generate-sign

```

### 插件参数
`token`(必须)

你可以通过加入和创建一个群组来添加钉钉自定义机器人，添加自定义机器人完成后即可获得所需要的`access token`。

`type`(必须)

消息类型，因个人能力有限，目前仅支持`markdown`和`text`，其中，使用`markdown`可以获得最好的体验。

`secret`

如果你设置了`加签`，可以把你的`加签`密钥填入此项完成`加签`操作。

`tpl`

你可以通过该字段来自定义你的消息模版。该字段可以是一个本地路径也可以是一个远程的URL。

`tips_title`

你可以通过该字段自定义钉钉机器人的消息通知提醒标题。（注意，不是消息内容的标题，是收到钉钉机器人发的消息后，会有一个外显的标题）

`success_color`

你可以通过该字段自定义打包成功的颜色。比如：`008000`。

`failure_color`

你可以通过该字段自定义打包失败的颜色。比如：`FF0000`。

`success_pic`

你可以通过该字段自定义打包成功的图片。

`failure_pic`

字符串，你可以通过该字段自定义打包失败的图片。

`tpl_commit_branch_name`

你可以通过该字段自定义分支的名称，可以在模版中通过[TPL_COMMIT_BRANCH]来使用该值。

`tpl_repo_short_name`

你可以通过该字段自定义仓库的名字，可以在模版中通过[TPL_REPO_SHORT_NAME]来使用该值。

`tpl_repo_full_name`

你可以通过该字段自定义仓库的全名（包含组织名称），可以在模版中通过[TPL_REPO_FULL_NAME]来使用该值。

`tpl_build_status_success`

你可以通过该字段自定义运行成功状态的值，可以在模版中通过[TPL_BUILD_STATUS]来使用该值。（仅当前方`step`运行结果为成功时该值会生效）

`tpl_build_status_failure`

你可以通过该字段自定义运行失败状态的值，可以在模版中通过[TPL_BUILD_STATUS]来使用该值。（仅当前方`step`运行结果为失败时该值会生效）

### 模版
> `tpl` 对 `link` 类型的消息并不支持 !!!

感天动地，我们终于支持自定义模版了！下面是一个`markdown`的自定义模版例子：

	# [TPL_REPO_FULL_NAME] build [TPL_BUILD_STATUS], takes [TPL_BUILD_CONSUMING]s
	[TPL_COMMIT_MSG]

	[TPL_COMMIT_SHA]([TPL_COMMIT_LINK])

	[[TPL_AUTHOR_NAME]([TPL_AUTHOR_EMAIL])](mailto:[TPL_AUTHOR_EMAIL])

	[Click To The Build Detail Page [TPL_STATUS_EMOTICON)]]([TPL_BUILD_LINK])

你可以写自己喜欢的模版，终于不用再对着默认模版发愁啦！并且模版的语法非常简单！比较可惜的是目前支持的变量还比较少，下面是当前支持的变量的列表：

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



### 截图展示
- 发送成功（Drone Web）

![send-success](https://i.imgur.com/cECppkW.jpg)

- 忘记填写Access Token（Drone Web）

![missing-access-token](https://i.imgur.com/Su7iiyw.jpg)

- 忘记填写消息类型或者不支持的消息类型

![message-type-error](https://i.imgur.com/qtJ4DsA.jpg)

- 默认的`markdown`消息

![markdown-message-default](https://i.imgur.com/Bl7cT1y.jpg)

- 带颜色和链接的`markdown`消息

![markdown-massage-customize](https://i.imgur.com/pzdFzIw.jpg)

- 带颜色、链接和图片的`markdown`消息

![markdown-massage-customize](https://i.imgur.com/xFrCTZp.jpg)


### 贡献代码
本项目使用了`go mod`来管理依赖，因此要编译本项目相当简单。

- 先把项目代码拷贝到本地
```shell
$ git clone https://github.com/lddsb/drone-dingtalk-message.git /path/to/you/want
```
- 然后直接执行编译即可
```shell
$ cd /path/to/you/want && GO111MODULE=on go build .
```
- 跑个`help`
```shell
$ ./drone-dingtalk-message -h
```

### 待办
- 实现更多的消息类型

# Drone CI的钉钉群组机器人通知插件
[![GitHub Actions](https://github.com/lddsb/drone-dingtalk-message/workflows/Publish%20to%20DockerHub%20and%20Github%20Package/badge.svg)](https://github.com/lddsb/drone-dingtalk-message/actions?query=workflow%3A%22Publish+to+DockerHub+and+Github+Package%22) [![Go Report Card](https://goreportcard.com/badge/github.com/lddsb/drone-dingtalk-message)](https://goreportcard.com/report/github.com/lddsb/drone-dingtalk-message) [![codecov](https://codecov.io/gh/lddsb/drone-dingtalk-message/branch/master/graph/badge.svg)](https://codecov.io/gh/lddsb/drone-dingtalk-message) [![Dependabot](https://api.dependabot.com/badges/status?host=github&repo=lddsb/drone-dingtalk-message&identifier=159822771)](https://app.dependabot.com/accounts/lddsb/repos/159822771) [![LICENSE: MIT](https://img.shields.io/github/license/lddsb/drone-dingtalk-message.svg?style=flat-square)](LICENSE)


<!-- toc -->

- [怎么使用本插件](#%E6%80%8E%E4%B9%88%E4%BD%BF%E7%94%A8%E6%9C%AC%E6%8F%92%E4%BB%B6)
- [插件参数](#%E6%8F%92%E4%BB%B6%E5%8F%82%E6%95%B0)
- [模版](#%E6%A8%A1%E7%89%88)
- [截图展示](#%E6%88%AA%E5%9B%BE%E5%B1%95%E7%A4%BA)
- [贡献代码](#%E8%B4%A1%E7%8C%AE%E4%BB%A3%E7%A0%81)
- [未来计划](#%E6%9C%AA%E6%9D%A5%E8%AE%A1%E5%88%92)
- [Kubernetes 用户请注意](#kubernetes-%E7%94%A8%E6%88%B7%E8%AF%B7%E6%B3%A8%E6%84%8F)

<!-- tocstop -->

### 怎么使用本插件
添加一个`step`到你的`.drone.yml`中，下面是例子：

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
steps:
...
- name: notification
  image: lddsb/drone-dingtalk-message
  settings:
    token: your-groupbot-token
    type: markdown
    secret: your-secret-for-generate-sign
    debug: true
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

`debug`

通过该值可以打开`debug`模式，打印所有环境变量。

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

`msg_at_mobiles`

你需要@的群成员的手机号，多个时用英文逗号(`,`)分隔。如过你使用的是 `markdown` 类型的消息，则需要在 `tpl` 文件中加入 `@手机号` 的内容。

### 模版
> `tpl` 对 `link` 类型的消息并不支持 !!!

感天动地，我们终于支持自定义模版了！下面是一个`markdown`的自定义模版例子：

	# [TPL_REPO_FULL_NAME] build [TPL_BUILD_STATUS], takes [TPL_BUILD_CONSUMING]s
    @mobile1 @mobile2
	[TPL_COMMIT_MSG]

	[TPL_COMMIT_SHA]([TPL_COMMIT_LINK])

	[[TPL_AUTHOR_NAME]([TPL_AUTHOR_EMAIL])](mailto:[TPL_AUTHOR_EMAIL])

	[Click To The Build Detail Page [TPL_STATUS_EMOTICON)]]([TPL_BUILD_LINK])

`mobile1` 和 `mobile2` 应该为钉钉对应的手机号码，可以放在自己想要放的位置。

你可以写自己喜欢的模版，终于不用再对默认模版发愁啦！并且模版的语法非常简单！比较可惜的是目前支持的变量还比较少，下面是当前支持的变量的列表：

|       Variable        |                        Value                        |
| :-------------------: | :-------------------------------------------------: |
| [TPL_REPO_SHORT_NAME] |            当前仓库的名称，比如本仓库 `drone-dingtalk-message`             |
| [TPL_REPO_FULL_NAME]  |   当前仓库的名称，比如本仓库 `lddsb/drone-dingtalk-message`    |
| [TPL_REPO_GROUP_NAME] |           当前仓库的组织名称，比如本仓库 `lddsb`            |
| [TPL_REPO_OWNER_NAME] |           当前仓库拥有者的名称            |
| [TPL_REPO_REMOTE_URL] |           当前仓库的远程地址            |
|  [TPL_BUILD_STATUS]   |    当前编译的状态(比如, success, failure)     |
|   [TPL_BUILD_LINK]    |                 当前编译的链接                  |
|   [TPL_BUILD_EVENT]   | 触发当前编译的动作(比如, push, pull request等) |
|  [TPL_BUILD_CONSUMING]  |    当前编译耗时，单位秒     |
|   [TPL_COMMIT_SHA]    |                 当前提交的sha                  |
|   [TPL_COMMIT_REF]    |  当前提交的ref(比如, refs/heads/master等)  |
|   [TPL_COMMIT_LINK]   |           当前提交的远程地址            |
|  [TPL_COMMIT_BRANCH]  |         当前分之名称(比如, dev, master等)         |
|   [TPL_COMMIT_MSG]    |               当前提交的信息                |
|   [TPL_AUTHOR_NAME]   |             当前提交作者名称              |
|  [TPL_AUTHOR_EMAIL]   |             当前提交作者邮箱地址             |
| [TPL_AUTHOR_USERNAME] |           当前提交作者的用户名            |
|  [TPL_AUTHOR_AVATAR]  |            当前提交作者的头像             |
|   [TPL_STATUS_PIC]    |             根据编译状态显示不同的图片             |
|  [TPL_STATUS_COLOR]   |            根据编译状态显示不同的颜色            |
| [TPL_STATUS_EMOTICON] |          根据编译状态显示不同的表情，比如 `:)` `:(`           |



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

### 未来计划
目前仅支持 `text`, `markdown` 以及 `link` 类型的消息，建议使用`markdown`类型。
- 实现更多的消息类型
- i18N国际化直接翻译环境变量
- 批量发送给多个群机器人
- 失败重试机制

### Kubernetes 用户请注意
因为`Drone CI` [官方缺陷](https://docs.drone.io/runner/kubernetes/overview) ，所以较早版本将无法正常获取到需要用到的变量，会导致部分功能异常。为了能正常使用，所以请使用以下版本：
- `1.1`(总会是`1.1.x`的最新版本)
- `>=1.1.4`
- `1.2`(总会是`1.2.x`的最新版本)
- `>=1.2.4`

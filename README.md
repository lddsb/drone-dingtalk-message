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
- Drone Build Step

![build-step](https://ws1.sinaimg.cn/large/006tNbRwgy1fym86hefglj30eo04i749.jpg)


-  Markdown DingTalk Message

![markdown](https://ws3.sinaimg.cn/large/006tNbRwgy1fym82mg57fj30bo04pjrd.jpg)


- Markdown DingTalk Message(beta tag)

![markdown-massage-beta-tag](https://ws3.sinaimg.cn/large/006tNc79gy1fzgcennwy3j30a00abwf3.jpg)

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
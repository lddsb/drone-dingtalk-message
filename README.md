# Drone CI DingTalk Message Plugin
### Development

> First get this repo
```shell
go get github.com/lddsb/drone-dingtalk-message
```
> get dependent lib
```shell
go get github.com/urfave/cli
```
> build
```shell
cd $GOPATH/src/github.com/lddsb/drone-dingtalk-message && go build .
```
> run
```shell
./drone-dingtalk-notification -h
```
### Drone CI Plugin Config
```yaml
pipeline:
    # other step here
    message:
      image: lddsb/drone-dingtalk-message
      environment:
        - PLUGIN_ACCESS_TOKEN=xxx
        - PLUGIN_MSG_TYPE=markdown
```

### Drone CI Plugin Configs

|ENV|description|is require|
|:-:|:-:|:-:|:-:|
|PLUGIN_ACCESS_TOKEN|dingtalk access token|Yes|
|PLUGIN_MSG_TYPE|dingtalk message type, optional: text, markdown, link, actionCard, feedCard|Yes|
|PLUGIN_MSG_AT_ALL|dingtalk message at all in a group by robot(`default close`)|No|
|PLUGIN_MSG_AT_MOBILES|dingtalk message at anyone in a group by mobiles(`default empty`)|No|
|PLUGIN_LANG|dingtalk message lang, optional: zh_CN and en_US(`default zh_CN`)|No|

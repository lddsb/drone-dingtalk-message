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
./drone-dingtalk-message -h
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
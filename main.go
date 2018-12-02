package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli"
)

var Version = "0.1.1202"

func main() {
	app := cli.NewApp()
	app.Name = "Drone Dingtalk Message Plugin"
	app.Usage = "Sending message to Dingtalk group by robot using webhook"
	app.Copyright = "© 2018 Dee Luo"
	app.Authors = []cli.Author{
		{
			Name:  "Dee Luo",
			Email: "luodi0128@gmail.com",
		},
	}
	app.Action = run
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config.token",
			Usage:  "dingtalk webhook access token",
			EnvVar: "PLUGIN_ACCESS_TOKEN",
		},
		cli.StringFlag{
			Name:   "config.lang",
			Value:  "zh_CN",
			Usage:  "the lang display (zh_CN or en_US, zh_CN is default)",
			EnvVar: "PLUGIN_LANG",
		},
		cli.StringFlag{
			Name:   "config.message.type",
			Usage:  "dingtalk message type, like text, markdown, action card, link and feed card...",
			EnvVar: "PLUGIN_MSG_TYPE",
		},
		cli.StringFlag{
			Name:   "config.message.at.all",
			Usage:  "at all in a message(only text and markdown type message can at)",
			EnvVar: "PLUGIN_MSG_AT_ALL",
		},
		cli.StringFlag{
			Name:   "config.message.at.mobiles",
			Usage:  "at someone in a dingtalk group need this guy bind's mobile",
			EnvVar: "PLUGIN_MSG_AT_MOBILES",
		},
		cli.BoolFlag{
			Name:   "drone",
			Usage:  "indicates the runtime environment is Drone",
			EnvVar: "DRONE",
		},
		cli.StringFlag{
			Name:   "branch",
			Usage:  "providers the branch for the current build",
			EnvVar: "DRONE_BRANCH",
		},
		// commit args start
		cli.StringFlag{
			Name:   "remote.url",
			Usage:  "git remote url",
			EnvVar: "DRONE_REMOTE_URL",
		},
		cli.StringFlag{
			Name:   "commit.author.avatar",
			Usage:  "providers the author avatar url for the current commit",
			EnvVar: "DRONE_COMMIT_AUTHOR_AVATAR",
		},
		cli.StringFlag{
			Name:   "commit.author.email",
			Usage:  "providers the author email for the current commit",
			EnvVar: "DRONE_COMMIT_AUTHOR_EMAIL",
		},
		cli.StringFlag{
			Name:   "commit.author.name",
			Usage:  "providers the author name for the current commit",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Usage:  "providers the branch for the current build",
			EnvVar: "DRONE_COMMIT_BRANCH",
			Value:  "master",
		},
		cli.StringFlag{
			Name:   "commit.link",
			Usage:  "providers the http link to the current commit in the remote source code management system(e.g.GitHub)",
			EnvVar: "DRONE_COMMIT_LINK",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "providers the commit message for the current build",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Usage:  "providers the reference for the current build",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "providers the commit sha for the current build",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		//  commit args end
		cli.StringFlag{
			Name:   "git.url.http",
			Usage:  "providers the repository git+http url",
			EnvVar: "DRONE_GIT_HTTP_URL",
		},
		cli.StringFlag{
			Name:   "git.url.ssh",
			Usage:  "providers the repository git+ssh url",
			EnvVar: "DRONE_GIT_SSH_URL",
		},
		cli.StringFlag{
			Name:   "machine",
			Usage:  "providers the Drone agent hostname",
			EnvVar: "DRONE_MACHINE",
		},
		cli.StringFlag{
			Name:   "pull.request",
			Usage:  "providers the pull request number for the current build.This value is only set if the build event is of type pull request",
			EnvVar: "DRONE_PULL_REQUEST",
		},
		//  repo args start
		cli.StringFlag{
			Name:   "repo.fullname",
			Usage:  "providers the full name of the repository",
			EnvVar: "DRONE_REPO",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.branch",
			Usage:  "providers the default repository branch(e.g.master)",
			EnvVar: "DRONE_REPO_BRANCH",
		},
		cli.StringFlag{
			Name:   "repo.link",
			Usage:  "providers the repository http link",
			EnvVar: "DRONE_REPO_LINK",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "providers the repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "repo.avatar",
			Usage:  "repository avatar",
			EnvVar: "DRONE_REPO_AVATAR",
		},
		cli.StringFlag{
			Name:   "repo.namespace",
			Usage:  "providers the repository namespace(e.g. account owner)",
			EnvVar: "DRONE_REPO_NAMESPACE",
		},
		cli.BoolFlag{
			Name:   "repo.private",
			Usage:  "indicates the repository is public or private",
			EnvVar: "DRONE_REPO_PRIVATE",
		},
		cli.BoolFlag{
			Name:   "repo.trusted",
			Usage:  "repository is trusted",
			EnvVar: "DRONE_REPO_TRUSTED",
		},
		// repo args end
		cli.StringFlag{
			Name:   "runner.host",
			Usage:  "provider are Drone agent hostname",
			EnvVar: "DRONE_RUNNER_HOST",
		},
		cli.StringFlag{
			Name:   "runner.hostname",
			Usage:  "providers the Drone agent hostname",
			EnvVar: "DRONE_RUNNER_HOSTNAME",
		},
		cli.StringFlag{
			Name:   "runner.platform",
			Usage:  "providers the Drone agent os and architecture",
			EnvVar: "DRONE_RUNNER_PLATFORM",
		},
		cli.StringFlag{
			Name:   "runner.label",
			Usage:  "404 not found",
			EnvVar: "DRONE_RUNNER_LABEL",
		},
		cli.StringFlag{
			Name:   "source.branch",
			Usage:  "providers the source branch for a pull request",
			EnvVar: "DRONE_SOURCE_BRANCH",
		},
		cli.StringFlag{
			Name:   "target.branch",
			Usage:  "providers the target branch for a pull request",
			EnvVar: "DRONE_TARGET_BRANCH",
		},
		cli.StringFlag{
			Name:   "system.host",
			Usage:  "providers the Drone server hostname",
			EnvVar: "DRONE_SYSTEM_HOST",
		},
		cli.StringFlag{
			Name:   "system.hostname",
			Usage:  "providers the Drone server hostname",
			EnvVar: "DRONE_SYSTEM_HOSTNAME",
		},
		cli.StringFlag{
			Name:   "system.version",
			Usage:  "providers the Drone server version",
			EnvVar: "DRONE_SYSTEM_VERSION",
		},
		cli.StringFlag{
			Name:   "tag",
			Usage:  "providers the tag name for the current build.This value is only set if the build event is of type tag",
			EnvVar: "DRONE_TAG",
		},
		//  build args start
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.IntFlag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.IntFlag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.IntFlag{
			Name:   "build.finished",
			Usage:  "build finished",
			EnvVar: "DRONE_BUILD_FINISHED",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.StringFlag{
			Name:   "build.deploy",
			Usage:  "build deployment target",
			EnvVar: "DRONE_DEPLOY_TO",
		},
		cli.BoolFlag{
			Name:   "yaml.verified",
			Usage:  "build yaml is verified",
			EnvVar: "DRONE_YAML_VERIFIED",
		},
		cli.BoolFlag{
			Name:   "yaml.signed",
			Usage:  "build yaml is signed",
			EnvVar: "DRONE_YAML_SIGNED",
		},
		//  build args end
		cli.Float64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_JOB_STARTED",
		},
		cli.Float64Flag{
			Name:   "job.finished",
			Usage:  "job finished",
			EnvVar: "DRONE_JOB_FINISHED",
		},
	}

	if err := app.Run(os.Args); nil != err {
		log.Println(err)
		os.Exit(1)
	}
}

//  run with args
func run(c *cli.Context) {
	plugin := Plugin{
		//  repo info
		Repo: Repo{
			FullName: c.String("repo.fullname"),
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		//  build info
		Build: Build{
			Action:  c.String("build.action"),
			Number:  c.Int("build.number"),
			Started: c.Float64("build.started"),
			Created: c.Float64("build.created"),
			Event:   c.String("build.event"),
			Status:  c.String("build.status"),
			Link:    c.String("build.link"),
		},
		Commit: Commit{
			Sha:     c.String("commit.sha"),
			Branch:  c.String("commit.branch"),
			Message: c.String("commit.message"),
			Link:    c.String("commit.link"),
			Authors: struct {
				Avatar string
				Email  string
				Name   string
			}{
				Avatar: c.String("commit.author.avatar"),
				Email:  c.String("commit.author.email"),
				Name:   c.String("commit.author.name"),
			},
		},
		//  custom config
		Config: Config{
			AccessToken: c.String("config.token"),
			Lang:        c.String("config.lang"),
			IsAtALL:     c.Bool("config.message.at.all"),
			MsgType:     c.String("config.message.type"),
			Mobiles:     c.String("config.message.at.mobiles"),
		},
	}

	if err := plugin.Exec(); nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
}

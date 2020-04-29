package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

// Version of cli
var Version = "0.2.1130"

func main() {
	app := cli.NewApp()
	app.Name = "Drone DingTalk Message Plugin"
	app.Usage = "Sending message to DingTalk group by robot using WebHook"
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
		cli.BoolFlag{
			Name:   "config.debug",
			Usage:  "debug mode",
			EnvVar: "PLUGIN_DEBUG",
		},
		cli.StringFlag{
			Name:   "config.tips.title",
			Usage:  "customize the tips title",
			EnvVar: "PLUGIN_TIPS_TITLE",
		},
		cli.StringFlag{
			Name:   "config.token,access_token,token",
			Usage:  "DingTalk webhook access token",
			EnvVar: "PLUGIN_ACCESS_TOKEN,PLUGIN_TOKEN",
		},
		cli.StringFlag{
			Name:   "config.secret,secret",
			Usage:  "DingTalk WebHook secret for generate sign",
			EnvVar: "PLUGIN_SECRET",
		},
		cli.StringFlag{
			Name:   "config.message.type,message_type",
			Usage:  "DingTalk message type, like text, markdown, action card, link and feed card...",
			EnvVar: "PLUGIN_MSG_TYPE,PLUGIN_TYPE,PLUGIN_MESSAGE_TYPE",
		},
		cli.StringFlag{
			Name:   "config.message.at.all",
			Usage:  "at all in a message(only text and markdown type message can at)",
			EnvVar: "PLUGIN_MSG_AT_ALL",
		},
		cli.StringFlag{
			Name:   "config.message.at.mobiles",
			Usage:  "at someone in a DingTalk group need this guy bind's mobile",
			EnvVar: "PLUGIN_MSG_AT_MOBILES",
		},
		cli.StringFlag{
			Name:   "commit.author.username",
			Usage:  "providers the author username for the current commit",
			EnvVar: "DRONE_COMMIT_AUTHOR",
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
			Name:   "commit.sha",
			Usage:  "providers the commit sha for the current build",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Usage:  "provider the commit ref for the current build",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "repo.full.name",
			Usage:  "providers the full name of the repository",
			EnvVar: "DRONE_REPO",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "provider the name of the repository",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "repo.group",
			Usage:  "provider the group of the repository",
			EnvVar: "DRONE_REPO_NAMESPACE",
		},
		cli.StringFlag{
			Name:   "repo.remote.url",
			Usage:  "provider the remote url of the repository",
			EnvVar: "DRONE_REMOTE_URL",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "provider the owner of the repository",
			EnvVar: "DRONE_REPO_OWNER",
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
			Name:   "build.event",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.StringFlag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.StringFlag{
			Name:   "build.finished",
			Usage:  "build finished",
			EnvVar: "DRONE_BUILD_FINISHED",
		},
		cli.StringFlag{
			Name:   "tpl.build.status.success",
			Usage:  "tpl.build status for replace success",
			EnvVar: "TPL_BUILD_STATUS_SUCCESS, PLUGIN_TPL_BUILD_STATUS_SUCCESS",
		},
		cli.StringFlag{
			Name:   "tpl.build.status.failure",
			Usage:  "tpl.build status for replace failure",
			EnvVar: "TPL_BUILD_STATUS_FAILURE, PLUGIN_TPL_BUILD_STATUS_FAILURE",
		},
		cli.StringFlag{
			Name:   "custom.pic.url.success",
			Usage:  "custom success picture url",
			EnvVar: "SUCCESS_PICTURE_URL,PLUGIN_SUCCESS_PIC",
		},
		cli.StringFlag{
			Name:   "custom.pic.url.failure",
			Usage:  "custom failure picture url",
			EnvVar: "FAILURE_PICTURE_URL,PLUGIN_FAILURE_PIC",
		},
		cli.StringFlag{
			Name:   "custom.color.success",
			Usage:  "custom success color for title in markdown",
			EnvVar: "SUCCESS_COLOR,PLUGIN_SUCCESS_COLOR",
		},
		cli.StringFlag{
			Name:   "custom.color.failure",
			Usage:  "custom failure color for title in markdown",
			EnvVar: "FAILURE_COLOR,PLUGIN_FAILURE_COLOR",
		},
		cli.StringFlag{
			Name:   "custom.tpl",
			Usage:  "custom tpl",
			EnvVar: "PLUGIN_TPL,PLUGIN_CUSTOM_TPL",
		},
		cli.StringFlag{
			Name:   "tpl.repo.full.name",
			Usage:  "tpl custom repo full name",
			EnvVar: "PLUGIN_TPL_REPO_FULL_NAME,TPL_REPO_FULL_NAME",
		},
		cli.StringFlag{
			Name:   "tpl.repo.short.name",
			Usage:  "tpl custom repo short name",
			EnvVar: "PLUGIN_TPL_REPO_SHORT_NAME,TPL_REPO_SHORT_NAME",
		},
		cli.StringFlag{
			Name:   "tpl.commit.branch.name",
			Usage:  "tpl custom commit branch name",
			EnvVar: "PLUGIN_TPL_COMMIT_BRANCH_NAME,TPL_COMMIT_BRANCH_NAME",
		},
	}

	// kubernetes runner patch
	if _, err := os.Stat("/run/drone/env"); err == nil {
		godotenv.Overload("/run/drone/env")
	}

	if err := app.Run(os.Args); nil != err {
		log.Println(err)
	}
}

//  run with args
func run(c *cli.Context) {
	plugin := Plugin{
		Drone: Drone{
			//  repo info
			Repo: Repo{
				ShortName: c.String("repo.name"),
				GroupName: c.String("repo.group"),
				OwnerName: c.String("repo.owner"),
				RemoteURL: c.String("repo.remote.url"),
				FullName:  c.String("repo.full.name"),
			},
			//  build info
			Build: Build{
				Status:     c.String("build.status"),
				Link:       c.String("build.link"),
				Event:      c.String("build.event"),
				StartAt:    c.Int64("build.started"),
				FinishedAt: c.Int64("build.finished"),
			},
			Commit: Commit{
				Sha:     c.String("commit.sha"),
				Branch:  c.String("commit.branch"),
				Message: c.String("commit.message"),
				Link:    c.String("commit.link"),
				Author: CommitAuthor{
					Avatar:   c.String("commit.author.avatar"),
					Email:    c.String("commit.author.email"),
					Name:     c.String("commit.author.name"),
					Username: c.String("commit.author.username"),
				},
			},
		},
		//  custom config
		Config: Config{
			AccessToken: c.String("config.token"),
			Secret:      c.String("config.secret"),
			IsAtALL:     c.Bool("config.message.at.all"),
			MsgType:     c.String("config.message.type"),
			Mobiles:     c.String("config.message.at.mobiles"),
			Debug:       c.Bool("config.debug"),
			TipsTitle:   c.String("config.tips.title"),
		},
		Custom: Custom{
			Pic: Pic{
				SuccessPicURL: c.String("custom.pic.url.success"),
				FailurePicURL: c.String("custom.pic.url.failure"),
			},
			Color: Color{
				SuccessColor: c.String("custom.color.success"),
				FailureColor: c.String("custom.color.failure"),
			},
			Tpl: c.String("custom.tpl"),
		},
		Tpl: Tpl{
			Repo: TplRepo{
				FullName:  c.String("tpl.repo.full.name"),
				ShortName: c.String("tpl.repo.short.name"),
			},
			Commit: TplCommit{
				Branch: c.String("tpl.commit.branch.name"),
			},
			Build: TplBuild{
				Status: Status{
					Success: c.String("tpl.build.status.success"),
					Failure: c.String("tpl.build.status.failure"),
				},
			},
		},
	}

	if err := plugin.Exec(); nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli"
)

// Version of cli
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
		cli.BoolFlag{
			Name:   "config.debug",
			Usage:  "debug mode",
			EnvVar: "PLUGIN_DEBUG",
		},
		cli.StringFlag{
			Name:   "config.token,access_token,token",
			Usage:  "dingtalk webhook access token",
			EnvVar: "PLUGIN_ACCESS_TOKEN,PLUGIN_TOKEN",
		},
		cli.StringFlag{
			Name:   "config.lang",
			Value:  "zh_CN",
			Usage:  "the lang display (zh_CN or en_US, zh_CN is default)",
			EnvVar: "PLUGIN_LANG",
		},
		cli.StringFlag{
			Name:   "config.message.type,message_type",
			Usage:  "dingtalk message type, like text, markdown, action card, link and feed card...",
			EnvVar: "PLUGIN_MSG_TYPE,PLUGIN_TYPE,PLUGIN_MESSAGE_TYPE",
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
			Name:   "repo.fullname",
			Usage:  "providers the full name of the repository",
			EnvVar: "DRONE_REPO",
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
			Name:   "config.success.pic.url",
			Usage:  "config success picture url",
			EnvVar: "SUCCESS_PICTURE_URL,PLUGIN_SUCCESS_PIC",
		},
		cli.StringFlag{
			Name:   "config.failure.pic.url",
			Usage:  "config failure picture url",
			EnvVar: "FAILURE_PICTURE_URL,PLUGIN_FAILURE_PIC",
		},
		cli.StringFlag{
			Name:   "config.success.color",
			Usage:  "config success color for title in markdown",
			EnvVar: "SUCCESS_COLOR,PLUGIN_SUCCESS_COLOR",
		},
		cli.StringFlag{
			Name:   "config.failure.color",
			Usage:  "config failure color for title in markdown",
			EnvVar: "FAILURE_COLOR,PLUGIN_FAILURE_COLOR",
		},
		cli.BoolFlag{
			Name:   "config.message.color",
			Usage:  "configure the message with color or not",
			EnvVar: "PLUGIN_COLOR,PLUGIN_MESSAGE_COLOR",
		},
		cli.BoolFlag{
			Name:   "config.message.pic",
			Usage:  "configure the message with picture or not",
			EnvVar: "PLUGIN_PIC,PLUGIN_MESSAGE_PIC",
		},
		cli.BoolFlag{
			Name:   "config.message.sha.link",
			Usage:  "link sha source page or not",
			EnvVar: "PLUGIN_SHA_LINK,PLUGIN_MESSAGE_SHA_LINK",
		},
		cli.StringFlag{
			Name:   "config.tips.title",
			Usage:  "tips title, just work for markdown type message",
			EnvVar: "PLUGIN_TIPS_TITLE",
		},
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
				FullName: c.String("repo.fullname"),
			},
			//  build info
			Build: Build{
				Status: c.String("build.status"),
				Link:   c.String("build.link"),
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
		},
		//  custom config
		Config: Config{
			AccessToken: c.String("config.token"),
			//Lang:          c.String("config.lang"),
			IsAtALL:   c.Bool("config.message.at.all"),
			MsgType:   c.String("config.message.type"),
			Mobiles:   c.String("config.message.at.mobiles"),
			Debug:     c.Bool("config.debug"),
			TipsTitle: c.String("config.tips.title"),
		},
		Extra: Extra{
			Pic: ExtraPic{
				WithPic:       c.Bool("config.message.pic"),
				SuccessPicURL: c.String("config.success.pic.url"),
				FailurePicURL: c.String("config.failure.pic.url"),
			},
			Color: ExtraColor{
				SuccessColor: c.String("config.success.color"),
				FailureColor: c.String("config.failure.color"),
				WithColor:    c.Bool("config.message.color"),
			},
			LinkSha: c.Bool("config.message.sha.link"),
		},
	}

	if err := plugin.Exec(); nil != err {
		fmt.Println(err)
	}
}

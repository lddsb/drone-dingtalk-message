package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type (
	//  repo base info
	Repo struct {
		Owner      string //  providers the repository owner name
		Name       string //  providers the repository name
		Branch     string //  providers the default repository branch(e.g.master)
		Link       string //  providers the repository http link
		NameSpace  string //  providers the repository namespace(e.g.account owner)
		Private    bool   //  indicates the repository is public or private
		Visibility string //  providers the repository visibility level.Possible values are public,private and internal
		SCM        string //  providers the repository version control system
		FullName   string //  repository full name
	}
	//  build info
	Build struct {
		Action  string  //  document description not found
		Created float64 //  providers the date and time when the build was created in the system
		Event   string  //  providers the current build event
		Number  int     //  providers the current build number
		Started float64 //  providers the date and time when the build was started
		Status  string  //  providers the current build status
		Link    string  //  providers the current build link
	}
	//  commit info
	Commit struct {
		After   string //  providers the commit sha for the current build
		Author  string //  providers the author username for the current commit
		Before  string //  providers the parent commit sha for the current build
		Branch  string //  providers the branch for the current commit
		Link    string //  providers the http link to the current commit in the remote source code management system(e.g.GitHub)
		Message string //  providers the commit message for the current build
		Ref     string //  providers the reference for the current build
		Sha     string //  providers the commit sha for the current build
		//  repo author info
		Authors struct {
			Avatar string //  providers the author avatar for the current commit
			Email  string //  providers the author email for the current commit
			Name   string //  providers the author name for the current commit
		}
	}
	//  git url info
	Git struct {
		HttpUrl string //  providers the repository git+http url
		SSHUrl  string //  providers the repository git+ssh url
	}
	//  Drone runner info
	Runner struct {
		Host     string //  providers the Drone agent hostname
		Hostname string //  providers the Drone agent hostname
		Platform string //  providers the Drone agent os and architecture
		Label    string //  document description not found
	}
	//  Drone system info
	System struct {
		Host     string //  providers the Drone server hostname
		Hostname string //  providers the Drone server hostname
		Version  string //  providers the Drone server version
	}
	//  Drone CI Info
	CI struct {
		RepoLink string
	}
	//  plugin private config
	Config struct {
		AccessToken    string
		Message        string
		Lang           string
		IsAtALL        bool
		Mobiles        string
		Username       string
		AvatarURL      string
		MsgType        string
		LinkUrls       string
		LinkTitles     string
		HideAvatar     bool
		BtnOrientation bool
		PicURL         string
		MsgURL         string
		SuccessPicUrl  string
		FailurePicUrl  string
		SuccessColor   string
		FailureColor   string
		WithColor      bool
		WithPic        bool
		LinkSha        bool
		RetryTime      int
	}
	// plugin all config
	Plugin struct {
		CI      CI
		Git     Git
		Runner  Runner
		System  System
		Commit  Commit
		Repo    Repo
		Build   Build
		Config  Config
		WebHook *WebHook
	}
)

func (p *Plugin) Exec() error {
	log.Println("start execute sending...")
	if 0 == len(p.Config.AccessToken) {
		msg := "missing dingtalk access token"
		log.Println(msg)
		return errors.New(msg)
	}
	log.Println("access token pass...")
	p.WebHook = NewWebHook(p.Config.AccessToken)
	mobiles := strings.Split(p.Config.Mobiles, ",")
	linkUrls := strings.Split(p.Config.LinkUrls, ",")
	linkTitles := strings.Split(p.Config.LinkTitles, ",")
	log.Println("sending message type: " + p.Config.MsgType)
	var err error
	retryTime := 1
	for retryTime <= p.Config.RetryTime {
		log.Printf("start a %d try", retryTime)
		switch strings.ToLower(p.Config.MsgType) {
		case "markdown":
			err = p.WebHook.SendMarkdownMsg(
				"You have a new message...",
				p.baseTpl(),
				p.Config.IsAtALL,
				mobiles...
			)
		case "text":
			err = p.WebHook.SendTextMsg(p.baseTpl(), p.Config.IsAtALL, mobiles...)
		case "actioncard":
			err = p.WebHook.SendActionCardMsg(
				"A actionCard title",
				p.baseTpl(),
				linkUrls,
				linkTitles,
				p.Config.HideAvatar,
				p.Config.BtnOrientation,
			)
		case "link":
			err = p.WebHook.SendLinkMsg(p.Build.Status, p.baseTpl(), p.Commit.Authors.Avatar, p.Build.Link)
		default:
			err = errors.New("not support message type")
		}

		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			retryTime++
		} else {
			break
		}
	}
	if nil != err {
		log.Println(err)
		return err
	}
	log.Println("send " + p.Config.MsgType + " message success!")
	return nil
}

func (p *Plugin) markdownTpl() string {
	var tpl string

	//  title
	title := fmt.Sprintf(" %s *Branch Build %s*",
		strings.Title(p.Commit.Branch),
		strings.Title(p.Build.Status))
	//  with color on title
	if p.Config.WithColor {
		title = fmt.Sprintf("<font color=%s>%s</font>", p.getColor(), title)
	}

	tpl = fmt.Sprintf("# %s \n", title)

	// with pic
	if p.Config.WithPic {
		tpl += fmt.Sprintf("![%s](%s)\n\n",
			p.Build.Status,
			p.getPicUrl())
	}

	//  commit message
	commitMsg := fmt.Sprintf("%s", p.Commit.Message)
	if p.Config.WithColor {
		commitMsg = fmt.Sprintf("<font color=%s>%s</font>", p.getColor(), commitMsg)
	}
	tpl += commitMsg + "\n\n"

	//  sha info
	commitSha := p.Commit.Sha
	if p.Config.LinkSha {
		commitSha = fmt.Sprintf("[Click To %s Commit Detail Page](%s)", commitSha[:6], p.Commit.Link)
	}
	tpl += commitSha + "\n\n"

	//  author info
	authorInfo := fmt.Sprintf("`%s(%s)`", p.Commit.Authors.Name, p.Commit.Authors.Email)
	tpl += authorInfo + "\n\n"

	//  build detail link
	buildDetail := fmt.Sprintf("[Click To The Build Detail Page %s](%s)",
		p.getEmoticon(),
		p.Build.Link)
	tpl += buildDetail
	return tpl
}

func (p *Plugin) baseTpl() string {
	tpl := ""
	switch strings.ToLower(p.Config.MsgType) {
	case "markdown":
		tpl = p.markdownTpl()
	case "text":
		tpl = fmt.Sprintf(`[%s] %s
%s (%s)
@%s
%s (%s)
`,
			p.Build.Status,
			strings.TrimSpace(p.Commit.Message),
			p.Repo.FullName,
			p.Commit.Branch,
			p.Commit.Sha,
			p.Commit.Authors.Name,
			p.Commit.Authors.Email)
	case "link":
		tpl = fmt.Sprintf(`%s(%s) @%s %s(%s)`,
			p.Repo.FullName,
			p.Commit.Branch,
			p.Commit.Sha[:6],
			p.Commit.Authors.Name,
			p.Commit.Authors.Email)
	case "actionCard":
		//  coming soon

	}

	return tpl
}

/**
get emoticon
 */
func (p *Plugin) getEmoticon() string {
	emoticons := make(map[string]string)
	emoticons["success"] = ":)"
	emoticons["failure"] = ":("

	emoticon, ok := emoticons[p.Build.Status]
	if ok {
		return emoticon
	}

	return ":("
}

/**
get picture url
 */
func (p *Plugin) getPicUrl() string {
	pics := make(map[string]string)
	//  success picture url
	pics["success"] = "https://ws4.sinaimg.cn/large/006tNc79gy1fz05g5a7utj30he0bfjry.jpg"
	if p.Config.SuccessPicUrl != "" {
		pics["success"] = p.Config.SuccessPicUrl
	}
	//  failure picture url
	pics["failure"] = "https://ws1.sinaimg.cn/large/006tNc79gy1fz0b4fghpnj30hd0bdmxn.jpg"
	if p.Config.FailurePicUrl != "" {
		pics["failure"] = p.Config.FailurePicUrl
	}

	url, ok := pics[p.Build.Status]
	if ok {
		return url
	}

	return ""
}

/**
get color for message title
 */
func (p *Plugin) getColor() string {
	colors := make(map[string]string)
	//  success color
	colors["success"] = "#008000"
	if p.Config.SuccessColor != "" {
		colors["success"] = "#" + p.Config.SuccessColor
	}
	//  failure color
	colors["failure"] = "#FF0000"
	if p.Config.FailureColor != "" {
		colors["failure"] = "#" + p.Config.FailureColor
	}

	color, ok := colors[p.Build.Status]
	if ok {
		return color
	}

	return ""
}

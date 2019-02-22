package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type (
	// Repo `repo base info`
	Repo struct {
		FullName string //  repository full name
	}
	// Build `build info`
	Build struct {
		Status string //  providers the current build status
		Link   string //  providers the current build link
	}
	// Commit `commit info`
	Commit struct {
		Branch  string //  providers the branch for the current commit
		Link    string //  providers the http link to the current commit in the remote source code management system(e.g.GitHub)
		Message string //  providers the commit message for the current build
		Sha     string //  providers the commit sha for the current build
		//  repo author info
		Authors struct {
			Avatar string //  providers the author avatar for the current commit
			Email  string //  providers the author email for the current commit
			Name   string //  providers the author name for the current commit
		}
	}
	// Config `plugin private config`
	Config struct {
		Debug          bool
		AccessToken    string
		IsAtALL        bool
		Mobiles        string
		Username       string
		MsgType        string
		LinkUrls       string
		LinkTitles     string
		HideAvatar     bool
		BtnOrientation bool
	}
	// Extra `extra variables`
	Extra struct {
		PicURL        string
		MsgURL        string
		SuccessPicUrl string
		FailurePicUrl string
		SuccessColor  string
		FailureColor  string
		WithColor     bool
		WithPic       bool
		LinkSha       bool
	}

	// Plugin `plugin all config`
	Plugin struct {
		Commit  Commit
		Repo    Repo
		Build   Build
		Config  Config
		Extra   Extra
		WebHook *WebHook
	}
)

// Exec `execute webhook`
func (p *Plugin) Exec() error {
	var err error
	if 0 == len(p.Config.AccessToken) {
		msg := "missing dingtalk access token"
		return errors.New(msg)
	}
	p.WebHook = NewWebHook(p.Config.AccessToken)
	mobiles := strings.Split(p.Config.Mobiles, ",")
	switch strings.ToLower(p.Config.MsgType) {
	case "markdown":
		err = p.WebHook.SendMarkdownMsg("You have a new message...", p.baseTpl(), p.Config.IsAtALL, mobiles...)
	case "text":
		err = p.WebHook.SendTextMsg(p.baseTpl(), p.Config.IsAtALL, mobiles...)
	case "link":
		err = p.WebHook.SendLinkMsg(p.Build.Status, p.baseTpl(), p.Commit.Authors.Avatar, p.Build.Link)
	default:
		msg := "not support message type"
		err = errors.New(msg)
	}

	if err != nil {
		return err
	}
	log.Println("send message success!")
	return nil
}

// markdownTpl `output the tpl of markdown`
func (p *Plugin) markdownTpl() string {
	var tpl string

	//  title
	title := fmt.Sprintf(" %s *Branch Build %s*",
		strings.Title(p.Commit.Branch),
		strings.Title(p.Build.Status))
	//  with color on title
	if p.Extra.WithColor {
		title = fmt.Sprintf("<font color=%s>%s</font>", p.getColor(), title)
	}

	tpl = fmt.Sprintf("# %s \n", title)

	// with pic
	if p.Extra.WithPic {
		tpl += fmt.Sprintf("![%s](%s)\n\n",
			p.Build.Status,
			p.getPicUrl())
	}

	//  commit message
	commitMsg := fmt.Sprintf("%s", p.Commit.Message)
	if p.Extra.WithColor {
		commitMsg = fmt.Sprintf("<font color=%s>%s</font>", p.getColor(), commitMsg)
	}
	tpl += commitMsg + "\n\n"

	//  sha info
	commitSha := p.Commit.Sha
	if p.Extra.LinkSha {
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
	if p.Extra.SuccessPicUrl != "" {
		pics["success"] = p.Extra.SuccessPicUrl
	}
	//  failure picture url
	pics["failure"] = "https://ws1.sinaimg.cn/large/006tNc79gy1fz0b4fghpnj30hd0bdmxn.jpg"
	if p.Extra.FailurePicUrl != "" {
		pics["failure"] = p.Extra.FailurePicUrl
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
	if p.Extra.SuccessColor != "" {
		colors["success"] = "#" + p.Extra.SuccessColor
	}
	//  failure color
	colors["failure"] = "#FF0000"
	if p.Extra.FailureColor != "" {
		colors["failure"] = "#" + p.Extra.FailureColor
	}

	color, ok := colors[p.Build.Status]
	if ok {
		return color
	}

	return ""
}

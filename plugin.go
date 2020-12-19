package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	webhook "github.com/lddsb/dingtalk-webhook"
)

type (
	// Repo repo base info
	Repo struct {
		ShortName string //  short name
		GroupName string //  group name
		FullName  string //  repository full name
		OwnerName string //  repo owner
		RemoteURL string //  repo remote url
	}

	// Build build info
	Build struct {
		Status     string //  providers the current build status
		Link       string //  providers the current build link
		Event      string //  trigger event
		StartAt    uint64 //  build start at ( unix timestamp )
		FinishedAt uint64 //  build finish at ( unix timestamp )
	}

	// Commit commit info
	Commit struct {
		Branch  string //  providers the branch for the current commit
		Link    string //  providers the http link to the current commit in the remote source code management system(e.g.GitHub)
		Message string //  providers the commit message for the current build
		Sha     string //  providers the commit sha for the current build
		Ref     string //  commit ref
		Author  CommitAuthor
	}

	// Stage drone stage env
	Stage struct {
		StartedAt  uint64
		FinishedAt uint64
	}

	// CommitAuthor commit author info
	CommitAuthor struct {
		Avatar   string //  providers the author avatar for the current commit
		Email    string //  providers the author email for the current commit
		Name     string //  providers the author name for the current commit
		Username string //  the author username for the current commit
	}

	// Drone drone info
	Drone struct {
		Repo   Repo
		Build  Build
		Commit Commit
		Stage  Stage
	}

	// Config plugin private config
	Config struct {
		Debug       bool
		AccessToken string
		Secret      string
		IsAtALL     bool
		Mobiles     string
		Username    string
		MsgType     string
		TipsTitle   string
	}

	// MessageConfig DingTalk message struct
	MessageConfig struct {
		ActionCard ActionCard
	}

	// ActionCard action card message struct
	ActionCard struct {
		LinkUrls       string
		LinkTitles     string
		HideAvatar     bool
		BtnOrientation bool
	}

	// Pic extra config for pic
	Pic struct {
		SuccessPicURL string
		FailurePicURL string
	}

	// Color extra config for color
	Color struct {
		SuccessColor string
		FailureColor string
	}

	// Plugin plugin all config
	Plugin struct {
		Tpl     Tpl
		Drone   Drone
		Config  Config
		Custom  Custom
		Message MessageConfig
	}

	// Custom user custom env
	Custom struct {
		Tpl       string
		Color     Color
		Pic       Pic
		Consuming Consuming
	}

	// Tpl TPL base
	Tpl struct {
		Repo   TplRepo
		Commit TplCommit
		Build  TplBuild
	}

	// TplRepo TPL repo
	TplRepo struct {
		FullName  string
		ShortName string
	}

	// TplCommit TPL commit
	TplCommit struct {
		Branch string
	}

	// TplBuild TPL build
	TplBuild struct {
		Status Status
	}

	// Status status
	Status struct {
		Success string
		Failure string
	}

	// Consuming custom consuming env
	Consuming struct {
		StartedEnv  string
		FinishedEnv string
	}
)

// Exec execute WebHook
func (p *Plugin) Exec() error {
	if p.Config.Debug {
		for _, e := range os.Environ() {
			log.Println(e)
		}
	}

	var err error
	if "" == p.Config.AccessToken {
		msg := "missing DingTalk access token"
		return errors.New(msg)
	}

	tpl, err := p.getMessage()
	if err != nil {
		return err
	}

	if p.Config.TipsTitle == "" {
		p.Config.TipsTitle = "you have a new message"
	}

	newWebHook := webhook.NewWebHook(p.Config.AccessToken)

	// add sign
	if "" != p.Config.Secret {
		newWebHook.Secret = p.Config.Secret
	}

	mobiles := strings.Split(p.Config.Mobiles, ",")
	switch strings.ToLower(p.Config.MsgType) {
	case "markdown":
		err = newWebHook.SendMarkdownMsg(p.Config.TipsTitle, tpl, p.Config.IsAtALL, mobiles...)
	case "text":
		err = newWebHook.SendTextMsg(tpl, p.Config.IsAtALL, mobiles...)
	case "link":
		err = newWebHook.SendLinkMsg(p.Drone.Build.Status, tpl, p.Drone.Commit.Author.Avatar, p.Drone.Build.Link)
	default:
		msg := "not support message type"
		err = errors.New(msg)
	}

	if err == nil {
		log.Println("send message success!")
	}

	return err
}

// fileExists check file is exists
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// getTpl get tpl from local file or remote file
func (p *Plugin) getTpl() (tpl string, err error) {
	//var tpl string
	tplDir := "/app/drone/dingtalk/message/tpls"
	if "" == p.Custom.Tpl {
		p.Custom.Tpl = fmt.Sprintf("%s/%s.tpl", tplDir, strings.ToLower(p.Config.MsgType))
	}

	u, err := url.Parse(p.Custom.Tpl)
	if err != nil {
		return "", err
	}

	if u.Scheme != "" {
		resp, err := http.Get(p.Custom.Tpl)
		if err != nil {
			return "", err
		}

		// check response
		if u.Path != resp.Request.URL.Path {
			return "", errors.New("cannot get tpl from url")
		}

		// defer close
		defer func() {
			_ = resp.Body.Close()
		}()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		tpl = string(body)
	} else {
		if !fileExists(p.Custom.Tpl) {
			return "", errors.New("tpl file not exists")
		}

		tplStr, err := ioutil.ReadFile(p.Custom.Tpl)
		if err != nil {
			return "", err
		}

		tpl = string(tplStr)
	}

	return tpl, nil
}

// fillTpl fill the tpl by valid keyword
func (p *Plugin) fillTpl(tpl string) string {
	envs := p.getEnvs()
	// replace regex
	reg := regexp.MustCompile(`\[([^\[\]]*)]`)
	match := reg.FindAllStringSubmatch(tpl, -1)
	for _, m := range match {
		// from environment
		if envStr := os.Getenv(m[1]); envStr != "" {
			tpl = strings.ReplaceAll(tpl, m[0], envStr)
		}

		// check if the keyword is legal
		if _, ok := envs[m[1]]; ok {
			// replace keyword
			tpl = strings.ReplaceAll(tpl, m[0], envs[m[1]].(string))
		}
	}

	return tpl
}

// getEnvs get available envs
func (p *Plugin) getEnvs() map[string]interface{} {
	var envs map[string]interface{}
	envs = make(map[string]interface{})
	envs["TPL_REPO_FULL_NAME"] = p.Drone.Repo.FullName
	if p.Tpl.Repo.FullName != "" {
		envs["TPL_REPO_FULL_NAME"] = p.Tpl.Repo.FullName
	}
	envs["TPL_REPO_SHORT_NAME"] = p.Drone.Repo.ShortName
	if p.Tpl.Repo.ShortName != "" {
		envs["TPL_REPO_SHORT_NAME"] = p.Tpl.Repo.ShortName
	}
	envs["TPL_REPO_GROUP_NAME"] = p.Drone.Repo.GroupName
	envs["TPL_REPO_OWNER_NAME"] = p.Drone.Repo.OwnerName
	envs["TPL_REPO_REMOTE_URL"] = p.Drone.Repo.RemoteURL

	envs["TPL_BUILD_STATUS"] = p.getStatus()
	envs["TPL_BUILD_LINK"] = p.Drone.Build.Link
	envs["TPL_BUILD_EVENT"] = p.Drone.Build.Event

	var consuming uint64
	// custom consuming env
	if p.Custom.Consuming.FinishedEnv != "" && p.Custom.Consuming.StartedEnv != "" {
		finishedAt, _ := strconv.ParseUint(os.Getenv(p.Custom.Consuming.FinishedEnv), 10, 64)
		startedAt, _ := strconv.ParseUint(os.Getenv(p.Custom.Consuming.StartedEnv), 10, 64)
		consuming = finishedAt - startedAt
	} else {
		consuming = p.Drone.Build.FinishedAt - p.Drone.Build.StartAt
		if consuming == 0 {
			consuming = p.Drone.Stage.FinishedAt - p.Drone.Stage.StartedAt
		}
	}
	envs["TPL_BUILD_CONSUMING"] = fmt.Sprintf("%v", consuming)

	envs["TPL_COMMIT_SHA"] = p.Drone.Commit.Sha
	envs["TPL_COMMIT_REF"] = p.Drone.Commit.Ref
	envs["TPL_COMMIT_LINK"] = p.Drone.Commit.Link
	envs["TPL_COMMIT_MSG"] = p.Drone.Commit.Message
	envs["TPL_COMMIT_BRANCH"] = p.Drone.Commit.Branch
	if p.Tpl.Commit.Branch != "" {
		envs["TPL_COMMIT_BRANCH"] = p.Tpl.Commit.Branch
	}

	envs["TPL_AUTHOR_NAME"] = p.Drone.Commit.Author.Name
	envs["TPL_AUTHOR_USERNAME"] = p.Drone.Commit.Author.Username
	envs["TPL_AUTHOR_EMAIL"] = p.Drone.Commit.Author.Email
	envs["TPL_AUTHOR_AVATAR"] = p.Drone.Commit.Author.Avatar

	envs["TPL_STATUS_PIC"] = p.getPicURL()
	envs["TPL_STATUS_COLOR"] = p.getColor()
	envs["TPL_STATUS_EMOTICON"] = p.getEmoticon()

	return envs
}

// getMessage get message tpl
func (p *Plugin) getMessage() (tpl string, err error) {
	tpl, err = p.getTpl()
	if err != nil {
		return "", err
	}
	return p.fillTpl(tpl), nil
}

// getStatus
func (p *Plugin) getStatus() string {
	if p.Drone.Build.Status == "success" {
		if p.Tpl.Build.Status.Success != "" {
			return p.Tpl.Build.Status.Success
		}

		return p.Drone.Build.Status
	}

	if p.Tpl.Build.Status.Failure != "" {
		return p.Tpl.Build.Status.Failure
	}

	return p.Drone.Build.Status
}

// get emoticon
func (p *Plugin) getEmoticon() string {
	emoticons := make(map[string]string)
	emoticons["success"] = ":)"
	emoticons["failure"] = ":("

	emoticon, ok := emoticons[p.Drone.Build.Status]
	if ok {
		return emoticon
	}

	return ":("
}

// get picture url
func (p *Plugin) getPicURL() string {
	pics := make(map[string]string)
	//  success picture url
	pics["success"] = "https://wx1.sinaimg.cn/large/006tNc79gy1fz05g5a7utj30he0bfjry.jpg"
	if p.Custom.Pic.SuccessPicURL != "" {
		pics["success"] = p.Custom.Pic.SuccessPicURL
	}
	//  failure picture url
	pics["failure"] = "https://wx1.sinaimg.cn/large/006tNc79gy1fz0b4fghpnj30hd0bdmxn.jpg"
	if p.Custom.Pic.FailurePicURL != "" {
		pics["failure"] = p.Custom.Pic.FailurePicURL
	}

	picURL, ok := pics[p.Drone.Build.Status]
	if ok {
		return picURL
	}

	return ""
}

// get color for message title
func (p *Plugin) getColor() string {
	colors := make(map[string]string)
	//  success color
	colors["success"] = "#008000"
	if p.Custom.Color.SuccessColor != "" {
		colors["success"] = "#" + p.Custom.Color.SuccessColor
	}

	//  failure color
	colors["failure"] = "#FF0000"
	if p.Custom.Color.FailureColor != "" {
		colors["failure"] = "#" + p.Custom.Color.FailureColor
	}

	color, ok := colors[p.Drone.Build.Status]
	if ok {
		return color
	}

	return ""
}

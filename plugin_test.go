package main

import (
	"testing"
)

func TestPlugin(t *testing.T) {
	p := Plugin{}
	err := p.Exec()
	if nil == err {
		t.Error("access token empty error should be catch!")
	}

	p.Config.AccessToken = "example-access-token"
	err = p.Exec()
	if nil == err {
		t.Error("commit sha length error should be catch!")
	}

	p.Drone.Commit.Sha = "53729847dfksj"
	err = p.Exec()
	if nil == err {
		t.Error("not support message type error should be catch!")
	}

	p.Config.MsgType = "text"
	err = p.Exec()
	if nil == err {
		t.Error("access token invalid error should be catch!")
	}

	p.Config.MsgType = "link"
	err = p.Exec()
	if nil == err {
		t.Error("access token invalid error should be catch!")
	}

	p.Extra.Color.WithColor = true
	p.Extra.Color.FailureColor = "#555555"
	p.Extra.Color.SuccessColor = "#222222"
	p.Extra.Pic.WithPic = true
	p.Extra.Pic.FailurePicURL = "https://www.baidu.com"
	p.Extra.Pic.SuccessPicURL = "https://www.baidu.com"
	p.Extra.LinkSha = true
	// p.Drone.Build.Status = "failure"
	p.Config.MsgType = "markdown"
	err = p.Exec()
	if nil == err {
		t.Error("access token invalid error should be catch!")
	}

	p.Drone.Build.Status = "failure"
	err = p.Exec()
	if nil == err {
		t.Error("access token invalid error should be catch!")
	}

	t.Log("plugin testing finished")
}

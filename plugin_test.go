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
	p.Custom.Tpl = "tpls/markdown.tpl"
	err = p.Exec()
	if nil == err {
		t.Error("not support message type error should be catch!")
	}

	p.Config.MsgType = "link"
	err = p.Exec()
	if nil == err {
		t.Error("access token invalid error should be catch!")
	}

	p.Custom.Tpl = "https://aaa.com"
	p.Config.MsgType = "text"
	err = p.Exec()
	if nil == err {
		t.Error("access token invalid error should be catch!")
	}


	p.Custom.Tpl = ""
	p.Config.MsgType = "link"
	err = p.Exec()
	if nil == err {
		t.Error("access token invalid error should be catch!")
	}

	p.Custom.Color.FailureColor = "#555555"
	p.Custom.Color.SuccessColor = "#222222"
	p.Custom.Pic.FailurePicURL = "https://www.baidu.com"
	p.Custom.Pic.SuccessPicURL = "https://www.baidu.com"
	p.Config.MsgType = "markdown"
	p.Custom.Tpl = "tpls/markdown.tpl"
	err = p.Exec()
	if nil == err {
		t.Error("access token invalid error should be catch!")
	}

	p.Custom.Tpl = "https://gist.githubusercontent.com/lddsb/87065e73678dcf56cd222a3c2f1f32b0/raw/fce9fb28b2c8c768eb93df5598beee8c98cba610/md.tpl"
	p.Drone.Build.Status = "failure"
	err = p.Exec()
	if nil == err {
		t.Error("access token invalid error should be catch!")
	}

	t.Log("plugin testing finished")
}

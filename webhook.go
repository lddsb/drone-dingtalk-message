package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// LinkMsg `link message struct`
type LinkMsg struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL"`
}

// ActionCard `action card message struct`
type ActionCard struct {
	Text           string `json:"text"`
	Title          string `json:"title"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
	BtnOrientation string `json:"btnOrientation"`
	HideAvatar     string `json:"hideAvatar"` //  robot message avatar
	Buttons        []struct {
		Title     string `json:"title"`
		ActionURL string `json:"actionURL"`
	} `json:"btns"`
}

// PayLoad payload
type PayLoad struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Link struct {
		Title      string `json:"title"`
		Text       string `json:"text"`
		PicUrl     string `json:"picUrl"`
		MessageUrl string `json:"messageUrl"`
	} `json:"link"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	ActionCard ActionCard `json:"actionCard"`
	FeedCard   struct {
		Links []LinkMsg `json:"links"`
	} `json:"feedCard"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

// WebHook `web hook base config`
type WebHook struct {
	AccessToken string `json:"accessToken"`
}

// NewWebHook `new a webhook`
func NewWebHook(accessToken string) *WebHook {
	return &WebHook{AccessToken: accessToken}
}

// Response `dingtalk webhook response struct`
type Response struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

var baseApi = "https://oapi.dingtalk.com/robot/send?access_token="
var reg = `^1([38][0-9]|14[57]|5[^4])\d{8}$`
var regx = regexp.MustCompile(reg)

//  real send request to api
func (w *WebHook) sendPayload(payload *PayLoad) error {
	//  get config
	bs, err := json.Marshal(payload)
	if nil != err {
		return err
	}
	//  request api
	resp, err := http.Post(baseApi+w.AccessToken, "application/json", bytes.NewReader(bs))
	if nil != err {
		return err
	}
	//  read response body
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return err
	}
	//  api unusual
	if 200 != resp.StatusCode {
		return fmt.Errorf("%d: %s", resp.StatusCode, string(body))
	}

	var result Response
	//  json decode
	err = json.Unmarshal(body, &result)
	if nil != err {
		return err
	}
	if 0 != result.ErrorCode {
		return fmt.Errorf("%d: %s", result.ErrorCode, result.ErrorMessage)
	}

	return nil
}

// SendTextMsg `send a text message`
func (w *WebHook) SendTextMsg(content string, isAtAll bool, mobiles ...string) error {
	//  send request
	return w.sendPayload(&PayLoad{
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
		At: struct {
			AtMobiles []string `json:"atMobiles"`
			IsAtAll   bool     `json:"isAtAll"`
		}{
			AtMobiles: mobiles,
			IsAtAll:   isAtAll,
		},
	})
}

// SendLinkMsg `send a link message`
func (w *WebHook) SendLinkMsg(title, content, picURL, msgURL string) error {
	return w.sendPayload(&PayLoad{
		MsgType: "link",
		Link: struct {
			Title      string `json:"title"`
			Text       string `json:"text"`
			PicUrl     string `json:"picUrl"`
			MessageUrl string `json:"messageUrl"`
		}{
			Title:      title,
			Text:       content,
			PicUrl:     picURL,
			MessageUrl: msgURL,
		},
	})
}

// SendMarkdownMsg `send a markdown msg`
func (w *WebHook) SendMarkdownMsg(title, content string, isAtAll bool, mobiles ...string) error {
	firstLine := false
	for _, mobile := range mobiles {
		if regx.MatchString(mobile) {
			if false == firstLine {
				content += "#####"
			}
			content += " @" + mobile
			firstLine = true
		}
	}
	//  send request
	return w.sendPayload(&PayLoad{
		MsgType: "markdown",
		Markdown: struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		}{
			Title: title,
			Text:  content,
		},
		At: struct {
			AtMobiles []string `json:"atMobiles"`
			IsAtAll   bool     `json:"isAtAll"`
		}{
			AtMobiles: mobiles,
			IsAtAll:   isAtAll,
		},
	})
}

// SendActionCardMsg `send single action card message`
func (w *WebHook) SendActionCardMsg(title, content string, linkTitles, linkUrls []string, hideAvatar, btnOrientation bool) error {
	//  validation is empty
	if 0 == len(linkTitles) || 0 == len(linkUrls) {
		return errors.New("links or titles is empty！")
	}
	//  validation is equal
	if len(linkUrls) != len(linkTitles) {
		return errors.New("links length and titles length is not equal！")
	}
	//  hide robot avatar
	var strHideAvatar = "0"
	if hideAvatar {
		strHideAvatar = "1"
	}
	//  button sort
	var strBtnOrientation = "0"
	if btnOrientation {
		strBtnOrientation = "1"
	}
	//  button struct
	var buttons []struct {
		Title     string `json:"title"`
		ActionURL string `json:"actionURL"`
	}
	//  inject to button
	for i := 0; i < len(linkTitles); i++ {
		buttons = append(buttons, struct {
			Title     string `json:"title"`
			ActionURL string `json:"actionURL"`
		}{
			Title:     linkTitles[i],
			ActionURL: linkUrls[i],
		})
	}
	//  send request
	return w.sendPayload(&PayLoad{
		MsgType: "actionCard",
		ActionCard: ActionCard{
			Title:          title,
			Text:           content,
			HideAvatar:     strHideAvatar,
			BtnOrientation: strBtnOrientation,
			Buttons:        buttons,
		},
	})
}

// SendLinkCardMsg `send link card message`
func (w *WebHook) SendLinkCardMsg(messages []LinkMsg) error {
	return w.sendPayload(&PayLoad{
		MsgType: "feedCard",
		FeedCard: struct {
			Links []LinkMsg `json:"links"`
		}{
			Links: messages,
		},
	})
}

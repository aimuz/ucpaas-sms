package ucpaassmsclient

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"bytes"
)

type Config struct {
	SID   string `json:"sid"`
	Token string `json:"token"`
	AppID string `json:"appid"`
}
type SendResponse struct {
	Code       string `json:"code"`
	Msg        string `json:"msg"`
	Count      string `json:"count"`
	CreateDate string `json:"create_date"`
	UID        string `json:"uid"`
	Smsid      string `json:"smsid"`
	Mobile     string `json:"mobile"`
}

func (m *Config) Send(templateid, param, mobile, uid string) (*SendResponse, error) {
	params := map[string]string{
		"sid":        m.SID,
		"token":      m.Token,
		"appid":      m.AppID,
		"templateid": templateid,
		"param":      param,
		"mobile":     mobile,
		"uid":        uid,
	}
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(SendSMSURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	sendResp := new(SendResponse)
	err = json.Unmarshal(body, sendResp)
	if err != nil {
		return nil, err
	}
	return sendResp, nil
}

type SendBatchResponse struct {
	CountSum   string   `json:"count_sum"`
	CreateDate string   `json:"create_date"`
	UID        string   `json:"uid"`
	Report     []Report `json:"report"`
	Code       string   `json:"code"`
	Msg        string   `json:"msg"`
}
type Report struct {
	Count  string `json:"count"`
	Code   string `json:"code"`
	Msg    string `json:"msg"`
	Smsid  string `json:"smsid"`
	Mobile string `json:"mobile"`
}

func (m *Config) SendBatch(templateid, param, mobile, uid string) (interface{}, error) {
	params := map[string]string{
		"sid":        m.SID,
		"token":      m.Token,
		"appid":      m.AppID,
		"templateid": templateid,
		"param":      param,
		"mobile":     mobile,
		"uid":        uid,
	}
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(SendSMSBatchURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	sendBatchResp := new(SendBatchResponse)
	err = json.Unmarshal(body, sendBatchResp)
	if err != nil {
		return nil, err
	}
	return sendBatchResp, nil
}

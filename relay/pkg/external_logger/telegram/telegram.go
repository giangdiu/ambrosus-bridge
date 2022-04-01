package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type externalLogger struct {
	Token      string
	ChatId     int
	HttpClient *http.Client
	Prefix     string
}

func NewExternalLogger(token string, chatId int, prefix string, httpClient *http.Client) *externalLogger {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &externalLogger{
		Token:      token,
		ChatId:     chatId,
		HttpClient: httpClient,
		Prefix:     prefix,
	}
}

type request struct {
	ChatId    int    `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

type response struct {
	Ok               bool   `json:"ok"`
	ErrorDescription string `json:"description"` // if Ok is false
}

func (t *externalLogger) LogError(msg string) (returningError error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.Token)
	body := &request{
		ChatId:    t.ChatId,
		Text:      fmt.Sprintf("%s <b>We got an unexpected error:</b>\n%s", t.Prefix, msg),
		ParseMode: "html",
	}

	payloadBuf := new(bytes.Buffer)
	if err := json.NewEncoder(payloadBuf).Encode(body); err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", payloadBuf)
	if err != nil {
		return err
	}
	defer func() {
		err = resp.Body.Close()
	}()

	respData := new(response)
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return err
	}
	if !respData.Ok {
		return fmt.Errorf(respData.ErrorDescription)
	}
	return err
}
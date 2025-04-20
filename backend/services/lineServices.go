package services

import (
	"backend/config"
	"backend/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Line 登入回調
func HandleLineCallback(code string) (*models.LineLoginResponse, error) {
	tokenURL := "https://api.line.me/oauth2/v2.1/token"
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", config.RedirectURI)
	data.Set("client_id", config.LineClientID)
	data.Set("client_secret", config.LineClientSecret)

	resp, err := http.Post(tokenURL, "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("token response:", string(body))

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	accessToken, ok := result["access_token"].(string)
	if !ok {
		return nil, errors.New("access_token missing")
	}

	req, _ := http.NewRequest("GET", "https://api.line.me/v2/profile", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	profileResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer profileResp.Body.Close()

	profileBody, _ := io.ReadAll(profileResp.Body)
	fmt.Println("👤 profile response:", string(profileBody))

	var profile models.LineProfile
	if err := json.Unmarshal(profileBody, &profile); err != nil {
		return nil, err
	}

	return &models.LineLoginResponse{
		Profile:     profile,
		AccessToken: accessToken,
	}, nil
}
// 發送消息到 Line
func SendLineMessage(userID, message string) error {
	url := "https://api.line.me/v2/bot/message/push"
	channelToken := config.LineChannelAccessToken

	body := map[string]interface{}{
		"to": userID,
		"messages": []map[string]string{
			{
				"type": "text",
				"text": message,
			},
		},
	}

	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization",channelToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("LINE API 回傳錯誤狀態碼：%d", resp.StatusCode)
	}

	return nil
}


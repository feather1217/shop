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


func HandleLineCallback(code string) (*models.LineProfile, error) {
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
	fmt.Println("🔑 token response:", string(body)) // 印出 access token 回應

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
	fmt.Println("👤 profile response:", string(profileBody)) // 印出使用者資料回應

	var profile models.LineProfile
	if err := json.Unmarshal(profileBody, &profile); err != nil {
		return nil, err
	}

	return &profile, nil
}

package services

import (
	"backend/config"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type LineProfile struct {
	UserID      string `json:"userId"`
	DisplayName string `json:"displayName"`
	PictureURL  string `json:"pictureUrl"`
}

func HandleLineCallback(code string) (*LineProfile, error) {
	// 換 access token
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
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	accessToken, ok := result["access_token"].(string)
	if !ok {
		return nil, errors.New("access_token missing")
	}

	// 拿 profile
	req, _ := http.NewRequest("GET", "https://api.line.me/v2/profile", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	profileResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer profileResp.Body.Close()

	profileBody, _ := io.ReadAll(profileResp.Body)
	var profile LineProfile
	if err := json.Unmarshal(profileBody, &profile); err != nil {
		return nil, err
	}

	return &profile, nil
}

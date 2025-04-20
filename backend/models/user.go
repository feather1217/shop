// models/line_profile.go
package models

type LineProfile struct {
    UserID      string `json:"userId"`
    DisplayName string `json:"displayName"`
    PictureURL  string `json:"pictureUrl"`
}

type LineLoginResponse struct {
    Profile     LineProfile `json:"profile"`
    AccessToken string      `json:"accessToken"`
}

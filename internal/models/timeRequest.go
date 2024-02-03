package models

type TimeRequest struct {
	AccessToken int    `json:"accessToken"`
	Signature   string `json:"signature"`
}

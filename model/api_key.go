package model

type ApiKey struct {
	ID        int    `json:"id"`
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
}

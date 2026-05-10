package models

type Configuration struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	Version    string            `json:"version"`
	Parameters map[string]string `json:"parameters"`
}

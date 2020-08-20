package model

type RequestType struct {
	Search string   `json:"search"`
	Sites  []string `json:"sites"`
}

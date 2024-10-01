package models

type ProjectInfo struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Version     Version `json:"version"`
}

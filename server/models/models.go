package models

import "time"

type ExecutionRequest struct {
	Language string `json:"language" binding:"required"`
	Code     string `json:"code"     binding:"required"`
	Stdin    string `json:"stdin"`
}

type RunResult struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Code   int    `json:"code"`
	Signal string `json:"signal"`
}

type ExecutionResponse struct {
	Run      RunResult `json:"run"`
	Language string    `json:"language"`
	Version  string    `json:"version"`
}

type Snippet struct {
	ID        string    `json:"id"`
	Language  string    `json:"language"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}

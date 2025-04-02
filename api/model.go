package api

type Result struct {
	Ip        string `json:"ip"`
	UUID      string `json:"uuid"`
	UserAgent string `json:"user_agent"`
}

type UserAgent struct {
	UserAgent string `json:"user-agent"`
}

type UUID struct {
	UUID string `json:"uuid"`
}

type IP struct {
	IP string `json:"origin"`
}

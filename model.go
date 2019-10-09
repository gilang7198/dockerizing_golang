package main

// Versions struct
type Versions struct {
	ID       int    `form:"id" json:"id"`
	NameApps string `form:"name_apps" json:"name_apps"`
	Version  string `form:"version" json:"version"`
}

// Response struct
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Versions
}

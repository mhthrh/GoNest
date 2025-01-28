package db

type DB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Driver   string `json:"driver"`
	MinCount int    `json:"minCount"`
	MaxCount int    `json:"maxCount"`
}

package model

type MetaData struct {
	AppName    string `json:"appName"`
	Version    string `json:"version"`
	IsTest     bool   `json:"isTest"`
	ExpireDate string `json:"expireDate"`
}
type Secret struct {
	SecretKey     string `json:"secretKey"`
	TokenDuration string `json:"tokenDuration"`
}
type DB struct {
	Host     string  `json:"host"`
	Port     int     `json:"port"`
	UserName string  `json:"user"`
	Password string  `json:"password"`
	DbName   string  `json:"dbName"`
	Driver   string  `json:"driver"`
	SSLMode  SslType `json:"sslMode"`
}

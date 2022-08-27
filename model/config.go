package model

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Title string
	QrcodeLink string
	Description string
}

var ConfigData Config
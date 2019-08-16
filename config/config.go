package config

var Conf = new(Config)

type Config struct {
	ServerHost string `json:"server_host"`
	Docker     string `json:"docker"`
}

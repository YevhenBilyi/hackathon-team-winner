package db

type Config struct {
	Connection       string `default:"mongodb://localhost:31060"`
	Name             string `default:"noonaNordar"`
	DirectConnection bool   `default:"true"`
}

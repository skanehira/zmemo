package main

// Config : config info
type Config struct {
	AppName string `default:"zmemo"`
	Port    string `default:"80"`
	DB      struct {
		Name     string `default:"zmemo"`
		User     string `default:"zmemo"`
		Password string `default:"zmemo"`
		Port     string `default:"3306"`
		Host     string `default:"localhost"`
	}
}

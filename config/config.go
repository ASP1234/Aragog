package config

// structure for decoding config.yml
type Init struct {
	OS struct {
		MaxProcessors int `yaml:"maxProcessors"`
	} `yaml:"os"`
	Processor struct {
		MaxRetryAttempts int `yaml:"maxRetryAttempts"`
		MaxRoutines      int `yaml:"maxRoutines"`
	} `yaml:"processor"`
	Producer struct {
		SeedUrl string `yaml:"seedUrl"`
	} `yaml:"producer"`
}

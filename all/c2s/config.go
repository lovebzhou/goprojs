package c2s

type TransportConfig struct {
	Type string `yaml:"type"`
	Port int    `yaml:"port"`
}

type Config struct {
	ID        string
	Transport TransportConfig
}

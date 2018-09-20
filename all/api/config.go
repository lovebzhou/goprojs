package api

type Config struct {
	Port int `yaml:"port"`
}

// func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
// 	log.Println("api.Config.UnmarshalYAML")

// 	return nil
// }

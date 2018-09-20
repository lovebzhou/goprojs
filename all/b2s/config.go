package b2s

type Config struct {
	Port int `yaml:"port"`
}

//
type proxyConfig struct {
	Port int `yaml:"port"`
}

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	p := proxyConfig{}
	if err := unmarshal(&p); err != nil {
		return err
	}

	c.Port = p.Port

	return nil
}

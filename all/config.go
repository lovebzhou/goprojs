package main

import (
	"bytes"
	"io/ioutil"

	"github.com/lovebzhou/goprojs/all/api"
	"github.com/lovebzhou/goprojs/all/b2s"
	"github.com/lovebzhou/goprojs/all/c2s"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	API api.Config
	B2S b2s.Config
	C2S []c2s.Config
}

// FromFile loads default global configuration from
// a specified file.
func (cfg *Config) FromFile(configFile string) error {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, cfg)
}

// FromBuffer loads default global configuration from
// a specified byte buffer.
func (cfg *Config) FromBuffer(buf *bytes.Buffer) error {
	return yaml.Unmarshal(buf.Bytes(), cfg)
}

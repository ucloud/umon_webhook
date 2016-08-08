package utils

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Raw  []byte
	data map[string]interface{}
}

var g_Cfg *Config

func GetGlobalConf() *Config {
	return g_Cfg
}

func SetGlobalConf(cfg *Config) {
	if cfg != nil {
		g_Cfg = cfg
	}
}

func NewConfig(path string) (*Config, error) {
	cfg := &Config{
		data: make(map[string]interface{}),
	}

	err := cfg.parse(path)
	if err != nil {
		cfg = nil
	}

	return cfg, err
}

func (cfg *Config) parse(path string) error {
	jsonRaw, err := ioutil.ReadFile(path)
	if err == nil {
		cfg.Raw = jsonRaw
		err = json.Unmarshal(jsonRaw, &cfg.data)

	}

	return err
}

func (cfg *Config) GetFloat(key string) float64 {
	x, ok := cfg.data[key]
	if !ok {
		return -1
	}

	return x.(float64)
}

func (cfg *Config) GetInt(key string) int64 {
	x, ok := cfg.data[key]
	if !ok {
		return -1
	}

	return int64(x.(float64))
}

func (cfg *Config) GetString(key string) string {
	x, ok := cfg.data[key]
	if !ok {
		return ""
	}

	return x.(string)
}

func (cfg *Config) GetBool(key string) bool {
	x, ok := cfg.data[key]
	if !ok {
		return false
	}

	return x.(bool)
}

func (cfg *Config) GetArray(key string) []interface{} {
	x, ok := cfg.data[key]
	if !ok {
		return []interface{}(nil)
	}

	return x.([]interface{})
}

func (cfg *Config) Set(key string, val interface{}) {
	cfg.data[key] = val
}

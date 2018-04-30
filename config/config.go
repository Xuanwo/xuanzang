package config

import (
	"errors"
	"io/ioutil"

	"github.com/pengsrc/go-shared/convert"
	"gopkg.in/yaml.v2"

	"github.com/Xuanwo/xuanzang/constants"
)

// Config contains all config needed by xuanzang.
type Config struct {
	Host      *string `yaml:"host"`
	Port      *int    `yaml:"port"`
	DBPath    *string `yaml:"db_path"`
	IndexPath *string `yaml:"index_path"`

	Dictionary *string `yaml:"dictionary"`
	StopTokens *string `yaml:"stop_tokens"`

	Source *Source `yaml:"source"`

	Logger *Logger `yaml:"logger"`
}

// Source will specify a source for xuanzang.
type Source struct {
	Type     string `yaml:"type"`
	URL      string `yaml:"url"`
	Duration int    `yaml:"duration"`
}

// New will create a new config instance.
func New() *Config {
	return &Config{}
}

// LoadFromFilePath loads configuration from a specified local path.
// It returns error if file not found or yaml decode failed.
func (c *Config) LoadFromFilePath(filePath string) error {
	cYAML, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	return c.LoadFromContent(cYAML)
}

// LoadFromContent loads configuration from a given bytes.
// It returns error if yaml decode failed.
func (c *Config) LoadFromContent(content []byte) error {
	return yaml.Unmarshal(content, c)
}

// Check will check if the config is valid.
func (c *Config) Check() error {
	if c.Host == nil {
		c.Host = convert.String("localhost")
	}
	if c.Port == nil {
		c.Port = convert.Int(8080)
	}
	if c.DBPath == nil {
		return errors.New("db path is empty")
	}

	if c.IndexPath == nil {
		return errors.New("index path is empty")
	}

	if c.Source == nil {
		return errors.New("source is empty")
	}

	if c.Dictionary == nil {
		return errors.New("dictionary is empty")
	}

	if c.StopTokens == nil {
		return errors.New("stop tokens is empty")
	}

	// Check logger.
	if c.Logger == nil {
		c.Logger = &Logger{
			Level:  constants.DefaultLogLevel,
			Output: constants.DefaultLogOutput,
		}
	}

	return nil
}

// TODO: Add check for source.

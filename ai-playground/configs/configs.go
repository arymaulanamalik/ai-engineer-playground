package configs

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port           int    `envconfig:"port" default:"9001"`
	Host           string `envconfig:"host" default:"localhost:9001"`
	BaseURL        string `envconfig:"base_url" default:"http://localhost:9001"`
	ServiceName    string `envconfig:"service_name" default:"playground"`
	ServiceVersion string `envconfig:"service_version" default:"1.0.0"`
	Environment    string `envconfig:"environment" default:"staging"`

	LogLevel string `envconfig:"log_level" default:"DEBUG"`

	LLM LLM `envconfig:"llm"`
}

type LLM struct {
	ClaudeURL       string `envconfig:"claude_url" default:"https://api.anthropic.com/v1"`
	ClaudeModel     string `envconfig:"claude_model" default:"claude-2"`
	ClaudeAPIKey    string `envconfig:"claude_api_key" default:""`
	ClaudeMaxTokens int    `envconfig:"claude_max_tokens" default:"2048"`
	ClaudeTimeout   int    `envconfig:"claude_timeout" default:"30"`
}

func NewConfig() *Config {
	var conf Config
	err := envconfig.Process("PLAYGROUND", &conf)
	if err != nil {
		log.Fatalf("fail to proceed the config: %v", err)
	}
	return &conf
}

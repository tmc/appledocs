package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// Config represents the configuration for framework binding generation.
type Config struct {
	Frameworks map[string]FrameworkConfig `yaml:"frameworks"`
}

// FrameworkConfig represents configuration for a specific framework.
type FrameworkConfig struct {
	ExcludeTestExamples []string `yaml:"exclude_test_examples"`
	Deprecated          bool     `yaml:"deprecated"`
	DeprecationReason   string   `yaml:"deprecation_reason"`
}

var config *Config

// loadConfig loads configuration from embedded config.yaml file.
func loadConfig() error {
	data, err := embeddedFS.ReadFile("config.yaml")
	if err != nil {
		// Config file is optional, return empty config
		config = &Config{
			Frameworks: make(map[string]FrameworkConfig),
		}
		return nil
	}

	config = &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return fmt.Errorf("failed to parse config.yaml: %w", err)
	}

	if config.Frameworks == nil {
		config.Frameworks = make(map[string]FrameworkConfig)
	}

	return nil
}

// isFrameworkDeprecated checks if a framework is marked as deprecated.
func isFrameworkDeprecated(framework string) bool {
	if config == nil {
		return false
	}

	frameworkConfig, ok := config.Frameworks[framework]
	if !ok {
		return false
	}

	return frameworkConfig.Deprecated
}

// shouldExcludeTestExample checks if a class should be excluded from test example generation.
func shouldExcludeTestExample(framework, className string) bool {
	if config == nil {
		return false
	}

	frameworkConfig, ok := config.Frameworks[framework]
	if !ok {
		return false
	}

	for _, excludedClass := range frameworkConfig.ExcludeTestExamples {
		if excludedClass == className {
			return true
		}
	}

	return false
}

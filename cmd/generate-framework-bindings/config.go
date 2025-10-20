package main

import (
	"fmt"
	"regexp"

	"gopkg.in/yaml.v3"
)

// Config represents the configuration for framework binding generation.
type Config struct {
	Frameworks map[string]FrameworkConfig `yaml:"frameworks"`
}

// FrameworkConfig represents configuration for a specific framework.
type FrameworkConfig struct {
	ExcludeTestExamples        []string `yaml:"exclude_test_examples"`         // Class names to exclude (deprecated, use exclude_test_patterns)
	ExcludeTestPatterns        []string `yaml:"exclude_test_patterns"`         // Regex patterns for test method names to exclude
	Deprecated                 bool     `yaml:"deprecated"`
	DeprecationReason          string   `yaml:"deprecation_reason"`
	excludeTestPatternsCompiled []*regexp.Regexp // Compiled regex patterns (not in YAML)
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

	// Compile regex patterns for each framework
	for fwName, fwConfig := range config.Frameworks {
		for _, pattern := range fwConfig.ExcludeTestPatterns {
			re, err := regexp.Compile(pattern)
			if err != nil {
				return fmt.Errorf("failed to compile regex pattern %q for framework %s: %w", pattern, fwName, err)
			}
			fwConfig.excludeTestPatternsCompiled = append(fwConfig.excludeTestPatternsCompiled, re)
		}
		// Update the config with compiled patterns
		config.Frameworks[fwName] = fwConfig
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
// This function is deprecated in favor of shouldExcludeTestMethod which allows more precise filtering.
func shouldExcludeTestExample(framework, className string) bool {
	if config == nil {
		return false
	}

	frameworkConfig, ok := config.Frameworks[framework]
	if !ok {
		return false
	}

	// Legacy class-based exclusion
	for _, excludedClass := range frameworkConfig.ExcludeTestExamples {
		if excludedClass == className {
			return true
		}
	}

	return false
}

// shouldExcludeTestMethod checks if a specific test method should be excluded from generation.
// testMethodName should be the Go test function name (e.g., "ExampleNewCKContainerWithIdentifier").
// This uses regex patterns from the config for flexible filtering.
func shouldExcludeTestMethod(framework, testMethodName string) bool {
	if config == nil {
		return false
	}

	frameworkConfig, ok := config.Frameworks[framework]
	if !ok {
		return false
	}

	// Check against compiled regex patterns
	for _, pattern := range frameworkConfig.excludeTestPatternsCompiled {
		if pattern.MatchString(testMethodName) {
			return true
		}
	}

	return false
}

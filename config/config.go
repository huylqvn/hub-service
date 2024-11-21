package config

import (
	"flag"
	"fmt"
	"os"

	"hub-service/util"

	"gopkg.in/yaml.v3"
)

// Config represents the composition of yml settings.
type Config struct {
	Database struct {
		Dialect   string `default:"postgres"`
		Host      string `default:"localhost"`
		Port      string `default:"5432"`
		Dbname    string `default:"postgres"`
		Username  string `default:"postgres"`
		Password  string `default:"postgres"`
		Migration bool   `default:"false"`
	}
	Redis struct {
		Enabled            bool   `default:"false"`
		ConnectionPoolSize int    `yaml:"connection_pool_size" default:"10"`
		Host               string `default:"localhost"`
		Port               string `default:"6379"`
	}
	Extension struct {
		MasterGenerator bool `yaml:"master_generator" default:"false"`
		CorsEnabled     bool `yaml:"cors_enabled" default:"false"`
		SecurityEnabled bool `yaml:"security_enabled" default:"false"`
	}
	Log struct {
		RequestLogFormat string `yaml:"request_log_format" default:"${remote_ip} ${User_name} ${uri} ${method} ${status}"`
	}
	Swagger struct {
		Enabled bool `default:"false"`
		Path    string
	}
	Security struct {
		AuthPath    []string `yaml:"auth_path"`
		ExculdePath []string `yaml:"exclude_path"`
		UserPath    []string `yaml:"user_path"`
		AdminPath   []string `yaml:"admin_path"`
	}
}

const (
	// DEV represents development environment
	DEV = "develop"
	// PRD represents production environment
	PRD = "production"
	// DOC represents docker container
	DOC = "docker"
)

// LoadAppConfig reads the settings written to the yml file
func LoadAppConfig() (*Config, string) {
	var env *string
	if value := os.Getenv("WEB_APP_ENV"); value != "" {
		env = &value
	} else {
		env = flag.String("env", "develop", "To switch configurations.")
		flag.Parse()
	}
	file, err := os.ReadFile(fmt.Sprintf(AppConfigPath, *env))
	if err != nil {
		fmt.Printf("Failed to read application.%s.yml: %s", *env, err)
		os.Exit(ErrExitStatus)
	}

	config := &Config{}
	if err := yaml.Unmarshal(file, config); err != nil {
		fmt.Printf("Failed to read application.%s.yml: %s", *env, err)
		os.Exit(ErrExitStatus)
	}

	return config, *env
}

// LoadMessagesConfig loads the messages.properties.
func LoadMessagesConfig() map[string]string {
	messages := util.ReadPropertiesFile(MessagesConfigPath)
	if messages == nil {
		fmt.Printf("Failed to load the messages.properties.")
		os.Exit(ErrExitStatus)
	}
	return messages
}

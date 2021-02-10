package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

// ConfigFile struct to `config.yaml`
type ConfigFile struct {
	Vars  map[string]string `mapstructure:"vars"`
	Rules []Rule            `mapstructure:"rule"`
}

// Config clink configs
type Config struct {
	WorkDIR    string // path to `config.yaml`
	DryRun     bool
	BackupPath string

	*ConfigFile
}

// AppConfig global configs
var AppConfig *Config

// SetupConfig initialize global configs
func SetupConfig(dryRun bool, configPath string) {
	absConfigPath, err := filepath.Abs(configPath)
	if err != nil {
		log.Fatalf("failed to resolve config path %v: %v", configPath, err)
	}
	fmt.Println(configPath, absConfigPath)

	workDIR := filepath.Dir(absConfigPath)
	backupPath := confirmBackupPath()
	configFile := readConfigFile(absConfigPath)

	cfg := &Config{
		DryRun:     dryRun,
		WorkDIR:    workDIR,
		BackupPath: backupPath,
		ConfigFile: configFile,
	}

	AppConfig = confirmConfig(cfg)
}

// unmarshall yaml config file content to struct
func readConfigFile(absPath string) (c *ConfigFile) {
	viper.SetConfigType("yaml")

	f, err := os.Open(absPath)
	if err != nil {
		log.Fatalf("failed to open config file %v: %v", absPath, err)
	}
	defer f.Close()

	if err = viper.ReadConfig(f); err != nil {
		log.Fatalf("failed to read config file %v: %v", absPath, err)
	}

	var configFile ConfigFile
	if err = viper.Unmarshal(&configFile); err != nil {
		log.Fatalf("wrong config file format %v: %v", absPath, err)
	}

	return &configFile
}

// ask user to confirm original config files (if exists) backup path
func confirmBackupPath() string {
	defaultPath, err := filepath.Abs("~/.clink")
	if err != nil {
		defaultPath = ""
	}

	p := promptui.Prompt{
		Label:   "Please specify your backup path",
		Default: defaultPath,
	}

	result, err := p.Run()
	fmt.Println(result)

	// TODO confirm backup path

	return ""
}

// ask user to confirm variables defined in config file
// and render config file with absolute pathes
func confirmConfig(cfg *Config) *Config {
	// TODO render variables
	return nil
}

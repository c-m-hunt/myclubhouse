package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// Config Struct holding the config object
type Config struct {
	AccessToken string
	SubDomain   string
}

func printTitle(title string) {
	fmt.Println(strings.Repeat("-", len(title)+6))
	fmt.Printf("-- %v --\n", title)
	fmt.Println(strings.Repeat("-", len(title)+6))
}

func getConfigLocation() string {
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Cannot get user home directory to save config")
	}
	configDir := path.Join(homePath, ".config", "myclubhouse")
	err = os.MkdirAll(configDir, 0644)
	if err != nil {
		log.Fatal("Could not create config directory")
	}
	return path.Join(configDir, "config")
}

// SetConfig Set the config data and save it to file
func SetConfig(accessToken string, subDomain string) error {
	c := Config{
		AccessToken: accessToken,
		SubDomain:   subDomain,
	}
	return c.Save(getConfigLocation())
}

// Save Saves the config to the path
func (c Config) Save(path string) error {
	j, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, j, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Config file saved to %v", path)
	return nil
}

// EnsureConfig Make sure the config data can be loaded otherwise panic
func (c *Config) EnsureConfig() {
	b, err := ioutil.ReadFile(getConfigLocation())
	if err != nil {
		log.Fatal("Cannot load config file")
	}
	err = json.Unmarshal(b, &c)
	if err != nil {
		log.Fatal("Config file not valid")
	}
}

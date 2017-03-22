package config

// Package lib contains a common library functionalities
import (
	"fmt"
	"os"
	"path"
	"strings"

	c "github.com/revel/config"
)

var (
	filename   = `app.conf`
	production = `production`
)

// Config struct provide access to configuration files
type Config struct {
	Pwd          string
	Filename     string
	IsProduction bool
	base         *c.Config
}

// NewConfig creates a new configuration struct
func NewConfig(cfgfile ...string) (config *Config, err error) {
	if len(cfgfile) > 0 {
		filename = cfgfile[0]
	}
	config = &Config{Filename: filename}
	if config.Pwd, err = os.Getwd(); err != nil {
		fmt.Errorf("| Error | %v \n", err)
		panic(err)
	}

	var file = config.File()
	//	fmt.Printf("App | Config will be loaded from %v \n", file)
	if config.base, err = c.ReadDefault(file); err != nil {
		fmt.Errorf("| Error | %v \n", err)
		panic(err)
	}
	//	fmt.Println("App | Config loaded successfully! \n")
	config.IsProduction = strings.EqualFold(config.Default("env"), production)
	return
}

// File gets  absolute configuration path
func (c *Config) File() (file string) {
	return path.Join(c.Pwd, c.Filename)
}

// Database gets any database configuration property
func (c *Config) Database(property string) (result string) {
	result, _ = c.base.String("database", property)
	return
}

// Default gets any configuration property from default section
func (c *Config) Default(property string) (result string) {
	result, _ = c.base.String("default", property)
	return
}

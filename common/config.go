package common

import (
	"encoding/xml"
	"io/ioutil"
	"log"

	"git.garena.com/duanzy/motto/motto"
)

// Cfg - get application config
func Cfg(app motto.Application) *Configuration {
	return app.Settings().(*Configuration)
}

// NewConfiguration - create a new configuration object
func NewConfiguration(filepath string) *Configuration {
	return &Configuration{
		filepath: filepath,
	}
}

// Configuration - app configuration that implements Motto's config interface
type Configuration struct {
	Settings
	filepath string
}

// Motto - return the Motto settings
func (c *Configuration) Motto() *motto.Settings {
	return c.Settings.Motto
}

// Load - load configuration
func (c *Configuration) Load() (err error) {
	var (
		content []byte
	)
	if content, err = ioutil.ReadFile(c.filepath); err != nil {
		return err
	}

	settings := Settings{}
	if err = xml.Unmarshal(content, &settings); err != nil {
		log.Fatalf("Failed to unmarshal application settings. err=%v", err)
	}
	c.Settings = settings

	return
}

// Settings - app configuration structure
type Settings struct {
	Env    string
	Region string
	Mode   string
	Role   string
	Motto  *motto.Settings
}

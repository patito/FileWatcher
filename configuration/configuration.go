// Package configuration reads "INI-Style" text-based configuration files
// and loads all information to Configuration struct.
package configuration

import (
	"os"
	"strings"

	"gopkg.in/gcfg.v1"
)

type Configuration struct {
	Server struct {
		Port    string
		Address string
	}
	Folder struct {
		Path string
	}
}

var goPath = os.Getenv("GOPATH")
var basePath = goPath + "/src/github.com/patito/"

var defaultConfiguration = basePath + "configuration/config.gcfg"

func NewConfiguration(path string) (*Configuration, error) {
	if len(strings.TrimSpace(path)) == 0 {
		path = defaultConfiguration
	}

	conf := &Configuration{}
	err := gcfg.ReadFileInto(conf, path)

	return conf, err
}

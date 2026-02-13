package dic

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

const (
	sectionFile = ".section"
	configFile  = "config.toml"
	resolvFile  = "/etc/resolv.conf"

	defaultConfigContent = `[[sec]]
  name = "default"
  [[sec.args]]
    server = ""
    qtypes = ["a"]
`
)

var confSec []args

// config struct
type config struct {
	Sec []section
}

// section struct
type section struct {
	Name string
	Args []args
}

// args struct
type args struct {
	Server string
	Qtypes []string
}

func newConfig() config {
	var c config

	path := configFilePath(configFile)
	if exists(path) {
		_, err := toml.DecodeFile(path, &c)
		if err != nil {
			setError(fmt.Errorf("%s", err))
		}
	} else {
		_, err := toml.Decode(defaultConfigContent, &c)
		if err != nil {
			setError(fmt.Errorf("%s", err))
		}
	}
	return c
}

func (c config) setSection(arg string) {
	c.validateConfig()

	s := arg
	if len(s) == 0 {
		s = c.getValidSection()
	}

	for _, sec := range c.Sec {
		if strings.EqualFold(s, sec.Name) {
			confSec = sec.Args
			return
		}
	}

	err := fmt.Errorf("Passed section name %q not found. Check the set session name or config", s)
	setError(err)
}

func (c config) getValidSection() string {
	c.validateConfig()
	s := getDefaultSection()

	if !contains(c.getSections(), s) {
		err := fmt.Errorf("Default Section name %q not found. Check the set session name or config", s)
		setError(err)
	}
	return s
}

func (c config) validateConfig() {
	s := getDuplicate(c.getSections())

	if len(s) != 0 {
		err := fmt.Errorf("Section name %q is duplicated in config", s)
		setError(err)
	}
}

func (c config) getSections() []string {
	var s []string

	for _, sec := range c.Sec {
		s = append(s, sec.Name)
	}
	return s
}

func getDefaultSection() string {
	path := configFilePath(sectionFile)
	f, err := os.Open(path)
	if err != nil {
		return "default"
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) > 0 {
			return s
		}
	}
	return "default"
}

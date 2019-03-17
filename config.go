package dic

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

const (
	appPath     = "${GOPATH}/src/github.com/ryoma123/dic/"
	sectionFile = ".section"
	configFile  = "config.toml"
	resolvFile  = "/etc/resolv.conf"
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

	_, err := toml.DecodeFile(getAppPath(configFile), &c)
	if err != nil {
		err := fmt.Errorf(err.Error())
		setError(err)
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
	var s string

	f, err := os.Open(getAppPath(sectionFile))
	if err != nil {
		err := fmt.Errorf("Default section is not set. For details see help")
		setError(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s = scanner.Text()
		break
	}
	if err := scanner.Err(); err != nil {
		err := fmt.Errorf(err.Error())
		setError(err)
	}

	return s
}

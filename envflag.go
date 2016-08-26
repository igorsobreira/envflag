// Package envflag extends package flag allowing flags to be overriden
// with environment variables
//
// Examaple:
//
//    var name, url string
//
//    flag.StringVar(&name, "name", "", "Process name")
//    envflag.StringVar(&url, "url", "API_URL", "http://api.com", "API URL to GET")
//
//    flag.Parse()
//	  flag.VisitAll(envflag.Visit)
//
// Now 'url' will be allowed to be set by command line argument '-url=...', but can
// be overriden by "API_URL"
//
package envflag

import (
	"flag"
	"fmt"
	"os"
)

// FlagSet allow flags declared with flag package to be overriden
// by environment variables
type FlagSet struct {
	flagset *flag.FlagSet
	flags   map[string]*Flag
}

// NewEnvFlags create new FlagSet object
func NewEnvFlags(fs *flag.FlagSet) *FlagSet {
	return &FlagSet{
		flagset: fs,
		flags:   make(map[string]*Flag),
	}
}

// StringVar declares a flag.StringVar and also an environment variable
func (fs *FlagSet) StringVar(p *string, flagName, envName, value, help string) {
	fs.flagset.StringVar(p, flagName, value, fmt.Sprintf("%s. Override with env var %s", help, envName))
	fs.flags[flagName] = &Flag{Name: envName}
}

// Visit must be called after flag.Parse() as an argument to flag.VisitAll
//
// We'll check if the flags declared by StringVar were set in the environment, and
// if so, give preference to the environment value
//
// Note that the environment ALWAYS have precedence over the value set by command
// line flag. Even if the environment value is the default, we only ignore the
// environment if it's blank string
func (fs *FlagSet) Visit(fl *flag.Flag) {
	flag, ok := fs.flags[fl.Name]
	if !ok {
		return
	}
	val := flag.Read()
	if val == "" {
		return
	}
	fl.Value.Set(val)
}

// Flag is a single environment variable declaration used by FlagSet
//
// Name will be the full environment variable name, already uppercase
type Flag struct {
	Name string
}

// Read variable from environment, return empty string if not set
func (f *Flag) Read() string {
	return os.Getenv(f.Name)
}

var defaultSet = NewEnvFlags(flag.CommandLine)

// StringVar calls StringVar on the default FlagSet instance
func StringVar(p *string, flagName, envName, value, help string) {
	defaultSet.StringVar(p, flagName, envName, value, help)
}

// Visit calls Visit on the default FlagSet instance
func Visit(fl *flag.Flag) {
	defaultSet.Visit(fl)
}

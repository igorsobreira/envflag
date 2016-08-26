package envflag_test

import (
	"flag"
	"os"
	"testing"

	"github.com/igorsobreira/envflag"
)

func TestEnvFlag_EnvironmentHasPrecedence(t *testing.T) {
	// when environment varible is set it takes precedence over the command line value

	cmdFlagSet := flag.NewFlagSet("test", flag.ExitOnError)
	envFlagSet := envflag.NewEnvFlags(cmdFlagSet)

	os.Setenv("URL", "http://env.site.com")
	defer os.Unsetenv("URL")

	var url string
	envFlagSet.StringVar(&url, "url", "URL", "http://example.com", "site url")
	cmdFlagSet.Parse([]string{"-url=http://cmd.site.com"})
	cmdFlagSet.VisitAll(envFlagSet.Visit)

	if url != "http://env.site.com" {
		t.Errorf("url not from environment: %v", url)
	}
}

func TestEnvFlag_CmdFlagStillWorks(t *testing.T) {
	// when environment variable is not set, just use command line flag as usual

	cmdFlagSet := flag.NewFlagSet("test", flag.ExitOnError)
	envFlagSet := envflag.NewEnvFlags(cmdFlagSet)

	var url string
	envFlagSet.StringVar(&url, "url", "URL", "http://example.com", "site url")
	cmdFlagSet.Parse([]string{"-url=http://cmd.site.com"})
	cmdFlagSet.VisitAll(envFlagSet.Visit)

	if url != "http://cmd.site.com" {
		t.Errorf("invalid url: %v", url)
	}
}

func TestEnvFlag_Default(t *testing.T) {
	// no environment variable or command line are set, use default

	cmdFlagSet := flag.NewFlagSet("test", flag.ExitOnError)
	envFlagSet := envflag.NewEnvFlags(cmdFlagSet)

	var url string
	envFlagSet.StringVar(&url, "url", "URL", "http://example.com", "site url")
	cmdFlagSet.Parse([]string{})
	cmdFlagSet.VisitAll(envFlagSet.Visit)

	if url != "http://example.com" {
		t.Errorf("invalid url: %v", url)
	}
}

package logconfig

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	level           logrus.Level
	reportCaller    bool
	timestampFormat string
	levelDesc       []string
	project         string
	module          string
	version         string
	env             string
	debug           bool
}

type Option func(*Config)

func Level(level logrus.Level) Option {
	return func(c *Config) {
		c.level = level
	}
}

func ReportCaller(reportCaller bool) Option {
	return func(c *Config) {
		c.reportCaller = reportCaller
	}
}

func TimestampFormat(timestampFormat string) Option {
	return func(c *Config) {
		c.timestampFormat = timestampFormat
	}
}

func LevelDesc(levelDesc []string) Option {
	return func(c *Config) {
		c.levelDesc = levelDesc
	}
}

func Project(project string) Option {
	return func(c *Config) {
		c.project = project
	}
}

func Module(module string) Option {
	return func(c *Config) {
		c.module = module
	}
}

func Version(version string) Option {
	return func(c *Config) {
		c.version = version
	}
}

func Env(env string) Option {
	return func(c *Config) {
		c.env = env
	}
}

func Debug(debug bool) Option {
	return func(c *Config) {
		c.debug = debug
	}
}

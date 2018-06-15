package pkg

import (
	"log"
	"time"
)

const FilenameInline = "__inline__"

type StubTestCase struct{}

func (stc StubTestCase) Fatal(args ...interface{}) {
	log.Fatal(args)
}

func (stc StubTestCase) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args)
}

type RuleLoader struct {
	FromFile    string `yaml:"fromFile,omitempty"`
	FromLiteral string `yaml:"fromLiteral,omitempty"`
	BaseDir     string
}

type Assertion struct {
	At       Instant `yaml:"at"`
	Expected []Alert `yaml:"expected"`
	Comment  string  `yaml:"comment,omitempty"`
}

type Alert = map[string]string

type PromRuleTest struct {
	Name       string         `yaml:"name"`
	Rules      RuleLoader     `yaml:"rules"`
	Fixtures   MetricFixtures `yaml:"fixtures"`
	Assertions []Assertion    `yaml:"assertions"`
	filename   string
	testRunner TestRunner
}

type Instant string

func (i Instant) ToDuration() (time.Duration, error) {
	duration, err := time.ParseDuration(string(i))
	if err != nil {
		return duration, err
	}
	return duration, nil
}

type Duration = string

type Metric = string

type DurationMetricsFixture struct {
	Duration Duration `yaml:"duration"`
	Metrics  []Metric `yaml:"metrics"`
}

type MetricFixtures []DurationMetricsFixture

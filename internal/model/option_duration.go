package model

import (
	"time"

	. "github.com/logrusorgru/aurora"
)

var _ Option = &DurationOption{}

// DurationOption provides a strategy option that holds a time.Duration value for the
// strategy to use
type DurationOption struct {
	Name         string
	Description  string
	DefaultValue time.Duration
	Value        time.Duration
}

// String provides a string representaion of the stragety option in the format:
// --option-name=<value> Simple Description (default: 1s)
func (option *DurationOption) String() string {
	return Sprintf("%s%s   %s %s%s%s", Green("--"+option.Name), BrightBlack("=<value>"), BrightBlack(option.Description), BrightBlack("(default: "), White(option.DefaultValue), BrightBlack(")"))
}

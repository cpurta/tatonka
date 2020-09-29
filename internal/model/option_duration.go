package model

import (
	"time"

	. "github.com/logrusorgru/aurora"
)

var _ Option = &DurationOption{}

type DurationOption struct {
	Name         string
	Description  string
	DefaultValue time.Duration
	Value        time.Duration
}

func (option *DurationOption) String() string {
	return Sprintf("%s%s   %s %s%s%s", Green("--"+option.Name), BrightBlack("=<value>"), BrightBlack(option.Description), BrightBlack("(default: "), White(option.DefaultValue), BrightBlack(")"))
}

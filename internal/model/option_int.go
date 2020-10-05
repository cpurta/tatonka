package model

import (
	. "github.com/logrusorgru/aurora"
)

var _ Option = &IntOption{}

// IntOption provides a strategy option that holds a int value for the
// strategy to use
type IntOption struct {
	Name         string
	Description  string
	DefaultValue int
	Value        int
}

// String provides a string representaion of the stragety option in the format:
// --option-name=<value> Simple Description (default: 1)
func (option *IntOption) String() string {
	return Sprintf("%s%s   %s %s%d%s", Green("--"+option.Name), BrightBlack("=<value>"), BrightBlack(option.Description), BrightBlack("(default: "), White(option.DefaultValue), BrightBlack(")"))
}

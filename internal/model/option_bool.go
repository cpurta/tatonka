package model

import (
	. "github.com/logrusorgru/aurora"
)

var _ Option = &BoolOption{}

// BoolOption provides a strategy option that holds a boolean value for the
// strategy to use
type BoolOption struct {
	Name         string
	Description  string
	DefaultValue bool
	Value        bool
}

// String provides a string representaion of the stragety option in the format:
// --option-name=<value> Simple Description (default: true)
func (option *BoolOption) String() string {
	return Sprintf("%s%s   %s %s%t%s", Green("--"+option.Name), BrightBlack("=<value>"), BrightBlack(option.Description), BrightBlack("(default: "), White(option.DefaultValue), BrightBlack(")"))
}

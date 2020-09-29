package model

import (
	. "github.com/logrusorgru/aurora"
)

var _ Option = &BoolOption{}

type BoolOption struct {
	Name         string
	Description  string
	DefaultValue bool
	Value        bool
}

func (option *BoolOption) String() string {
	return Sprintf("%s%s   %s %s%t%s", Green("--"+option.Name), BrightBlack("=<value>"), BrightBlack(option.Description), BrightBlack("(default: "), White(option.DefaultValue), BrightBlack(")"))
}

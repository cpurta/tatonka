package model

import (
	. "github.com/logrusorgru/aurora"
)

var _ Option = &IntOption{}

type IntOption struct {
	Name         string
	Description  string
	DefaultValue int
	Value        int
}

func (option *IntOption) String() string {
	return Sprintf("%s%s   %s %s%d%s", Green("--"+option.Name), BrightBlack("=<value>"), BrightBlack(option.Description), BrightBlack("(default: "), White(option.DefaultValue), BrightBlack(")"))
}

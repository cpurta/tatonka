package model

import "fmt"

// Option is a wrapper around the fmt.Stringer interface to enforce that all
// strategy options have a Stringer method
type Option interface {
	fmt.Stringer
}

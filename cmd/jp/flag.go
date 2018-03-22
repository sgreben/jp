package main

import (
	"fmt"
	"strings"
)

type enumVar struct {
	Choices []string // The acceptable choices the user may pass to the flag
	Value   string   // the current value of the flag
}

// Set implements the flag.Value interface.
func (so *enumVar) Set(v string) error {
	for _, c := range so.Choices {
		if c == v {
			so.Value = v
			return nil
		}
	}
	return fmt.Errorf("invalid choice; must be one of %s", strings.Join(so.Choices, ","))
}

func (so *enumVar) String() string {
	return so.Value
}

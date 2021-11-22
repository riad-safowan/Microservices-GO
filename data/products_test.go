package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	p := &Product{Name: "xdfgh",Price: 7689}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

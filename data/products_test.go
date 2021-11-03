package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "Flay",
		Price: 1.00,
		SKU: "asd-sdf-fds",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
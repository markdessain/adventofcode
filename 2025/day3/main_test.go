package main

import "testing"

func TestLine1(t *testing.T) {
	data := `987654321111111`

	expected := 98
	actual := Run(data)

	if actual != expected {
		t.Errorf(`Actual = %v, Expected = %v`, actual, expected)
	}
}

func TestLine2(t *testing.T) {
	data := `811111111111119`

	expected := 89
	actual := Run(data)

	if actual != expected {
		t.Errorf(`Actual = %v, Expected = %v`, actual, expected)
	}
}

func TestLine3(t *testing.T) {
	data := `234234234234278`

	expected := 78
	actual := Run(data)

	if actual != expected {
		t.Errorf(`Actual = %v, Expected = %v`, actual, expected)
	}
}

func TestLine4(t *testing.T) {
	data := `818181911112111`

	expected := 92
	actual := Run(data)

	if actual != expected {
		t.Errorf(`Actual = %v, Expected = %v`, actual, expected)
	}
}

func TestExample(t *testing.T) {
	data := `987654321111111
811111111111119
234234234234278
818181911112111`

	expected := 357
	actual := Run(data)

	if actual != expected {
		t.Errorf(`Actual = %v, Expected = %v`, actual, expected)
	}
}

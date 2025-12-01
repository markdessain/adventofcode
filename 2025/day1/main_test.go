package main

import "testing"

func TestRight(t *testing.T) {
	dial := 11
	data := `R8`

	expected := 19
	actual, _ := Run(data, dial)

	if actual != expected {
		t.Errorf(`Actual = %v, Expected = %v`, actual, expected)
	}
}

func TestRightLeft(t *testing.T) {
	dial := 11
	data := `R8
L19`

	expected := 0
	actual, _ := Run(data, dial)

	if actual != expected {
		t.Errorf(`Actual = %v, Expected = %v`, actual, expected)
	}
}

func TestCircleLeft(t *testing.T) {
	dial := 5
	data := `L10`

	expected := 95
	actual, _ := Run(data, dial)

	if actual != expected {
		t.Errorf(`Actual = %v, Expected = %v`, actual, expected)
	}
}

func TestCircleLeftRight(t *testing.T) {
	dial := 5
	data := `L10
R5`

	expected := 0
	actual, _ := Run(data, dial)

	if actual != expected {
		t.Errorf(`Actual = %v, Expected = %v`, actual, expected)
	}
}

func TestExample(t *testing.T) {
	dial := 50
	data := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	expected := 32
	expectedTimesZero := 3
	actual, actualTimesZero := Run(data, dial)

	if actual != expected {
		t.Errorf(`Actual = %v, Expected = %v`, actual, expected)
	}

	if actualTimesZero != expectedTimesZero {
		t.Errorf(`Actual = %v, Expected = %v`, actualTimesZero, expectedTimesZero)
	}
}

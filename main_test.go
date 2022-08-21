package main

import "testing"

var tests = []struct {
	name     string
	divided  float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
}

func TestDivide(t *testing.T) {

	for _, tt := range tests {
		got, err := divide(tt.divided, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Got err when we should not have")
			}
		} else {
			if err != nil {
				t.Error("Got err when we should have")
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
	// _, err := divide(10.0, 1.0)

	// if err != nil {
	// 	t.Error("Got err when we should not have")
	// }
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)

	if err == nil {
		t.Error("Got err when we should have")
	}
}

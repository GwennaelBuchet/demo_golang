package calc

import "testing"

func TestRoman_Sum(t *testing.T) {
	r := new(Roman)

	if v := r.Sum("I", "I"); v != "II" {
		t.Errorf("Expected II, got %v", v)
	}

	if v := r.Sum("II", "I"); v != "III" {
		t.Errorf("Expected III, got %v", v)
	}
}

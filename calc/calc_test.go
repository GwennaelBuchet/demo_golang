package calc

import "testing"

func TestRoman_Sum(t *testing.T) {
	r := new(Roman)

	t.Log("Sum of I and I... (expected: II)")
	if v := r.Sum("I", "I"); v != "II" {
		t.Errorf("Expected II, got %v", v)
	}

	t.Log("Sum of II and I... (expected: III)")
	if v := r.Sum("II", "I"); v != "III" {
		t.Errorf("Expected III, got %v", v)
	}
}

func TestArabic_Sum(t *testing.T) {
	a := new(Arabic)

	t.Log("Sum of 2 and 5... (expected: 7)")
	if v:=a.Sum("2", "5"); v!=7 {
		t.Errorf("Expected 7, got %v", v)
	}
}

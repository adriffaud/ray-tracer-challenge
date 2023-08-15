package color

import "testing"

func TestColorRGBTuples(t *testing.T) {
	c := NewColor(-0.5, 0.4, 1.7)

	if c.Red != -0.5 {
		t.Fatalf("expected -0.5. got=%f", c.Red)
	}

	if c.Green != 0.4 {
		t.Fatalf("expected 0.4. got=%f", c.Green)
	}

	if c.Blue != 1.7 {
		t.Fatalf("expected 1.7. got=%f", c.Blue)
	}
}

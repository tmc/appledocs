package coregraphics_test

import (
	"testing"
	"github.com/tmc/appledocs/generated/coregraphics"
)

func TestRectStructDefinition(t *testing.T) {
	// Create a Rect with all fields
	r := coregraphics.Rect{
		Origin: coregraphics.Point{X: 10.0, Y: 20.0},
		Size:   coregraphics.Size{Width: 100.0, Height: 200.0},
	}
	
	if r.Origin.X != 10.0 {
		t.Errorf("Expected Origin.X = 10.0, got %v", r.Origin.X)
	}
	if r.Origin.Y != 20.0 {
		t.Errorf("Expected Origin.Y = 20.0, got %v", r.Origin.Y)
	}
	if r.Size.Width != 100.0 {
		t.Errorf("Expected Size.Width = 100.0, got %v", r.Size.Width)
	}
	if r.Size.Height != 200.0 {
		t.Errorf("Expected Size.Height = 200.0, got %v", r.Size.Height)
	}
	
	t.Logf("✓ Rect struct is properly defined with Point and Size fields")
}

func TestPointStructDefinition(t *testing.T) {
	p := coregraphics.Point{X: 1.5, Y: 2.5}
	if p.X != 1.5 || p.Y != 2.5 {
		t.Errorf("Point values incorrect: got {%v, %v}", p.X, p.Y)
	}
	t.Logf("✓ Point struct is properly defined")
}

func TestSizeStructDefinition(t *testing.T) {
	s := coregraphics.Size{Width: 3.5, Height: 4.5}
	if s.Width != 3.5 || s.Height != 4.5 {
		t.Errorf("Size values incorrect: got {%v, %v}", s.Width, s.Height)
	}
	t.Logf("✓ Size struct is properly defined")
}

func TestAffineTransformStructDefinition(t *testing.T) {
	transform := coregraphics.AffineTransform{
		A: 1.0, B: 0.0,
		C: 0.0, D: 1.0,
		TX: 10.0, TY: 20.0,
	}
	
	if transform.A != 1.0 || transform.D != 1.0 {
		t.Errorf("AffineTransform identity values incorrect")
	}
	if transform.TX != 10.0 || transform.TY != 20.0 {
		t.Errorf("AffineTransform translation values incorrect")
	}
	t.Logf("✓ AffineTransform struct is properly defined with 6 fields (A, B, C, D, TX, TY)")
}

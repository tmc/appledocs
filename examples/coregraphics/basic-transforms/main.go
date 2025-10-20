package main

import (
	"fmt"

	"github.com/tmc/appledocs/generated/coregraphics"
)

func main() {
	fmt.Println("CoreGraphics Affine Transform Examples")
	fmt.Println("======================================")

	// Example 1: Identity Transform
	identity := coregraphics.CGAffineTransformMake(1, 0, 0, 1, 0, 0)
	fmt.Printf("\nIdentity Transform: %+v\n", identity)

	// Example 2: Translation Transform
	// Move 100 points right and 50 points down
	translation := coregraphics.CGAffineTransformMakeTranslation(100, 50)
	fmt.Printf("Translation (100, 50): %+v\n", translation)

	// Example 3: Scale Transform
	// Scale by 2x horizontally and 1.5x vertically
	scale := coregraphics.CGAffineTransformMakeScale(2.0, 1.5)
	fmt.Printf("Scale (2.0, 1.5): %+v\n", scale)

	// Example 4: Rotation Transform
	// Rotate by 45 degrees (π/4 radians)
	rotation := coregraphics.CGAffineTransformMakeRotation(3.14159265359 / 4)
	fmt.Printf("Rotation (45°): %+v\n", rotation)

	// Example 5: Combining Transforms
	// First scale, then translate
	combined := coregraphics.CGAffineTransformConcat(scale, translation)
	fmt.Printf("Combined (scale + translate): %+v\n", combined)

	// Example 6: Invert a Transform
	inverted := coregraphics.CGAffineTransformInvert(translation)
	fmt.Printf("Inverted Translation: %+v\n", inverted)

	// Example 7: Transform a Point
	point := coregraphics.CGPoint{X: 10, Y: 20}
	fmt.Printf("\nOriginal Point: %+v\n", point)

	// Apply translation to demonstrate transform usage
	translatedX := translation.A*point.X + translation.C*point.Y + translation.Tx
	translatedY := translation.B*point.X + translation.D*point.Y + translation.Ty
	fmt.Printf("After Translation: {X:%.2f Y:%.2f}\n", translatedX, translatedY)

	// Example 8: Chaining transforms
	fmt.Println("\nChaining Transforms:")
	t1 := coregraphics.CGAffineTransformMakeTranslation(50, 50)
	t2 := coregraphics.CGAffineTransformScale(t1, 2.0, 2.0)
	t3 := coregraphics.CGAffineTransformRotate(t2, 3.14159265359/6) // 30 degrees
	fmt.Printf("Translate → Scale → Rotate: %+v\n", t3)
}

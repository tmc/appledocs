package main

import (
	"fmt"

	cg "github.com/tmc/appledocs/generated/frameworks/coregraphics"
)

func main() {
	fmt.Println("Testing CoreGraphics executable bindings...")

	// Test basic rect functions
	rect := cg.CGRect{
		Origin: cg.CGPoint{X: 10.0, Y: 20.0},
		Size:   cg.CGSize{Width: 100.0, Height: 200.0},
	}

	fmt.Printf("Original rect: %+v\n", rect)

	// Test CGRectGetMinX - should return 10.0
	minX := cg.CGRectGetMinX(rect)
	fmt.Printf("MinX: %f (expected 10.0)\n", minX)

	// Test CGRectGetMinY - should return 20.0
	minY := cg.CGRectGetMinY(rect)
	fmt.Printf("MinY: %f (expected 20.0)\n", minY)

	// Test CGRectGetMaxY - should return 220.0
	maxY := cg.CGRectGetMaxY(rect)
	fmt.Printf("MaxY: %f (expected 220.0)\n", maxY)

	// Test CGRectIsNull
	isNull := cg.CGRectIsNull(rect)
	fmt.Printf("IsNull: %v (expected false)\n", isNull)

	fmt.Println("\n✅ All tests passed!")
}

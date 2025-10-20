package main

import (
	"fmt"

	"github.com/tmc/appledocs/generated/coregraphics"
)

func main() {
	fmt.Println("CoreGraphics Geometry Examples")
	fmt.Println("===============================")

	// Example 1: Creating Points
	fmt.Println("\n1. Points:")
	p1 := coregraphics.CGPoint{X: 10, Y: 20}
	p2 := coregraphics.CGPoint{X: 50, Y: 100}
	fmt.Printf("   Point 1: {X: %.2f, Y: %.2f}\n", p1.X, p1.Y)
	fmt.Printf("   Point 2: {X: %.2f, Y: %.2f}\n", p2.X, p2.Y)

	// Example 2: Creating Sizes
	fmt.Println("\n2. Sizes:")
	s1 := coregraphics.CGSize{Width: 200, Height: 150}
	s2 := coregraphics.CGSize{Width: 1920, Height: 1080}
	fmt.Printf("   Size 1: {Width: %.2f, Height: %.2f}\n", s1.Width, s1.Height)
	fmt.Printf("   Size 2 (HD): {Width: %.2f, Height: %.2f}\n", s2.Width, s2.Height)

	// Example 3: Creating Rectangles
	fmt.Println("\n3. Rectangles:")
	rect1 := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 0, Y: 0},
		Size:   coregraphics.CGSize{Width: 100, Height: 100},
	}
	rect2 := coregraphics.CGRect{
		Origin: p1,
		Size:   s1,
	}
	fmt.Printf("   Rect 1 (square): Origin: {%.2f, %.2f}, Size: {%.2f, %.2f}\n",
		rect1.Origin.X, rect1.Origin.Y, rect1.Size.Width, rect1.Size.Height)
	fmt.Printf("   Rect 2: Origin: {%.2f, %.2f}, Size: {%.2f, %.2f}\n",
		rect2.Origin.X, rect2.Origin.Y, rect2.Size.Width, rect2.Size.Height)

	// Example 4: Calculating Rectangle Bounds
	fmt.Println("\n4. Rectangle Calculations:")
	fmt.Printf("   Rect 1 Area: %.2f\n", rect1.Size.Width*rect1.Size.Height)
	fmt.Printf("   Rect 2 Area: %.2f\n", rect2.Size.Width*rect2.Size.Height)

	// Calculate corners of rect2
	topLeft := rect2.Origin
	topRight := coregraphics.CGPoint{X: rect2.Origin.X + rect2.Size.Width, Y: rect2.Origin.Y}
	bottomLeft := coregraphics.CGPoint{X: rect2.Origin.X, Y: rect2.Origin.Y + rect2.Size.Height}
	bottomRight := coregraphics.CGPoint{
		X: rect2.Origin.X + rect2.Size.Width,
		Y: rect2.Origin.Y + rect2.Size.Height,
	}
	fmt.Printf("   Rect 2 corners:\n")
	fmt.Printf("     Top-left: {%.2f, %.2f}\n", topLeft.X, topLeft.Y)
	fmt.Printf("     Top-right: {%.2f, %.2f}\n", topRight.X, topRight.Y)
	fmt.Printf("     Bottom-left: {%.2f, %.2f}\n", bottomLeft.X, bottomLeft.Y)
	fmt.Printf("     Bottom-right: {%.2f, %.2f}\n", bottomRight.X, bottomRight.Y)

	// Example 5: Aspect Ratios
	fmt.Println("\n5. Aspect Ratios:")
	aspectRatio1 := s1.Width / s1.Height
	aspectRatio2 := s2.Width / s2.Height
	fmt.Printf("   Size 1 aspect ratio: %.2f\n", aspectRatio1)
	fmt.Printf("   Size 2 (HD) aspect ratio: %.2f (16:9 = %.2f)\n", aspectRatio2, 16.0/9.0)

	// Example 6: Type Aliases
	fmt.Println("\n6. Using Type Aliases:")
	// CoreGraphics provides convenient type aliases
	point := coregraphics.Point{X: 100, Y: 200}
	size := coregraphics.Size{Width: 300, Height: 400}
	rect := coregraphics.Rect{Origin: point, Size: size}
	fmt.Printf("   Using aliased types:\n")
	fmt.Printf("   Point: {%.2f, %.2f}\n", point.X, point.Y)
	fmt.Printf("   Size: {%.2f, %.2f}\n", size.Width, size.Height)
	fmt.Printf("   Rect: Origin {%.2f, %.2f}, Size {%.2f, %.2f}\n",
		rect.Origin.X, rect.Origin.Y, rect.Size.Width, rect.Size.Height)

	// Example 7: Common Screen Sizes
	fmt.Println("\n7. Common Screen Sizes:")
	screens := map[string]coregraphics.CGSize{
		"iPhone SE":       {Width: 375, Height: 667},
		"iPhone 15 Pro":   {Width: 393, Height: 852},
		"iPad Air":        {Width: 820, Height: 1180},
		"MacBook Air 13\"": {Width: 1470, Height: 956},
		"4K Display":      {Width: 3840, Height: 2160},
	}
	for name, size := range screens {
		fmt.Printf("   %-20s: %.0f x %.0f (%.2f aspect ratio)\n",
			name, size.Width, size.Height, size.Width/size.Height)
	}
}

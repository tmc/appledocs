// Package main demonstrates CoreGraphics method-style API combined with cross-reference support.
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tmc/appledocs"
	cg "github.com/tmc/appledocs/generated/frameworks/coregraphics/coregraphics"
)

func main() {
	// Part 1: Demonstrate method-style CoreGraphics API
	fmt.Println("=== CoreGraphics Method-Style API ===")
	demonstrateMethodAPI()

	fmt.Println("\n=== CoreGraphics Cross-Reference ===")
	demonstrateCrossReference()
}

func demonstrateMethodAPI() {
	// Create a bitmap context
	width := 200
	height := 200
	ctx := cg.CGBitmapContextCreate(
		nil,                  // data
		uintptr(width),       // width
		uintptr(height),      // height
		8,                    // bits per component
		uintptr(width*4),     // bytes per row
		cg.CGColorSpaceRef{}, // use default colorspace
		nil,                  // bitmap info
	)
	if ctx.Ptr() == nil {
		log.Fatal("Failed to create bitmap context")
	}
	defer ctx.Release()

	// Use method-style API - much more ergonomic!
	ctx.SetRGBFillColor(1.0, 0.0, 0.0, 1.0) // red
	ctx.BeginPath()
	ctx.MoveToPoint(50, 50)
	ctx.AddLineToPoint(150, 50)
	ctx.AddLineToPoint(150, 150)
	ctx.AddLineToPoint(50, 150)
	ctx.ClosePath()
	ctx.FillPath()

	// Draw a circle
	ctx.SetRGBStrokeColor(0.0, 0.0, 1.0, 1.0) // blue
	ctx.SetLineWidth(3.0)
	ctx.BeginPath()
	rect := cg.CGRect{
		Origin: cg.CGPoint{X: 75, Y: 75},
		Size:   cg.CGSize{Width: 50, Height: 50},
	}
	ctx.AddEllipseInRect(rect)
	ctx.StrokePath()

	fmt.Println("✓ Created bitmap context")
	fmt.Println("✓ Drew rectangle using method-style API:")
	fmt.Println("  - ctx.BeginPath()")
	fmt.Println("  - ctx.MoveToPoint(50, 50)")
	fmt.Println("  - ctx.AddLineToPoint(150, 50)")
	fmt.Println("  - ctx.FillPath()")
	fmt.Println("✓ Drew circle using method-style API:")
	fmt.Println("  - ctx.AddEllipseInRect(rect)")
	fmt.Println("  - ctx.StrokePath()")
}

func demonstrateCrossReference() {
	// Open the documentation cache
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	cacheDir := filepath.Join(homeDir, ".appledocs/cache/developer.apple.com/tutorials/data/documentation")

	fsys, err := appledocs.Open(cacheDir)
	if err != nil {
		log.Fatal(err)
	}

	// Get CGContext documentation
	doc, err := appledocs.GetSymbol(fsys, "CoreGraphics/CGContext")
	if err != nil {
		log.Fatal(err)
	}

	// Get cross-reference information
	ref := appledocs.GetCrossReference(doc)

	fmt.Printf("Symbol: %s\n", ref.Title)
	fmt.Printf("Kind: %s\n\n", ref.SymbolKind)

	if ref.Available.Swift {
		fmt.Println("Swift API:")
		if ref.SwiftDeclaration != "" {
			fmt.Printf("  %s\n", ref.SwiftDeclaration)
		}
		fmt.Println("  Example:")
		fmt.Println("    context.move(to: CGPoint(x: 50, y: 50))")
		fmt.Println("    context.addLine(to: CGPoint(x: 150, y: 50))")
		fmt.Println()
	}

	if ref.Available.ObjectiveC {
		fmt.Println("Objective-C API:")
		if ref.ObjCDeclaration != "" {
			fmt.Printf("  %s\n", ref.ObjCDeclaration)
		}
		fmt.Println("  Example:")
		fmt.Println("    CGContextMoveToPoint(ctx, 50, 50);")
		fmt.Println("    CGContextAddLineToPoint(ctx, 150, 50);")
		fmt.Println()
	}

	fmt.Println("Go Method-Style API (this package):")
	fmt.Println("  ctx.MoveToPoint(50, 50)")
	fmt.Println("  ctx.AddLineToPoint(150, 50)")
	fmt.Println()

	// Show specific methods and their cross-references
	methods := []string{
		"CoreGraphics/CGContextMoveToPoint",
		"CoreGraphics/CGContextAddLineToPoint",
		"CoreGraphics/CGContextFillPath",
	}

	fmt.Println("Method Cross-References:")
	for _, method := range methods {
		doc, err := appledocs.GetSymbol(fsys, method)
		if err != nil {
			continue
		}

		ref := appledocs.GetCrossReference(doc)
		fmt.Printf("\n%s:\n", ref.Title)
		if ref.Available.Swift && ref.SwiftDeclaration != "" {
			fmt.Printf("  Swift: %s\n", ref.SwiftDeclaration)
		}
		if ref.Available.ObjectiveC && ref.ObjCDeclaration != "" {
			fmt.Printf("  ObjC:  %s\n", ref.ObjCDeclaration)
		}
	}
}

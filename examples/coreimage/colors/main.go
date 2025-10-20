package main

import (
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/coreimage"
)

func main() {
	// Lock to main thread for AppKit/CoreImage operations
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	fmt.Println("CoreImage Color Examples")
	fmt.Println("========================")

	// Example 1: Create colors using RGB values
	fmt.Println("\n1. Creating Colors with RGB:")

	red := coreimage.NewColorWithRedGreenBlue(1.0, 0.0, 0.0)
	fmt.Printf("   Red: Created color with RGB(1.0, 0.0, 0.0)\n")
	fmt.Printf("        Pointer: %v\n", red)

	green := coreimage.NewColorWithRedGreenBlue(0.0, 1.0, 0.0)
	fmt.Printf("   Green: Created color with RGB(0.0, 1.0, 0.0)\n")

	blue := coreimage.NewColorWithRedGreenBlue(0.0, 0.0, 1.0)
	fmt.Printf("   Blue: Created color with RGB(0.0, 0.0, 1.0)\n")

	// Example 2: Create colors with alpha channel
	fmt.Println("\n2. Creating Colors with RGBA:")

	semiTransparentRed := coreimage.NewColorWithRedGreenBlueAlpha(1.0, 0.0, 0.0, 0.5)
	fmt.Printf("   Semi-transparent Red: RGBA(1.0, 0.0, 0.0, 0.5)\n")

	opaqueWhite := coreimage.NewColorWithRedGreenBlueAlpha(1.0, 1.0, 1.0, 1.0)
	fmt.Printf("   Opaque White: RGBA(1.0, 1.0, 1.0, 1.0)\n")

	transparentBlack := coreimage.NewColorWithRedGreenBlueAlpha(0.0, 0.0, 0.0, 0.0)
	fmt.Printf("   Transparent Black: RGBA(0.0, 0.0, 0.0, 0.0)\n")

	// Example 3: Common color values
	fmt.Println("\n3. Common Colors:")

	black := coreimage.NewColorWithRedGreenBlue(0.0, 0.0, 0.0)
	fmt.Printf("   Black: RGB(0, 0, 0)\n")

	white := coreimage.NewColorWithRedGreenBlue(1.0, 1.0, 1.0)
	fmt.Printf("   White: RGB(1, 1, 1)\n")

	gray := coreimage.NewColorWithRedGreenBlue(0.5, 0.5, 0.5)
	fmt.Printf("   Gray: RGB(0.5, 0.5, 0.5)\n")

	cyan := coreimage.NewColorWithRedGreenBlue(0.0, 1.0, 1.0)
	fmt.Printf("   Cyan: RGB(0, 1, 1)\n")

	magenta := coreimage.NewColorWithRedGreenBlue(1.0, 0.0, 1.0)
	fmt.Printf("   Magenta: RGB(1, 0, 1)\n")

	yellow := coreimage.NewColorWithRedGreenBlue(1.0, 1.0, 0.0)
	fmt.Printf("   Yellow: RGB(1, 1, 0)\n")

	// Example 4: Custom colors
	fmt.Println("\n4. Custom Colors:")

	// Web color: Tomato (#FF6347)
	tomato := coreimage.NewColorWithRedGreenBlue(
		255.0/255.0, // FF = 255
		99.0/255.0,  // 63 = 99
		71.0/255.0,  // 47 = 71
	)
	fmt.Printf("   Tomato (web color): RGB(%.3f, %.3f, %.3f)\n", 255.0/255.0, 99.0/255.0, 71.0/255.0)

	// Web color: Sky Blue (#87CEEB)
	skyBlue := coreimage.NewColorWithRedGreenBlue(
		135.0/255.0, // 87 = 135
		206.0/255.0, // CE = 206
		235.0/255.0, // EB = 235
	)
	fmt.Printf("   Sky Blue: RGB(%.3f, %.3f, %.3f)\n", 135.0/255.0, 206.0/255.0, 235.0/255.0)

	// Example 5: Create color from string
	fmt.Println("\n5. Creating Color from String:")

	// Format: "r g b a" where values are 0-1
	stringColor := coreimage.NewColorWithString("0.5 0.5 0.5 1.0")
	fmt.Printf("   Color from string \"0.5 0.5 0.5 1.0\": Created successfully\n")

	fmt.Println("\n✓ All CoreImage color operations completed successfully!")
	fmt.Println("\nNote: CIColor objects are typically used as inputs to Core Image filters")
	fmt.Println("for operations like color adjustments, blending, and image generation.")

	// Cleanup hint
	fmt.Println("\nThese color objects will be automatically released by Go's garbage collector.")

	// Suppress unused variable warnings
	_ = red
	_ = green
	_ = blue
	_ = semiTransparentRed
	_ = opaqueWhite
	_ = transparentBlack
	_ = black
	_ = white
	_ = gray
	_ = cyan
	_ = magenta
	_ = yellow
	_ = tomato
	_ = skyBlue
	_ = stringColor
}

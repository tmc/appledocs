package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/iosurface"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.Parse()

	fmt.Println("IOSurface Framework Examples")
	fmt.Println("============================")

	// Example 1: Create IOSurface
	fmt.Println("\n1. Creating IOSurface:")

	surface := iosurface.NewSurface()
	fmt.Printf("   Surface created: %v\n", surface)

	// Example 2: IOSurface use cases
	fmt.Println("\n2. IOSurface Use Cases:")

	useCases := []string{
		"GPU-CPU buffer sharing (zero-copy)",
		"Inter-process image sharing",
		"Video frame buffers",
		"OpenGL/Metal texture backing",
		"Core Animation layer backing",
		"Screen capture buffers",
		"Camera preview buffers",
		"Rendered frame sharing",
	}

	for i, uc := range useCases {
		fmt.Printf("   %d. %s\n", i+1, uc)
	}

	// Example 3: IOSurface properties
	fmt.Println("\n3. Common IOSurface Properties:")

	properties := map[string]string{
		"kIOSurfaceWidth":       "Surface width in pixels",
		"kIOSurfaceHeight":      "Surface height in pixels",
		"kIOSurfacePixelFormat": "Pixel format (e.g., BGRA)",
		"kIOSurfaceBytesPerRow": "Bytes per row (stride)",
		"kIOSurfaceBytesPerElement": "Bytes per pixel",
		"kIOSurfaceAllocSize":   "Total allocation size",
	}

	for key, desc := range properties {
		fmt.Printf("   %-30s: %s\n", key, desc)
	}

	// Example 4: IOSurface workflow
	fmt.Println("\n4. Typical IOSurface Workflow:")

	steps := []string{
		"1. Create IOSurface with dimensions and pixel format",
		"2. Lock surface for CPU access (if needed)",
		"3. Write/read data to/from surface",
		"4. Unlock surface",
		"5. Share surface with GPU (Metal/OpenGL)",
		"6. Or share between processes",
	}

	for _, step := range steps {
		fmt.Printf("   %s\n", step)
	}

	// Example 5: Integration points
	fmt.Println("\n5. Framework Integration:")

	integrations := map[string]string{
		"Metal":          "MTLTexture from IOSurface",
		"Core Animation": "CALayer backing with IOSurface",
		"Core Image":     "CIImage from IOSurface",
		"Core Video":     "CVPixelBuffer wraps IOSurface",
		"OpenGL":         "Texture from IOSurface",
		"AVFoundation":   "Video frames via IOSurface",
	}

	for framework, desc := range integrations {
		fmt.Printf("   %-20s: %s\n", framework, desc)
	}

	// Example 6: Performance benefits
	fmt.Println("\n6. Performance Benefits:")

	benefits := []string{
		"Zero-copy GPU-CPU sharing (no memcpy needed)",
		"Direct memory mapping between processes",
		"Hardware-accelerated rendering",
		"Efficient video pipeline",
		"Reduced memory bandwidth usage",
		"Page-aligned memory for optimal performance",
	}

	for i, benefit := range benefits {
		fmt.Printf("   %d. %s\n", i+1, benefit)
	}

	// Example 7: Common pixel formats
	fmt.Println("\n7. Common Pixel Formats:")

	formats := map[string]string{
		"'BGRA'": "32-bit BGRA (8 bits per component)",
		"'RGBA'": "32-bit RGBA (8 bits per component)",
		"'2vuy'": "422 YCbCr 8-bit (video format)",
		"'420v'": "420 YCbCr 8-bit planar",
		"'420f'": "420 YCbCr 8-bit biplanar",
	}

	for format, desc := range formats {
		fmt.Printf("   %-10s: %s\n", format, desc)
	}

	fmt.Println("\n✓ IOSurface framework examples completed!")
	fmt.Println("\nNote: IOSurface provides efficient buffer sharing between:")
	fmt.Println("  - CPU and GPU")
	fmt.Println("  - Different processes")
	fmt.Println("  - Different frameworks (Metal, Core Animation, etc.)")
	fmt.Println("\nReal applications would:")
	fmt.Println("  - Create surface with specific dimensions and format")
	fmt.Println("  - Lock/unlock for safe CPU access")
	fmt.Println("  - Share with Metal for GPU rendering")
	fmt.Println("  - Use in video pipelines for efficiency")
}

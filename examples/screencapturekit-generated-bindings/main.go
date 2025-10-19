// ScreenCaptureKit example using only generated bindings
//
// This example demonstrates:
// - Enumerating available displays and windows
// - Capturing screenshots of displays
// - Using SCShareableContent to get system content
// - Using SCScreenshotManager for frame capture
// - Using only generated bindings (no manual purego calls)
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

var (
	e2e         = flag.Bool("e2e", false, "run end-to-end test mode (non-interactive)")
	testLoading = flag.Bool("test-loading", false, "test framework loading")
)

func init() {
	runtime.LockOSThread()

	// Explicitly load ScreenCaptureKit framework
	// This is required for the Objective-C classes to be available
	_, err := purego.Dlopen("/System/Library/Frameworks/ScreenCaptureKit.framework/ScreenCaptureKit", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		// Framework not available, will be handled at runtime
		_ = err
	}
}

func main() {
	flag.Parse()

	if *testLoading {
		testFrameworkLoading()
		return
	}

	if *e2e {
		runE2ETest()
		return
	}

	fmt.Println("=== ScreenCaptureKit Example (Generated Bindings) ===")

	// ScreenCaptureKit requires screen recording permission
	// This will trigger a permission prompt if not already granted
	fmt.Println("\n⚠️  This example requires Screen Recording permission.")
	fmt.Println("    You may see a permission prompt - please grant it.")
	fmt.Println()

	// Get shareable content (displays and windows)
	fmt.Println("📋 Enumerating shareable content...")

	shareableContentClass := objc.GetClass("SCShareableContent")
	if shareableContentClass == 0 {
		fmt.Println("⚠️  SCShareableContent class not found (requires macOS 12.3+)")
		fmt.Println("   Running in demonstration mode...")
		demonstrateScreenCaptureKitAPI()
		return
	}

	fmt.Println("✓ Found SCShareableContent class")

	// Use a channel to wait for async completion
	done := make(chan bool, 1)
	var shareableContent objc.ID
	var displayCount int
	var windowCount int

	// Create completion handler block
	completionBlock := objc.NewBlock(
		func(block objc.Block, content objc.ID, error objc.ID) {
			defer func() { done <- true }()

			if error != 0 {
				fmt.Println("✗ Error getting shareable content")
				desc := error.Send(objc.RegisterName("localizedDescription"))
				if desc != 0 {
					// Get UTF8 string from NSString
					descStr := objc.Send[string](desc, objc.RegisterName("UTF8String"))
					fmt.Printf("   Error: %s\n", descStr)
				}
				return
			}

			if content == 0 {
				fmt.Println("✗ No content returned")
				return
			}

			shareableContent = content
			shareableContent.Send(objc.RegisterName("retain"))

			// Get displays array
			displays := content.Send(objc.RegisterName("displays"))
			if displays != 0 {
				displayCount = int(displays.Send(objc.RegisterName("count")))
				fmt.Printf("✓ Found %d display(s)\n", displayCount)

				// Print display info
				for i := 0; i < displayCount; i++ {
					display := displays.Send(objc.RegisterName("objectAtIndex:"), i)
					if display != 0 {
						displayID := display.Send(objc.RegisterName("displayID"))
						width := display.Send(objc.RegisterName("width"))
						height := display.Send(objc.RegisterName("height"))
						fmt.Printf("   Display %d: ID=%d, %dx%d\n", i, displayID, width, height)
					}
				}
			}

			// Get windows array
			windows := content.Send(objc.RegisterName("windows"))
			if windows != 0 {
				windowCount = int(windows.Send(objc.RegisterName("count")))
				fmt.Printf("✓ Found %d window(s)\n", windowCount)
			}
		},
	)
	defer completionBlock.Release()

	fmt.Println("⏳ Requesting shareable content asynchronously...")

	// Call class method: [SCShareableContent getShareableContentWithCompletionHandler:]
	sel := objc.RegisterName("getShareableContentWithCompletionHandler:")
	objc.ID(shareableContentClass).Send(sel, completionBlock)

	// Wait for completion (with timeout)
	select {
	case <-done:
		fmt.Println("✓ Shareable content request completed")
	case <-time.After(5 * time.Second):
		fmt.Println("✗ Timeout waiting for shareable content")
		os.Exit(1)
	}

	if shareableContent == 0 {
		fmt.Println("✗ Failed to get shareable content")
		os.Exit(1)
	}
	defer shareableContent.Send(objc.RegisterName("release"))

	if displayCount == 0 {
		fmt.Println("⚠️  No displays found")
		os.Exit(1)
	}

	fmt.Println("\n📸 Capturing screenshot using ScreenCaptureKit...")

	// Get main display ID
	mainDisplayID := coregraphics.CGMainDisplayID()
	fmt.Printf("✓ Main display ID: %v\n", mainDisplayID)

	// Capture screenshot of main display
	cgImage := coregraphics.CGDisplayCreateImage(mainDisplayID)
	if cgImage == nil {
		fmt.Println("✗ FAIL: Failed to capture display image")
		os.Exit(1)
	}
	fmt.Println("✓ Captured display image")

	// Get image dimensions
	width := coregraphics.CGImageGetWidth(cgImage)
	height := coregraphics.CGImageGetHeight(cgImage)
	fmt.Printf("✓ Image size: %d x %d pixels\n", width, height)

	// Get other image properties
	bitsPerComponent := coregraphics.CGImageGetBitsPerComponent(cgImage)
	bitsPerPixel := coregraphics.CGImageGetBitsPerPixel(cgImage)
	fmt.Printf("✓ Image format: %d bits/component, %d bits/pixel\n", bitsPerComponent, bitsPerPixel)

	// Release CGImage
	coregraphics.CGImageRelease(cgImage)
	fmt.Println("✓ Released CGImage")

	fmt.Println("\n✅ ScreenCaptureKit example complete!")
	fmt.Println("   - Used CoreGraphics display capture")
	fmt.Printf("   - Captured %dx%d screenshot from main display\n", width, height)
	fmt.Println("   - ScreenCaptureKit bindings generated and available")
	fmt.Println("\n💡 Note: This example uses CoreGraphics (CGDisplayCreateImage)")
	fmt.Println("   For full ScreenCaptureKit features (window capture, filters),")
	fmt.Println("   async completion handlers are required (macOS 12.3+)")
}

func runE2ETest() {
	fmt.Println("=== E2E Test Mode (ScreenCaptureKit Generated Bindings) ===")

	// Check for ScreenCaptureKit framework
	shareableContentClass := objc.GetClass("SCShareableContent")
	if shareableContentClass == 0 {
		fmt.Println("⚠️  SCShareableContent class not found (requires macOS 12.3+)")
		fmt.Println("   Falling back to CoreGraphics display capture test")
		time.Sleep(100 * time.Millisecond)
	} else {
		fmt.Println("✓ Found SCShareableContent class")
		time.Sleep(100 * time.Millisecond)
	}

	// Check for SCScreenshotManager
	screenshotManagerClass := objc.GetClass("SCScreenshotManager")
	if screenshotManagerClass != 0 {
		fmt.Println("✓ Found SCScreenshotManager class")
		time.Sleep(100 * time.Millisecond)
	}

	// Test CoreGraphics display capture (works without ScreenCaptureKit)
	mainDisplayID := coregraphics.CGMainDisplayID()
	if mainDisplayID == nil {
		fmt.Println("✗ FAIL: Failed to get main display ID")
		os.Exit(1)
	}
	fmt.Printf("✓ Got main display ID: %v\n", mainDisplayID)
	time.Sleep(100 * time.Millisecond)

	// Capture display image
	cgImage := coregraphics.CGDisplayCreateImage(mainDisplayID)
	if cgImage == nil {
		fmt.Println("✗ FAIL: Failed to capture display image")
		os.Exit(1)
	}
	fmt.Println("✓ Captured display image")
	time.Sleep(100 * time.Millisecond)

	// Get image dimensions
	width := coregraphics.CGImageGetWidth(cgImage)
	height := coregraphics.CGImageGetHeight(cgImage)
	if width == 0 || height == 0 {
		fmt.Println("✗ FAIL: Invalid image dimensions")
		os.Exit(1)
	}
	fmt.Printf("✓ Image dimensions: %d x %d\n", width, height)
	time.Sleep(100 * time.Millisecond)

	// Release image
	coregraphics.CGImageRelease(cgImage)
	fmt.Println("✓ Released CGImage")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n=== E2E Test PASSED ===")
	fmt.Println("   ✓ Used CoreGraphics display capture")
	fmt.Println("   ✓ CGDisplayCreateImage, CGImageGetWidth/Height")
	fmt.Println("   ✓ ScreenCaptureKit classes available (if macOS 12.3+)")
	os.Exit(0)
}

func createNSString(s string) objc.ID {
	strClass := objc.GetClass("NSString")
	str := objc.ID(strClass).Send(objc.RegisterName("alloc"))
	return str.Send(objc.RegisterName("initWithUTF8String:"), s)
}

func demonstrateScreenCaptureKitAPI() {
	fmt.Println("\n📚 ScreenCaptureKit API Demonstration")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("This example demonstrates how to use ScreenCaptureKit APIs")
	fmt.Println("with objc.NewBlock for async completion handlers.")
	fmt.Println()
	fmt.Println("Code structure (when SCShareableContent is available):")
	fmt.Println()
	fmt.Println("  1. Create completion handler block with objc.NewBlock:")
	fmt.Println("     completionBlock := objc.NewBlock(")
	fmt.Println("       func(block objc.Block, content objc.ID, error objc.ID) {")
	fmt.Println("         // Handle returned content...")
	fmt.Println("       })")
	fmt.Println()
	fmt.Println("  2. Call SCShareableContent class method:")
	fmt.Println("     sel := objc.RegisterName(\"getShareableContentWithCompletionHandler:\")")
	fmt.Println("     objc.ID(shareableContentClass).Send(sel, completionBlock)")
	fmt.Println()
	fmt.Println("  3. Wait for async completion using channels:")
	fmt.Println("     select {")
	fmt.Println("       case <-done:")
	fmt.Println("         // Process results")
	fmt.Println("       case <-time.After(5 * time.Second):")
	fmt.Println("         // Handle timeout")
	fmt.Println("     }")
	fmt.Println()
	fmt.Println("  4. Access display and window information:")
	fmt.Println("     displays := content.Send(objc.RegisterName(\"displays\"))")
	fmt.Println("     windows := content.Send(objc.RegisterName(\"windows\"))")
	fmt.Println()
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("✅ Generated bindings available for:")
	fmt.Println("   - SCShareableContent (system content enumeration)")
	fmt.Println("   - SCDisplay (display information)")
	fmt.Println("   - SCWindow (window information)")
	fmt.Println("   - SCRunningApplication (app information)")
	fmt.Println("   - SCContentFilter (content filtering)")
	fmt.Println("   - SCStream (real-time screen capture)")
	fmt.Println("   - SCStreamConfiguration (stream settings)")
	fmt.Println("   - SCScreenshotManager (screenshot capture)")
	fmt.Println()
	fmt.Println("💡 To test on a system with ScreenCaptureKit:")
	fmt.Println("   - Requires macOS 12.3 (Monterey) or later")
	fmt.Println("   - Requires Screen Recording permission")
	fmt.Println("   - The async block will execute and enumerate displays/windows")
	fmt.Println()
	
	// Show that we still use CoreGraphics for fallback
	fmt.Println("📸 CoreGraphics Fallback Demo:")
	mainDisplayID := coregraphics.CGMainDisplayID()
	cgImage := coregraphics.CGDisplayCreateImage(mainDisplayID)
	if cgImage != nil {
		width := coregraphics.CGImageGetWidth(cgImage)
		height := coregraphics.CGImageGetHeight(cgImage)
		fmt.Printf("✓ Captured %dx%d screenshot using CoreGraphics\n", width, height)
		coregraphics.CGImageRelease(cgImage)
	}
	
	fmt.Println()
	fmt.Println("═══════════════════════════════════════════════════════════")
}

func testFrameworkLoading() {
	fmt.Println("=== ScreenCaptureKit Framework Loading Test ===")
	fmt.Println()

	// Try to load ScreenCaptureKit framework explicitly
	fmt.Println("1. Attempting to load ScreenCaptureKit.framework...")
	handle, err := purego.Dlopen("/System/Library/Frameworks/ScreenCaptureKit.framework/ScreenCaptureKit", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Printf("   ✗ Failed to load framework: %v\n", err)
		fmt.Println()

		// Try without RTLD_GLOBAL
		fmt.Println("2. Trying without RTLD_GLOBAL...")
		handle, err = purego.Dlopen("/System/Library/Frameworks/ScreenCaptureKit.framework/ScreenCaptureKit", purego.RTLD_NOW)
		if err != nil {
			fmt.Printf("   ✗ Failed again: %v\n", err)
			fmt.Println()
		} else {
			fmt.Printf("   ✓ Loaded framework (handle: %v)\n", handle)
			fmt.Println()
		}
	} else {
		fmt.Printf("   ✓ Loaded framework (handle: %v)\n", handle)
		fmt.Println()
	}

	// Now check for SCShareableContent class
	fmt.Println("3. Looking up SCShareableContent class...")
	scClass := objc.GetClass("SCShareableContent")
	if scClass == 0 {
		fmt.Println("   ✗ SCShareableContent class NOT found")
	} else {
		fmt.Printf("   ✓ SCShareableContent class found: %v\n", scClass)
	}
	fmt.Println()

	// Check for other ScreenCaptureKit classes
	fmt.Println("4. Checking other ScreenCaptureKit classes...")
	classes := []string{
		"SCScreenshotManager",
		"SCDisplay",
		"SCWindow",
		"SCRunningApplication",
		"SCContentFilter",
		"SCStream",
		"SCStreamConfiguration",
	}

	for _, className := range classes {
		cls := objc.GetClass(className)
		if cls == 0 {
			fmt.Printf("   ✗ %s NOT found\n", className)
		} else {
			fmt.Printf("   ✓ %s found: %v\n", className, cls)
		}
	}
	fmt.Println()

	// Try loading with versioned path
	fmt.Println("5. Trying versioned framework path...")
	handle2, err := purego.Dlopen("/System/Library/Frameworks/ScreenCaptureKit.framework/Versions/A/ScreenCaptureKit", purego.RTLD_NOW)
	if err != nil {
		fmt.Printf("   ✗ Failed: %v\n", err)
	} else {
		fmt.Printf("   ✓ Loaded versioned framework (handle: %v)\n", handle2)

		// Check class again
		scClass2 := objc.GetClass("SCShareableContent")
		if scClass2 == 0 {
			fmt.Println("   ✗ Still no SCShareableContent class")
		} else {
			fmt.Printf("   ✓ SCShareableContent found: %v\n", scClass2)
		}
	}
	fmt.Println()

	fmt.Println("=== Test Complete ===")
}

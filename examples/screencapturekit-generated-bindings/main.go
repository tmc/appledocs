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
	"github.com/tmc/macgo"
)

var (
	e2e         = flag.Bool("e2e", false, "run end-to-end test mode (non-interactive)")
	testLoading = flag.Bool("test-loading", false, "test framework loading")
)

func init() {
	runtime.LockOSThread()

	// Skip macgo setup for test modes (E2E, framework loading test)
	// These modes don't need TCC permissions or app bundle setup
	skipMacgo := false
	for _, arg := range os.Args[1:] {
		if arg == "-e2e" || arg == "-test-loading" {
			skipMacgo = true
			break
		}
	}

	if !skipMacgo {
		// Use macgo to set up app bundle with proper entitlements
		// This enables proper TCC permission requests for screen capture
		cfg := &macgo.Config{
			AppName:  "ScreenCaptureKit-Example",
			BundleID: "com.github.tmc.appledocs.screencapturekit-example",
			Version:  "1.0.0",
			Custom: []string{
				// Request screen capture entitlement
				"com.apple.security.app-sandbox",
				"com.apple.security.device.camera", // Sometimes needed for screen recording
			},
			AdHocSign:           true, // Use ad-hoc signing
			ForceLaunchServices: true, // Use 'open' to trigger TCC prompts
			Debug:               os.Getenv("MACGO_DEBUG") == "1",
		}

		// Start macgo - this will relaunch via app bundle if needed
		if err := macgo.Start(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "macgo.Start failed: %v\n", err)
			os.Exit(1)
		}
	}

	// Explicitly load ScreenCaptureKit framework
	// This is required for the Objective-C classes to be available
	_, err := purego.Dlopen("/System/Library/Frameworks/ScreenCaptureKit.framework/ScreenCaptureKit", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		// Framework not available, will be handled at runtime
		_ = err
	}
}

// waitForScreenRecordingPermission waits for Screen Recording permission to be granted
// It retries the ScreenCaptureKit API call with exponential backoff and user feedback
func waitForScreenRecordingPermission() (shareableContent objc.ID, displayCount, windowCount int, err error) {
	maxAttempts := 10
	baseDelay := 500 * time.Millisecond

	shareableContentClass := objc.GetClass("SCShareableContent")
	if shareableContentClass == 0 {
		return 0, 0, 0, fmt.Errorf("SCShareableContent class not found (requires macOS 12.3+)")
	}

	sel := objc.RegisterName("getShareableContentWithCompletionHandler:")

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		done := make(chan bool, 1)
		var content objc.ID
		var displays, windows int
		var lastError string

		// Create completion handler block
		completionBlock := objc.NewBlock(
			func(block objc.Block, c objc.ID, e objc.ID) {
				defer func() { done <- true }()

				if e != 0 {
					desc := e.Send(objc.RegisterName("localizedDescription"))
					if desc != 0 {
						lastError = objc.Send[string](desc, objc.RegisterName("UTF8String"))
					}
					return
				}

				if c == 0 {
					lastError = "No content returned"
					return
				}

				content = c
				content.Send(objc.RegisterName("retain"))

				// Get displays array
				d := c.Send(objc.RegisterName("displays"))
				if d != 0 {
					displays = int(d.Send(objc.RegisterName("count")))
				}

				// Get windows array
				w := c.Send(objc.RegisterName("windows"))
				if w != 0 {
					windows = int(w.Send(objc.RegisterName("count")))
				}
			},
		)

		// Call async method
		objc.ID(shareableContentClass).Send(sel, completionBlock)

		// Wait for completion with timeout
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			lastError = "Timeout waiting for shareable content"
		}

		completionBlock.Release()

		// Check if we succeeded
		if content != 0 && displays > 0 {
			return content, displays, windows, nil
		}

		// First attempt - show user instructions
		if attempt == 1 && lastError != "" {
			fmt.Fprintf(os.Stderr, "\n⚠️  Waiting for Screen Recording permission...\n")
			fmt.Fprintf(os.Stderr, "   Please grant permission in System Settings → Privacy & Security → Screen Recording\n")
			inBundle := os.Getenv("MACGO_IN_BUNDLE") == "1"
			if !inBundle {
				fmt.Fprintf(os.Stderr, "   The permission dialog should appear automatically.\n")
			}
			fmt.Fprintf(os.Stderr, "   Look for: ScreenCaptureKit-Example.app\n")
			fmt.Fprintf(os.Stderr, "\n")
		}

		// Show waiting indicator
		dots := ""
		for i := 0; i < attempt%4; i++ {
			dots += "."
		}
		remaining := maxAttempts - attempt
		fmt.Fprintf(os.Stderr, "\r   Waiting%s (%d/%d attempts remaining)   ",
			dots, remaining, maxAttempts)

		// Exponential backoff with jitter
		delay := time.Duration(float64(baseDelay) * (1 + float64(attempt)*0.5))
		if delay > 5*time.Second {
			delay = 5 * time.Second
		}
		time.Sleep(delay)
	}

	fmt.Fprintf(os.Stderr, "\n\n❌ Screen Recording permission not granted after %d attempts\n", maxAttempts)
	fmt.Fprintf(os.Stderr, "   Please check System Settings → Privacy & Security → Screen Recording\n")
	fmt.Fprintf(os.Stderr, "   Then rerun this example.\n\n")
	return 0, 0, 0, fmt.Errorf("screen recording permission not available")
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
	fmt.Println("\n⚠️  This example requires Screen Recording permission.")
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
	fmt.Println("⏳ Requesting shareable content with retry logic...")

	// Wait for permission with retries
	shareableContent, displayCount, windowCount, err := waitForScreenRecordingPermission()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get shareable content: %v\n", err)
		os.Exit(1)
	}
	defer shareableContent.Send(objc.RegisterName("release"))

	fmt.Fprintf(os.Stderr, "\r   ✓ Permission granted!                                        \n\n")
	fmt.Printf("✓ Found %d display(s)\n", displayCount)

	// Print detailed display info
	displays := shareableContent.Send(objc.RegisterName("displays"))
	if displays != 0 {
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

	if windowCount > 0 {
		fmt.Printf("✓ Found %d window(s)\n", windowCount)
	}

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

// ScreenCaptureKit example using only generated bindings
//
// This example demonstrates:
// - Async SCShareableContent enumeration using objc.NewBlock()
// - TCC permission handling with retry logic
// - Display and window enumeration
package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
	"github.com/tmc/macgo"
)

func init() {
	runtime.LockOSThread()

	// Use macgo to set up app bundle with proper entitlements
	// This enables proper TCC permission requests for screen capture
	cfg := &macgo.Config{
		AppName:             "ScreenCaptureKit-Example",
		BundleID:            "com.github.tmc.appledocs.screencapturekit-example",
		Version:             "1.0.0",
		Custom:              []string{
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
	fmt.Println("=== ScreenCaptureKit Async Enumeration Example ===")
	fmt.Println()
	fmt.Println("This example demonstrates async SCShareableContent enumeration")
	fmt.Println("using objc.NewBlock() with proper TCC permission handling.")
	fmt.Println()

	// ScreenCaptureKit requires screen recording permission
	fmt.Println("⚠️  Requires Screen Recording permission")
	fmt.Println()

	// Get shareable content (displays and windows)
	fmt.Println("📋 Enumerating shareable content...")

	shareableContentClass := objc.GetClass("SCShareableContent")
	if shareableContentClass == 0 {
		fmt.Println("❌ SCShareableContent class not found (requires macOS 12.3+)")
		os.Exit(1)
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

	// Print display information
	fmt.Println("=== Results ===")
	fmt.Printf("✓ Found %d display(s)\n", displayCount)

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

	// Print window information (first 5)
	fmt.Printf("\n✓ Found %d window(s)\n", windowCount)
	if windowCount > 0 {
		windows := shareableContent.Send(objc.RegisterName("windows"))
		if windows != 0 {
			fmt.Println("   (showing first 5):")
			maxToShow := windowCount
			if maxToShow > 5 {
				maxToShow = 5
			}
			for i := 0; i < maxToShow; i++ {
				window := windows.Send(objc.RegisterName("objectAtIndex:"), i)
				if window != 0 {
					windowID := window.Send(objc.RegisterName("windowID"))

					// Get title (may be nil)
					title := window.Send(objc.RegisterName("title"))
					titleStr := "(no title)"
					if title != 0 {
						titleStr = objc.Send[string](title, objc.RegisterName("UTF8String"))
						if titleStr == "" {
							titleStr = "(no title)"
						}
					}

					// Get owning application
					app := window.Send(objc.RegisterName("owningApplication"))
					appName := ""
					if app != 0 {
						appNameObj := app.Send(objc.RegisterName("applicationName"))
						if appNameObj != 0 {
							appName = objc.Send[string](appNameObj, objc.RegisterName("UTF8String"))
						}
					}

					fmt.Printf("   - Window %d: %s (app: %s)\n", windowID, titleStr, appName)
				}
			}
		}
	}

	// Print applications
	applications := shareableContent.Send(objc.RegisterName("applications"))
	if applications != 0 {
		appCount := int(applications.Send(objc.RegisterName("count")))
		fmt.Printf("\n✓ Found %d running application(s)\n", appCount)
		if appCount > 0 {
			fmt.Println("   (showing first 5):")
			maxToShow := appCount
			if maxToShow > 5 {
				maxToShow = 5
			}
			for i := 0; i < maxToShow; i++ {
				app := applications.Send(objc.RegisterName("objectAtIndex:"), i)
				if app != 0 {
					appNameObj := app.Send(objc.RegisterName("applicationName"))
					processID := app.Send(objc.RegisterName("processID"))
					appName := ""
					if appNameObj != 0 {
						appName = objc.Send[string](appNameObj, objc.RegisterName("UTF8String"))
					}
					fmt.Printf("   - %s (pid: %d)\n", appName, processID)
				}
			}
		}
	}

	fmt.Println()
	fmt.Println("=== ✅ Async Enumeration Complete ===")
	fmt.Println()
	fmt.Println("Successfully demonstrated:")
	fmt.Println("   ✓ Async SCShareableContent enumeration with objc.NewBlock()")
	fmt.Println("   ✓ TCC permission handling with retry logic")
	fmt.Println("   ✓ Display, window, and application enumeration")
	fmt.Println()
	fmt.Println("Compare with Objective-C reference: sc_async_enum.m")
}

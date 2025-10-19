// ScreenCaptureKit example using only generated bindings
//
// This example demonstrates:
// - Async SCShareableContent enumeration using objc.NewBlock()
// - TCC permission handling with retry logic
// - Display and window enumeration
// - Screen recording with SCStream
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
	"github.com/tmc/macgo"
)

var (
	listAll    = flag.Bool("all", false, "list all windows (not just first 5)")
	record     = flag.Bool("record", false, "record screen content using SCStream")
	duration   = flag.Duration("duration", 5*time.Second, "recording duration")
	displayNum = flag.Int("display", 0, "display number to record (0-based)")
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
	flag.Parse()

	fmt.Println("=== ScreenCaptureKit Async Enumeration Example ===")
	fmt.Println()
	fmt.Println("This example demonstrates async SCShareableContent enumeration")
	fmt.Println("using objc.NewBlock() with proper TCC permission handling.")
	if *record {
		fmt.Println("and screen recording with SCStream.")
	}
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

	// Print window information
	fmt.Printf("\n✓ Found %d window(s)\n", windowCount)
	if windowCount > 0 {
		windows := shareableContent.Send(objc.RegisterName("windows"))
		if windows != 0 {
			maxToShow := windowCount
			if !*listAll && maxToShow > 5 {
				fmt.Println("   (showing first 5, use -all to show all):")
				maxToShow = 5
			} else if *listAll {
				fmt.Println("   (showing all):")
			} else {
				fmt.Println("   (showing all):")
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

	// If -record flag is set, start screen recording
	if *record {
		if displayCount == 0 {
			fmt.Println("❌ No displays available for recording")
			os.Exit(1)
		}

		if *displayNum >= displayCount {
			fmt.Printf("❌ Display %d not found (available: 0-%d)\n", *displayNum, displayCount-1)
			os.Exit(1)
		}

		fmt.Println("=== Starting Screen Recording ===")
		fmt.Println()

		// Get the selected display
		display := displays.Send(objc.RegisterName("objectAtIndex:"), *displayNum)
		if display == 0 {
			fmt.Println("❌ Failed to get display")
			os.Exit(1)
		}

		displayID := display.Send(objc.RegisterName("displayID"))
		width := display.Send(objc.RegisterName("width"))
		height := display.Send(objc.RegisterName("height"))
		fmt.Printf("Recording display %d: ID=%d, %dx%d\n", *displayNum, displayID, width, height)
		fmt.Printf("Duration: %v\n", *duration)
		fmt.Println()

		if err := recordScreen(display, *duration); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Recording failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Println()
		fmt.Println("✅ Recording complete!")
	} else {
		fmt.Println("💡 Use -record flag to record screen content with SCStream")
	}

	fmt.Println()
	fmt.Println("Compare with Objective-C reference: sc_async_enum.m")
}

// recordScreen uses SCStream to record screen content
func recordScreen(display objc.ID, duration time.Duration) error {
	fmt.Println("⏺  Setting up SCStream...")

	// Get display width and height
	width := int(display.Send(objc.RegisterName("width")))
	height := int(display.Send(objc.RegisterName("height")))

	// Create SCStreamConfiguration
	configClass := objc.GetClass("SCStreamConfiguration")
	if configClass == 0 {
		return fmt.Errorf("SCStreamConfiguration class not found")
	}
	config := objc.ID(configClass).Send(objc.RegisterName("alloc"))
	config = config.Send(objc.RegisterName("init"))
	defer config.Send(objc.RegisterName("release"))

	// Configure stream settings
	config.Send(objc.RegisterName("setPixelFormat:"), uint32(0x42475241)) // 'BGRA'
	config.Send(objc.RegisterName("setQueueDepth:"), 5)
	config.Send(objc.RegisterName("setWidth:"), width)
	config.Send(objc.RegisterName("setHeight:"), height)
	config.Send(objc.RegisterName("setShowsCursor:"), true)

	// Create content filter for the display
	filterClass := objc.GetClass("SCContentFilter")
	if filterClass == 0 {
		return fmt.Errorf("SCContentFilter class not found")
	}

	// Create empty arrays for excluding/excepting
	arrayClass := objc.GetClass("NSArray")
	emptyArray := objc.ID(arrayClass).Send(objc.RegisterName("array"))

	// initWithDisplay:excludingApplications:exceptingWindows:
	initSel := objc.RegisterName("initWithDisplay:excludingApplications:exceptingWindows:")
	filter := objc.ID(filterClass).Send(objc.RegisterName("alloc"))
	filter = filter.Send(initSel, display, emptyArray, emptyArray)
	defer filter.Send(objc.RegisterName("release"))

	// Create SCStream
	streamClass := objc.GetClass("SCStream")
	if streamClass == 0 {
		return fmt.Errorf("SCStream class not found")
	}

	// Create delegate instance
	frameCount := 0
	delegate, err := createStreamOutputDelegate(&frameCount)
	if err != nil {
		return fmt.Errorf("failed to create delegate: %w", err)
	}
	defer delegate.Send(objc.RegisterName("release"))

	fmt.Println("✓ Created SCStreamOutput delegate using objc.RegisterClass!")
	fmt.Println()

	// Create SCStream with filter and config
	initStreamSel := objc.RegisterName("initWithFilter:configuration:delegate:")
	stream := objc.ID(streamClass).Send(objc.RegisterName("alloc"))
	stream = stream.Send(initStreamSel, filter, config, objc.ID(0)) // nil delegate for now
	if stream == 0 {
		return fmt.Errorf("failed to create SCStream")
	}
	defer stream.Send(objc.RegisterName("release"))

	// Create dispatch queue for stream output
	queueClass := objc.GetClass("OS_dispatch_queue")
	if queueClass == 0 {
		// Try to create queue using dispatch_queue_create
		return fmt.Errorf("dispatch queue not available - need to use dispatch_queue_create")
	}

	// Add stream output with our delegate
	addOutputSel := objc.RegisterName("addStreamOutput:type:sampleHandlerQueue:error:")

	// For now, use nil queue (main queue)
	var errorPtr objc.ID
	success := stream.Send(addOutputSel, delegate, 0, objc.ID(0), &errorPtr)
	if success == 0 || errorPtr != 0 {
		if errorPtr != 0 {
			desc := errorPtr.Send(objc.RegisterName("localizedDescription"))
			if desc != 0 {
				errMsg := objc.Send[string](desc, objc.RegisterName("UTF8String"))
				return fmt.Errorf("failed to add stream output: %s", errMsg)
			}
		}
		return fmt.Errorf("failed to add stream output")
	}

	fmt.Println("✓ Added delegate to SCStream")
	fmt.Println()

	// Start capture
	fmt.Println("⏺  Starting capture...")
	startDone := make(chan error, 1)
	startBlock := objc.NewBlock(func(block objc.Block, err objc.ID) {
		if err != 0 {
			desc := err.Send(objc.RegisterName("localizedDescription"))
			if desc != 0 {
				errMsg := objc.Send[string](desc, objc.RegisterName("UTF8String"))
				startDone <- fmt.Errorf("%s", errMsg)
				return
			}
		}
		startDone <- nil
	})
	defer startBlock.Release()

	stream.Send(objc.RegisterName("startCaptureWithCompletionHandler:"), startBlock)

	if err := <-startDone; err != nil {
		return fmt.Errorf("failed to start capture: %w", err)
	}

	fmt.Printf("✓ Capture started! Recording for %v...\n", duration)
	fmt.Println("   (Frames will be logged every 30 frames)")
	fmt.Println()

	// Record for the specified duration
	time.Sleep(duration)

	// Stop capture
	fmt.Println("\n⏹  Stopping capture...")
	stopDone := make(chan error, 1)
	stopBlock := objc.NewBlock(func(block objc.Block, err objc.ID) {
		if err != 0 {
			desc := err.Send(objc.RegisterName("localizedDescription"))
			if desc != 0 {
				errMsg := objc.Send[string](desc, objc.RegisterName("UTF8String"))
				stopDone <- fmt.Errorf("%s", errMsg)
				return
			}
		}
		stopDone <- nil
	})
	defer stopBlock.Release()

	stream.Send(objc.RegisterName("stopCaptureWithCompletionHandler:"), stopBlock)

	if err := <-stopDone; err != nil {
		return fmt.Errorf("failed to stop capture: %w", err)
	}

	fmt.Printf("\n✓ Capture stopped!\n")
	fmt.Printf("   Total frames received: %d\n", frameCount)
	fps := float64(frameCount) / duration.Seconds()
	fmt.Printf("   Frame rate: %.1f FPS\n", fps)
	fmt.Printf("   Resolution: %dx%d\n", width, height)

	return nil
}

type streamOutputHandler struct {
	frameCount *int
}

// createStreamOutputDelegate creates a custom Objective-C class that implements SCStreamOutput protocol
func createStreamOutputDelegate(frameCount *int) (objc.ID, error) {
	// Get NSObject as our superclass
	nsObjectClass := objc.GetClass("NSObject")
	if nsObjectClass == 0 {
		return 0, fmt.Errorf("NSObject class not found")
	}

	// Note: SCStreamOutput protocol may not be formally registered at runtime,
	// but Objective-C uses duck typing - as long as we implement the right methods,
	// the class will work as a delegate. We don't need to declare protocol conformance.

	// Create the delegate method
	// Signature: - (void)stream:(SCStream *)stream didOutputSampleBuffer:(CMSampleBufferRef)sampleBuffer ofType:(SCStreamOutputType)type
	streamDidOutputSampleBuffer := func(self objc.ID, cmd objc.SEL, stream objc.ID, sampleBuffer uintptr, outputType int) {
		// Only process screen output
		if outputType != 0 { // SCStreamOutputTypeScreen = 0
			return
		}

		*frameCount++
		if *frameCount%30 == 0 {
			fmt.Printf("\r   📹 Received frame %d", *frameCount)
		}

		// TODO: Extract CVPixelBuffer and save as PNG
		// This requires CoreMedia and CoreVideo bindings
	}

	// Register the class without protocol (duck typing will make it work)
	className := fmt.Sprintf("GoStreamOutputDelegate_%d", time.Now().UnixNano())
	delegateClass, err := objc.RegisterClass(
		className,
		nsObjectClass,
		nil, // no protocols - duck typing will handle it
		nil, // no ivars
		[]objc.MethodDef{
			{
				Cmd: objc.RegisterName("stream:didOutputSampleBuffer:ofType:"),
				Fn:  streamDidOutputSampleBuffer,
			},
		},
	)
	if err != nil {
		return 0, fmt.Errorf("failed to register delegate class: %w", err)
	}

	// Create an instance
	delegate := objc.ID(delegateClass).Send(objc.RegisterName("alloc"))
	delegate = delegate.Send(objc.RegisterName("init"))

	return delegate, nil
}

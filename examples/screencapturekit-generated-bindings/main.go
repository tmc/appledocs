// ScreenCaptureKit example using generated bindings with property accessor support
//
// This example demonstrates:
// - Async SCShareableContent enumeration using objc.NewBlock()
// - TCC permission handling with retry logic
// - Display and window enumeration
// - Screen recording with SCStream
// - Typed ScreenCaptureKit bindings (SCShareableContent, SCDisplay, SCWindow, etc.)
// - Using generated property accessors (e.g., shareableContent.Displays())
//
// NOTE: Property accessors are generated and return unsafe.Pointer for NSArray types,
// which are then cast to objc.ID for array operations. This is the correct approach
// since NSArray is a dynamic Objective-C collection type.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/coreimage"
	"github.com/tmc/appledocs/generated/coremedia"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/imageio"
	"github.com/tmc/appledocs/generated/screencapturekit"
	"github.com/tmc/macgo"
)

var (
	listAll    = flag.Bool("all", false, "list all windows (not just first 5)")
	record     = flag.Bool("record", false, "record screen content using SCStream")
	duration   = flag.Duration("duration", 5*time.Second, "recording duration")
	displayNum = flag.Int("display", 0, "display number to record (0-based)")
	saveFrames = flag.Bool("save", false, "save frames as PNG files")
)

// Note: All CGImageDestination functions are now in generated ImageIO bindings!

func init() {
	runtime.LockOSThread()

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

	// Explicitly load ScreenCaptureKit framework
	// This is required for the Objective-C classes to be available
	_, err := purego.Dlopen("/System/Library/Frameworks/ScreenCaptureKit.framework/ScreenCaptureKit", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		// Framework not available, will be handled at runtime
		_ = err
	}

	// Note: ImageIO framework is loaded automatically by generated bindings
	// All CGImageDestination functions are available via imageio package
}

// waitForScreenRecordingPermission waits for Screen Recording permission to be granted
// It retries the ScreenCaptureKit API call with exponential backoff and user feedback
// Returns typed SCShareableContent once F7D7's property accessor work is complete
func waitForScreenRecordingPermission() (content screencapturekit.SCShareableContent, displayCount, windowCount int, err error) {
	maxAttempts := 10
	baseDelay := 500 * time.Millisecond

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		done := make(chan bool, 1)
		var shareableContent screencapturekit.SCShareableContent
		var displays, windows int
		var lastError string

		// Create completion handler block
		// NOTE: Once property accessors are generated, we can use:
		//   displays = len(shareableContent.Displays())
		//   windows = len(shareableContent.Windows())
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

				// Convert objc.ID to typed SCShareableContent
				shareableContent = screencapturekit.SCShareableContentFrom(unsafe.Pointer(c))
				// Retain to prevent deallocation
				c.Send(objc.RegisterName("retain"))

				// Use generated property accessors with helper functions
				displays = len(shareableContent.Displays())
				windows = len(shareableContent.Windows())
			},
		)

		// Call async class method using manual objc.Send
		// Note: GetShareableContentWithCompletionHandler is not in generated bindings
		//       (only GetCurrentProcessShareableContentWithCompletionHandler exists)
		shareableContentClass := objc.GetClass("SCShareableContent")
		if shareableContentClass == 0 {
			return screencapturekit.SCShareableContent{}, 0, 0, fmt.Errorf("SCShareableContent class not found (requires macOS 12.3+)")
		}
		sel := objc.RegisterName("getShareableContentWithCompletionHandler:")
		objc.ID(shareableContentClass).Send(sel, completionBlock)

		// Wait for completion with timeout
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			lastError = "Timeout waiting for shareable content"
		}

		completionBlock.Release()

		// Check if we succeeded
		if shareableContent.ID != 0 && displays > 0 {
			return shareableContent, displays, windows, nil
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
	return screencapturekit.SCShareableContent{}, 0, 0, fmt.Errorf("screen recording permission not available")
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
	fmt.Println("⏳ Requesting shareable content with retry logic...")

	// Wait for permission with retries - now returns typed SCShareableContent
	shareableContent, displayCount, windowCount, err := waitForScreenRecordingPermission()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get shareable content: %v\n", err)
		os.Exit(1)
	}
	// Release using objc.Send (NSObject methods not yet in generated bindings)
	defer shareableContent.ID.Send(objc.RegisterName("release"))

	fmt.Fprintf(os.Stderr, "\r   ✓ Permission granted!                                        \n\n")

	// Print display information using generated property accessors
	fmt.Println("=== Results ===")
	fmt.Printf("✓ Found %d display(s)\n", displayCount)

	displays := shareableContent.Displays()
	for i, display := range displays {
		// DisplayID() returns unsafe.Pointer, Width/Height not yet generated
		id := display.DisplayID()
		width := display.ID.Send(objc.RegisterName("width"))
		height := display.ID.Send(objc.RegisterName("height"))
		fmt.Printf("   Display %d: ID=%v, %dx%d\n", i, id, width, height)
	}

	// Print window information
	fmt.Printf("\n✓ Found %d window(s)\n", windowCount)
	if windowCount > 0 {
		windows := shareableContent.Windows()
		maxToShow := len(windows)
		if !*listAll && maxToShow > 5 {
			fmt.Println("   (showing first 5, use -all to show all):")
			maxToShow = 5
		} else {
			fmt.Println("   (showing all):")
		}
		for _, window := range windows[:maxToShow] {
			// Property accessors not yet generated for SCWindow
			id := window.ID.Send(objc.RegisterName("windowID"))
			titleObj := window.ID.Send(objc.RegisterName("title"))
			titleStr := nsStringToGo(titleObj)
			if titleStr == "" {
				titleStr = "(no title)"
			}

			appObj := window.ID.Send(objc.RegisterName("owningApplication"))
			appNameStr := ""
			if appObj != 0 {
				appNameObj := appObj.Send(objc.RegisterName("applicationName"))
				appNameStr = nsStringToGo(appNameObj)
			}

			fmt.Printf("   - Window %d: %s (app: %s)\n", id, titleStr, appNameStr)
		}
	}

	// Print applications
	applications := shareableContent.Applications()
	fmt.Printf("\n✓ Found %d running application(s)\n", len(applications))
	if len(applications) > 0 {
		fmt.Println("   (showing first 5):")
		maxToShow := len(applications)
		if maxToShow > 5 {
			maxToShow = 5
		}
		for _, app := range applications[:maxToShow] {
			// Property accessors not yet generated for SCRunningApplication
			appNameObj := app.ID.Send(objc.RegisterName("applicationName"))
			processID := app.ID.Send(objc.RegisterName("processID"))
			appName := nsStringToGo(appNameObj)
			fmt.Printf("   - %s (pid: %d)\n", appName, processID)
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

		// Get the selected display from the slice
		displays := shareableContent.Displays()
		if *displayNum >= len(displays) {
			fmt.Println("❌ Invalid display index")
			os.Exit(1)
		}

		display := displays[*displayNum]
		displayID := display.DisplayID()
		width := display.ID.Send(objc.RegisterName("width"))
		height := display.ID.Send(objc.RegisterName("height"))
		fmt.Printf("Recording display %d: ID=%v, %dx%d\n", *displayNum, displayID, width, height)
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
// Now accepts typed SCDisplay
func recordScreen(display screencapturekit.SCDisplay, duration time.Duration) error {
	fmt.Println("⏺  Setting up SCStream...")

	// Get display width and height
	// Using property accessors (return unsafe.Pointer for objects):
	//   width := display.Width()
	//   height := display.Height()
	width := int(display.ID.Send(objc.RegisterName("width")))
	height := int(display.ID.Send(objc.RegisterName("height")))

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
	// Init methods are generated; keeping manual approach for clarity:
	//   emptyApps := []screencapturekit.SCRunningApplication{}
	//   emptyWindows := []screencapturekit.SCWindow{}
	//   filter := screencapturekit.NewSCContentFilterWithDisplayExcludingApplicationsExceptingWindows(
	//       display, emptyApps, emptyWindows)

	filterClass := objc.GetClass("SCContentFilter")
	if filterClass == 0 {
		return fmt.Errorf("SCContentFilter class not found")
	}

	// Create empty arrays for excluding/excepting
	arrayClass := objc.GetClass("NSArray")
	emptyArray := objc.ID(arrayClass).Send(objc.RegisterName("array"))

	// initWithDisplay:excludingApplications:exceptingWindows:
	initSel := objc.RegisterName("initWithDisplay:excludingApplications:exceptingWindows:")
	filterID := objc.ID(filterClass).Send(objc.RegisterName("alloc"))
	filterID = filterID.Send(initSel, display.ID, emptyArray, emptyArray)
	defer filterID.Send(objc.RegisterName("release"))

	// Convert to typed SCContentFilter (for future use)
	_ = screencapturekit.SCContentFilterFrom(unsafe.Pointer(filterID))

	// Create SCStream
	streamClass := objc.GetClass("SCStream")
	if streamClass == 0 {
		return fmt.Errorf("SCStream class not found")
	}

	// Create delegate using helper API
	handler := &FrameHandler{
		shouldSave: *saveFrames,
	}

	// Setup output directory if saving frames
	if *saveFrames {
		handler.outputDir = filepath.Join(os.TempDir(), "screencapture_frames")
		if err := os.MkdirAll(handler.outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
		fmt.Printf("📁 Output directory: %s\n", handler.outputDir)
	}

	// Use the helper API to create delegate
	delegate := screencapturekit.NewSCStreamOutputDelegate(handler)
	defer delegate.Send(objc.RegisterName("release"))

	fmt.Println("✓ Created SCStreamOutput delegate using helper API!")
	if *saveFrames {
		fmt.Println("   💾 Frame saving enabled (every 30th frame)")
	}
	fmt.Println()

	// Create SCStream with filter and config
	// Init methods are generated; keeping manual approach for clarity:
	//   stream := screencapturekit.NewSCStreamWithFilterConfigurationDelegate(
	//       filter, config, nil)

	initStreamSel := objc.RegisterName("initWithFilter:configuration:delegate:")
	streamID := objc.ID(streamClass).Send(objc.RegisterName("alloc"))
	streamID = streamID.Send(initStreamSel, filterID, config, objc.ID(0)) // nil delegate for now
	if streamID == 0 {
		return fmt.Errorf("failed to create SCStream")
	}
	defer streamID.Send(objc.RegisterName("release"))

	// Convert to typed SCStream (for future use)
	_ = screencapturekit.SCStreamFrom(unsafe.Pointer(streamID))

	// Create dispatch queue for stream output
	queueClass := objc.GetClass("OS_dispatch_queue")
	if queueClass == 0 {
		// Try to create queue using dispatch_queue_create
		return fmt.Errorf("dispatch queue not available - need to use dispatch_queue_create")
	}

	// Add stream output with our delegate
	addOutputSel := objc.RegisterName("addStreamOutput:type:sampleHandlerQueue:error:")

	// Add stream output with our delegate
	// Methods exist; using manual objc.Send for flexibility:
	//   err := stream.AddStreamOutputTypeSampleHandlerQueueError(delegate, 0, nil, &errorPtr)

	// For now, use nil queue (main queue)
	var errorPtr objc.ID
	success := streamID.Send(addOutputSel, delegate, 0, objc.ID(0), &errorPtr)
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

	// Methods exist; using manual objc.Send for flexibility:
	//   stream.StartCaptureWithCompletionHandler(startBlock)
	streamID.Send(objc.RegisterName("startCaptureWithCompletionHandler:"), startBlock)

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

	// Methods exist; using manual objc.Send for flexibility:
	//   stream.StopCaptureWithCompletionHandler(stopBlock)
	streamID.Send(objc.RegisterName("stopCaptureWithCompletionHandler:"), stopBlock)

	if err := <-stopDone; err != nil {
		return fmt.Errorf("failed to stop capture: %w", err)
	}

	fmt.Printf("\n✓ Capture stopped!\n")
	fmt.Printf("   Total frames received: %d\n", handler.frameCount)
	fps := float64(handler.frameCount) / duration.Seconds()
	fmt.Printf("   Frame rate: %.1f FPS\n", fps)
	fmt.Printf("   Resolution: %dx%d\n", width, height)

	if *saveFrames && handler.outputDir != "" {
		savedCount := handler.frameCount / 30
		if handler.frameCount%30 == 0 {
			savedCount = handler.frameCount / 30
		}
		fmt.Printf("\n💾 Saved frames:\n")
		fmt.Printf("   Directory: %s\n", handler.outputDir)
		fmt.Printf("   Total saved: %d PNG files (every 30th frame)\n", savedCount)
	}

	return nil
}

// FrameHandler implements screencapturekit.SCStreamOutputHandler to handle incoming frames
type FrameHandler struct {
	frameCount int
	shouldSave bool
	outputDir  string
}

// StreamDidOutputSampleBuffer implements the SCStreamOutputHandler interface
func (h *FrameHandler) StreamDidOutputSampleBuffer(stream screencapturekit.SCStream, sampleBuffer uintptr, outputType int) {
	// Only process screen output
	if outputType != 0 { // SCStreamOutputTypeScreen = 0
		return
	}

	h.frameCount++
	if h.frameCount%30 == 0 {
		fmt.Printf("\r   📹 Received frame %d", h.frameCount)
	}

	if !h.shouldSave {
		return
	}

	// Save every 30th frame
	if h.frameCount%30 != 0 {
		return
	}

	// Get pixel buffer from sample buffer using generated CoreMedia binding
	pixelBuffer := coremedia.CMSampleBufferGetImageBuffer(unsafe.Pointer(sampleBuffer))
	if pixelBuffer == nil {
		fmt.Fprintf(os.Stderr, "\n⚠️  No pixel buffer in sample\n")
		return
	}

	// Create CIImage from pixel buffer using generated bindings
	ciImage := coreimage.NewImageWithCVPixelBuffer(unsafe.Pointer(pixelBuffer))
	if ciImage.ID == 0 {
		fmt.Fprintf(os.Stderr, "\n⚠️  Failed to create CIImage\n")
		return
	}

	// Create CIContext using generated bindings
	context := coreimage.NewContext()
	if context.ID == 0 {
		fmt.Fprintf(os.Stderr, "\n⚠️  Failed to create CIContext\n")
		return
	}

	// Get image extent
	extent := ciImage.Send(objc.RegisterName("extent"))

	// Create CGImage from CIImage using generated bindings
	cgImage := context.CreateCGImageFromRect(unsafe.Pointer(ciImage.ID), unsafe.Pointer(extent))
	if cgImage == nil {
		fmt.Fprintf(os.Stderr, "\n⚠️  Failed to create CGImage\n")
		return
	}
	// Use generated CoreGraphics binding for CGImageRelease
	defer coregraphics.CGImageRelease(coregraphics.CGImageRef(cgImage))

	// Create file URL for PNG using Foundation bindings
	filename := filepath.Join(h.outputDir, fmt.Sprintf("frame_%04d.png", h.frameCount))

	// Create NSURL from path using generated bindings (takes string directly)
	fileURL := foundation.NewURLFileURLWithPath(filename)
	if fileURL.ID == 0 {
		fmt.Fprintf(os.Stderr, "\n⚠️  Failed to create file URL\n")
		return
	}

	// Get UTTypePNG identifier
	utTypePNGClass := objc.GetClass("UTType")
	png := objc.RegisterName("PNG")
	utTypePNG := objc.ID(utTypePNGClass).Send(png)
	identifier := utTypePNG.Send(objc.RegisterName("identifier"))

	// Create CGImageDestination using generated ImageIO binding
	destination := imageio.CGImageDestinationCreateWithURL(
		unsafe.Pointer(fileURL.ID),
		unsafe.Pointer(identifier),
		1,   // count
		nil, // options
	)
	if destination == nil {
		fmt.Fprintf(os.Stderr, "\n⚠️  Failed to create image destination\n")
		return
	}
	defer func() {
		// CFRelease(destination)
		objc.ID(destination).Send(objc.RegisterName("release"))
	}()

	// Add image to destination using generated ImageIO binding
	imageio.CGImageDestinationAddImage(destination, imageio.CGImageRef(cgImage), nil)

	// Finalize (write to disk) using generated ImageIO binding
	if imageio.CGImageDestinationFinalize(destination) {
		fmt.Printf("\n   💾 Saved frame to: %s\n", filename)
	} else {
		fmt.Fprintf(os.Stderr, "\n⚠️  Failed to finalize image\n")
	}
}

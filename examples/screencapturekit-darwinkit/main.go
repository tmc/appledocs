package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/progrium/darwinkit/dispatch"
	"github.com/progrium/darwinkit/helper/action"
	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/coremedia"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/screencapturekit"
	"github.com/progrium/darwinkit/objc"
)

var e2e = flag.Bool("e2e", false, "Run end-to-end test mode")

var (
	stream           screencapturekit.Stream
	frameCount       int
	isCapturing      bool
	shareableContent screencapturekit.ShareableContent
)

func main() {
	flag.Parse()
	screencapturekit.InitClasses()
	macos.RunApp(launched)
}

func launched(app appkit.Application, delegate *appkit.ApplicationDelegate) {
	w := appkit.NewWindowWithSize(900, 700)
	objc.Retain(&w)
	w.SetTitle("ScreenCaptureKit Demo")
	w.Center()

	// Title label
	titleLabel := appkit.NewTextFieldWithFrame(foundation.Rect{
		Origin: foundation.Point{X: 20, Y: 660},
		Size:   foundation.Size{Width: 860, Height: 30},
	})
	titleLabel.SetStringValue("ScreenCaptureKit - Screen Recording & Capture")
	titleLabel.SetEditable(false)
	titleLabel.SetBezeled(false)
	titleLabel.SetDrawsBackground(false)
	titleLabel.SetFont(appkit.Font_BoldSystemFontOfSize(18))
	titleLabel.SetTextColor(appkit.Color_LabelColor())

	// Info label
	infoLabel := appkit.NewTextFieldWithFrame(foundation.Rect{
		Origin: foundation.Point{X: 20, Y: 630},
		Size:   foundation.Size{Width: 860, Height: 20},
	})
	infoLabel.SetStringValue("Demonstrates screen capture, window listing, and stream recording")
	infoLabel.SetEditable(false)
	infoLabel.SetBezeled(false)
	infoLabel.SetDrawsBackground(false)

	// Status label
	statusLabel := appkit.NewTextFieldWithFrame(foundation.Rect{
		Origin: foundation.Point{X: 20, Y: 600},
		Size:   foundation.Size{Width: 860, Height: 20},
	})
	statusLabel.SetStringValue("Ready - Click 'Get Shareable Content' to start")
	statusLabel.SetEditable(false)
	statusLabel.SetBezeled(false)
	statusLabel.SetDrawsBackground(false)

	// Get Shareable Content button
	getContentBtn := appkit.NewButtonWithFrame(foundation.Rect{
		Origin: foundation.Point{X: 20, Y: 560},
		Size:   foundation.Size{Width: 180, Height: 32},
	})
	getContentBtn.SetTitle("Get Shareable Content")
	getContentBtn.SetBezelStyle(appkit.BezelStyleRounded)

	// Start Capture button
	startCaptureBtn := appkit.NewButtonWithFrame(foundation.Rect{
		Origin: foundation.Point{X: 210, Y: 560},
		Size:   foundation.Size{Width: 150, Height: 32},
	})
	startCaptureBtn.SetTitle("Start Capture")
	startCaptureBtn.SetBezelStyle(appkit.BezelStyleRounded)
	startCaptureBtn.SetEnabled(false)

	// Stop Capture button
	stopCaptureBtn := appkit.NewButtonWithFrame(foundation.Rect{
		Origin: foundation.Point{X: 370, Y: 560},
		Size:   foundation.Size{Width: 150, Height: 32},
	})
	stopCaptureBtn.SetTitle("Stop Capture")
	stopCaptureBtn.SetBezelStyle(appkit.BezelStyleRounded)
	stopCaptureBtn.SetEnabled(false)

	// Results scroll view
	scrollView := appkit.NewScrollViewWithFrame(foundation.Rect{
		Origin: foundation.Point{X: 20, Y: 20},
		Size:   foundation.Size{Width: 860, Height: 530},
	})
	scrollView.SetHasVerticalScroller(true)
	scrollView.SetAutohidesScrollers(true)
	scrollView.SetBorderType(1)

	resultsTextView := appkit.NewTextViewWithFrame(foundation.Rect{
		Size: foundation.Size{Width: 860, Height: 530},
	})
	resultsTextView.SetEditable(false)
	resultsTextView.SetTextColor(appkit.Color_LabelColor())
	resultsTextView.SetBackgroundColor(appkit.Color_TextBackgroundColor())
	resultsTextView.SetFont(appkit.Font_MonospacedSystemFontOfSizeWeight(11, appkit.FontWeightRegular))
	resultsTextView.SetTextColor(appkit.Color_LabelColor())
	scrollView.SetDocumentView(resultsTextView)

	// Update results display
	updateResults := func(text string) {
		str := foundation.NewStringWithString(text)
		resultsTextView.SetString(str.String())
	}

	// Initial message
	updateResults("ScreenCaptureKit Demo\n\n" +
		"This example demonstrates the ScreenCaptureKit framework for screen recording.\n\n" +
		"Steps:\n" +
		"1. Click 'Get Shareable Content' to retrieve available displays and windows\n" +
		"2. Click 'Start Capture' to begin capturing the main display\n" +
		"3. Click 'Stop Capture' to end the capture\n\n" +
		"ScreenCaptureKit enables high-performance screen recording on macOS.")

	// Get Shareable Content action
	action.Set(getContentBtn, func(sender objc.Object) {
		statusLabel.SetStringValue("Retrieving shareable content...")
		statusLabel.SetTextColor(appkit.Color_SecondaryLabelColor())
		log.Println("Requesting shareable content...")

		screencapturekit.ShareableContent_GetShareableContentWithCompletionHandler(func(content screencapturekit.ShareableContent, error foundation.Error) {
			dispatch.MainQueue().DispatchAsync(func() {
				if !error.IsNil() {
					msg := fmt.Sprintf("Error: %s", error.LocalizedDescription())
					statusLabel.SetStringValue(msg)
					statusLabel.SetTextColor(appkit.Color_SecondaryLabelColor())
					log.Printf("Error getting shareable content: %s", error.LocalizedDescription())
					return
				}

				if content.IsNil() {
					statusLabel.SetStringValue("Error: No shareable content returned")
					statusLabel.SetTextColor(appkit.Color_SecondaryLabelColor())
					log.Println("No shareable content returned")
					return
				}

				// Retain the content to prevent it from being released
				content.Retain()
				shareableContent = content

				displays := content.Displays()
				windows := content.Windows()
				apps := content.Applications()

				output := "Shareable Content Retrieved\n\n"
				output += fmt.Sprintf("Displays: %d\n", len(displays))
				for i, display := range displays {
					output += fmt.Sprintf("  Display %d:\n", i)
					output += fmt.Sprintf("    ID: %d\n", display.DisplayID())
					output += fmt.Sprintf("    Size: %dx%d\n", display.Width(), display.Height())
					frame := display.Frame()
					output += fmt.Sprintf("    Frame: (%.0f, %.0f, %.0f, %.0f)\n",
						frame.Origin.X, frame.Origin.Y, frame.Size.Width, frame.Size.Height)
				}

				output += fmt.Sprintf("\nApplications: %d\n", len(apps))
				for i, app := range apps {
					if i >= 10 {
						output += fmt.Sprintf("  ... and %d more applications\n", len(apps)-10)
						break
					}
					output += fmt.Sprintf("  App %d:\n", i)
					output += fmt.Sprintf("    Name: %s\n", app.ApplicationName())
					output += fmt.Sprintf("    Bundle ID: %s\n", app.BundleIdentifier())
					output += fmt.Sprintf("    PID: %d\n", app.ProcessID())
				}

				output += fmt.Sprintf("\nWindows: %d\n", len(windows))
				visibleWindows := 0
				for i, window := range windows {
					if !window.IsOnScreen() {
						continue
					}
					visibleWindows++
					if visibleWindows > 10 {
						output += fmt.Sprintf("  ... and %d more on-screen windows\n", len(windows)-10)
						break
					}
					output += fmt.Sprintf("  Window %d:\n", i)
					output += fmt.Sprintf("    Title: %s\n", window.Title())
					output += fmt.Sprintf("    ID: %d\n", window.WindowID())
					output += fmt.Sprintf("    Layer: %d\n", window.WindowLayer())
					output += fmt.Sprintf("    On Screen: %v\n", window.IsOnScreen())
					frame := window.Frame()
					output += fmt.Sprintf("    Frame: (%.0f, %.0f, %.0f, %.0f)\n",
						frame.Origin.X, frame.Origin.Y, frame.Size.Width, frame.Size.Height)
					if !window.OwningApplication().IsNil() {
						output += fmt.Sprintf("    Owner: %s\n", window.OwningApplication().ApplicationName())
					}
				}

				updateResults(output)
				statusLabel.SetStringValue(fmt.Sprintf("Found %d displays, %d apps, %d windows", len(displays), len(apps), len(windows)))
				startCaptureBtn.SetEnabled(len(displays) > 0)
				log.Printf("Found %d displays, %d apps, %d windows", len(displays), len(apps), len(windows))
			})
		})
	})

	// Start Capture action
	action.Set(startCaptureBtn, func(sender objc.Object) {
		if isCapturing {
			return
		}

		displays := shareableContent.Displays()
		if len(displays) == 0 {
			statusLabel.SetStringValue("Error: No displays available")
			statusLabel.SetTextColor(appkit.Color_SecondaryLabelColor())
			return
		}

		statusLabel.SetStringValue("Starting capture...")
		log.Println("Starting capture...")

		// Use the main display
		mainDisplay := displays[0]

		// Create content filter for the display
		filter := screencapturekit.NewContentFilterWithDisplayExcludingWindows(mainDisplay, nil)
		objc.Retain(&filter)

		// Create stream configuration
		config := screencapturekit.NewStreamConfiguration()
		objc.Retain(&config)

		// Configure stream settings
		config.SetWidth(1920)
		config.SetHeight(1080)
		config.SetQueueDepth(5)
		config.SetShowsCursor(true)
		config.SetPixelFormat(0x42475241) // 'BGRA' format

		// Set minimum frame interval (30 fps)
		frameInterval := coremedia.Time{
			Value:     1,
			Timescale: 30,
			Flags:     1,
		}
		config.SetMinimumFrameInterval(frameInterval)

		// Create stream delegate
		streamDelegate := &screencapturekit.StreamDelegate{}
		streamDelegate.SetStreamDidStopWithError(func(stream screencapturekit.Stream, error foundation.Error) {
			dispatch.MainQueue().DispatchAsync(func() {
				if !error.IsNil() {
					msg := fmt.Sprintf("Stream stopped with error: %s", error.LocalizedDescription())
					statusLabel.SetStringValue(msg)
					log.Printf("Stream error: %s", error.LocalizedDescription())
				} else {
					statusLabel.SetStringValue("Stream stopped")
					log.Println("Stream stopped normally")
				}
				isCapturing = false
				startCaptureBtn.SetEnabled(true)
				stopCaptureBtn.SetEnabled(false)
			})
		})

		// Create stream
		stream = screencapturekit.NewStreamWithFilterConfigurationDelegate(filter, config, streamDelegate)
		objc.Retain(&stream)

		// Create stream output handler
		streamOutput := &screencapturekit.StreamOutput{}
		streamOutput.SetStreamDidOutputSampleBufferOfType(func(s screencapturekit.Stream, sampleBuffer coremedia.SampleBufferRef, outputType screencapturekit.StreamOutputType) {
			frameCount++
			if frameCount%30 == 0 {
				dispatch.MainQueue().DispatchAsync(func() {
					statusLabel.SetStringValue(fmt.Sprintf("Capturing... (Frame %d)", frameCount))
					statusLabel.SetTextColor(appkit.Color_SecondaryLabelColor())
				})
				log.Printf("Received frame %d (type: %d)", frameCount, outputType)
			}
		})

		// Add stream output
		success := stream.AddStreamOutput(streamOutput, screencapturekit.StreamOutputTypeScreen)
		if !success {
			statusLabel.SetStringValue("Error: Failed to add stream output")
			statusLabel.SetTextColor(appkit.Color_SecondaryLabelColor())
			log.Println("Failed to add stream output")
			return
		}

		// Start capture
		stream.StartCaptureWithCompletionHandler(func(error foundation.Error) {
			dispatch.MainQueue().DispatchAsync(func() {
				if !error.IsNil() {
					msg := fmt.Sprintf("Error starting capture: %s", error.LocalizedDescription())
					statusLabel.SetStringValue(msg)
					statusLabel.SetTextColor(appkit.Color_SecondaryLabelColor())
					log.Printf("Error starting capture: %s", error.LocalizedDescription())
					return
				}

				isCapturing = true
				frameCount = 0
				startCaptureBtn.SetEnabled(false)
				stopCaptureBtn.SetEnabled(true)
				statusLabel.SetStringValue("Capturing... (Frame 0)")
				statusLabel.SetTextColor(appkit.Color_SecondaryLabelColor())
				log.Println("Capture started successfully")

				output := "Screen Capture Active\n\n"
				output += fmt.Sprintf("Display: %d\n", mainDisplay.DisplayID())
				output += fmt.Sprintf("Resolution: %dx%d\n", mainDisplay.Width(), mainDisplay.Height())
				output += fmt.Sprintf("Output Size: 1920x1080\n")
				output += fmt.Sprintf("Frame Rate: 30 fps\n")
				output += fmt.Sprintf("Pixel Format: BGRA\n")
				output += fmt.Sprintf("Show Cursor: true\n\n")
				output += "Receiving video frames...\n"
				output += "Frame count will update in the status bar.\n"
				updateResults(output)
			})
		})
	})

	// Stop Capture action
	action.Set(stopCaptureBtn, func(sender objc.Object) {
		if !isCapturing {
			return
		}

		statusLabel.SetStringValue("Stopping capture...")
		log.Println("Stopping capture...")

		stream.StopCaptureWithCompletionHandler(func(error foundation.Error) {
			dispatch.MainQueue().DispatchAsync(func() {
				if !error.IsNil() {
					msg := fmt.Sprintf("Error stopping capture: %s", error.LocalizedDescription())
					statusLabel.SetStringValue(msg)
					statusLabel.SetTextColor(appkit.Color_SecondaryLabelColor())
					log.Printf("Error stopping capture: %s", error.LocalizedDescription())
				} else {
					statusLabel.SetStringValue(fmt.Sprintf("Capture stopped - Total frames: %d", frameCount))
					statusLabel.SetTextColor(appkit.Color_SecondaryLabelColor())
					log.Printf("Capture stopped - Total frames: %d", frameCount)
				}

				isCapturing = false
				startCaptureBtn.SetEnabled(true)
				stopCaptureBtn.SetEnabled(false)

				output := "Screen Capture Stopped\n\n"
				output += fmt.Sprintf("Total frames captured: %d\n", frameCount)
				output += fmt.Sprintf("Average frame rate: %.2f fps\n", float64(frameCount)/10.0)
				output += "\nCapture session completed successfully."
				updateResults(output)
			})
		})
	})

	// Add all subviews
	contentView := w.ContentView()
	contentView.AddSubview(titleLabel)
	contentView.AddSubview(infoLabel)
	contentView.AddSubview(statusLabel)
	contentView.AddSubview(getContentBtn)
	contentView.AddSubview(startCaptureBtn)
	contentView.AddSubview(stopCaptureBtn)
	contentView.AddSubview(scrollView)

	w.MakeKeyAndOrderFront(nil)

	// Enable window close button to actually close the window
	w.SetReleasedWhenClosed(false)

	delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		// Clean up
		if isCapturing && !stream.IsNil() {
			stream.StopCaptureWithCompletionHandler(func(error foundation.Error) {
				log.Println("Cleaned up stream on close")
			})
		}
		return true
	})
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)

	// E2E automation mode
	if *e2e {
		runE2E(app, w, getContentBtn, startCaptureBtn, stopCaptureBtn)
	}
}

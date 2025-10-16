package main

import (
	"log"
	"os"
	"time"

	"github.com/progrium/darwinkit/dispatch"
	"github.com/progrium/darwinkit/macos/appkit"
)

func runE2E(app appkit.Application, w appkit.Window, getContentBtn, startCaptureBtn, stopCaptureBtn appkit.Button) {
	log.Println("Running E2E test mode")

	go func() {
		// Wait for UI to settle
		time.Sleep(1 * time.Second)

		// Test 1: Get shareable content
		dispatch.MainQueue().DispatchAsync(func() {
			log.Println("E2E: Clicking Get Shareable Content button")
			getContentBtn.PerformClick(nil)
		})

		// Wait for content to be retrieved
		time.Sleep(2 * time.Second)

		// Test 2: Start capture
		dispatch.MainQueue().DispatchAsync(func() {
			log.Println("E2E: Clicking Start Capture button")
			startCaptureBtn.PerformClick(nil)
		})

		// Capture for 3 seconds
		time.Sleep(3 * time.Second)

		// Test 3: Stop capture
		dispatch.MainQueue().DispatchAsync(func() {
			log.Println("E2E: Clicking Stop Capture button")
			stopCaptureBtn.PerformClick(nil)
		})

		// Wait for capture to stop
		time.Sleep(1 * time.Second)

		// Verify results
		dispatch.MainQueue().DispatchAsync(func() {
			if frameCount == 0 {
				log.Println("E2E FAIL: No frames captured")
				os.Exit(1)
			}

			if frameCount < 30 {
				log.Printf("E2E WARNING: Only captured %d frames (expected ~90 for 3 seconds at 30fps)", frameCount)
			}

			log.Printf("E2E PASS: Captured %d frames successfully", frameCount)

			// Close the window and quit
			w.Close()
			app.Terminate(nil)
		})

		time.Sleep(500 * time.Millisecond)
		os.Exit(0)
	}()
}

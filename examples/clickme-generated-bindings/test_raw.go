package main

import (
	"fmt"
	"runtime"
	"unsafe"
	
	"github.com/ebitengine/purego/objc"
)

func testRawWindow() {
	runtime.LockOSThread()
	
	// Get NSWindow class directly
	windowClass := objc.GetClass("NSWindow")
	fmt.Printf("NSWindow class: %v\n", windowClass)
	
	// Create window using raw objc
	type NSPoint struct{ X, Y float64 }
	type NSSize struct{ Width, Height float64 }
	type NSRect struct {
		Origin NSPoint
		Size   NSSize
	}
	rect := NSRect{
		Origin: NSPoint{X: 100, Y: 100},
		Size:   NSSize{Width: 400, Height: 300},
	}
	
	styleMask := uintptr(0xb) // Titled | Closable | Resizable
	
	// alloc
	windowID := objc.ID(windowClass).Send(objc.RegisterName("alloc"))
	fmt.Printf("After alloc, windowID: %v\n", windowID)
	
	// initWithContentRect:styleMask:backing:defer:
	windowID = windowID.Send(objc.RegisterName("initWithContentRect:styleMask:backing:defer:"),
		unsafe.Pointer(&rect), styleMask, uintptr(2), false)
	fmt.Printf("After init, windowID: %v\n", windowID)
	
	// Check canBecomeKeyWindow immediately after init
	canBecomeKey := windowID.Send(objc.RegisterName("canBecomeKeyWindow"))
	fmt.Printf("After init - canBecomeKey: %v (raw value: %d)\n", canBecomeKey != 0, canBecomeKey)
	
	// Check styleMask
	actualStyleMask := windowID.Send(objc.RegisterName("styleMask"))
	fmt.Printf("Actual styleMask: 0x%x\n", actualStyleMask)
}

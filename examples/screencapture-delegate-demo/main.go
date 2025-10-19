package main

import (
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/screencapturekit"
)

func init() {
	runtime.LockOSThread()
}

// FrameCounter demonstrates interface-based delegate
type FrameCounter struct {
	count int
}

func (f *FrameCounter) StreamDidOutputSampleBuffer(stream screencapturekit.SCStream, sampleBuffer uintptr, outputType int) {
	f.count++
	fmt.Printf("Frame %d received (type=%d)\n", f.count, outputType)
}

func main() {
	fmt.Println("ScreenCaptureKit Delegate API Demo\n")

	// Example 1: Interface-based delegate (recommended for stateful handlers)
	fmt.Println("1. Interface-based delegate:")
	counter := &FrameCounter{}
	delegate1 := screencapturekit.NewSCStreamOutputDelegate(counter)
	fmt.Printf("   Created: %#x\n\n", uintptr(delegate1))

	// Example 2: Function-based delegate (recommended for simple callbacks)
	fmt.Println("2. Function-based delegate:")
	frameCount := 0
	delegate2 := screencapturekit.NewSimpleSCStreamOutputDelegate(
		func(stream screencapturekit.SCStream, buf uintptr, typ int) {
			frameCount++
			fmt.Printf("   Frame %d\n", frameCount)
		},
	)
	fmt.Printf("   Created: %#x\n\n", uintptr(delegate2))

	fmt.Println("Benefits:")
	fmt.Println("  ✓ Type-safe parameters (SCStream vs objc.ID)")
	fmt.Println("  ✓ No manual selector registration")
	fmt.Println("  ✓ No objc.RegisterClass boilerplate")
	fmt.Println("  ✓ Clean, idiomatic Go code")
}

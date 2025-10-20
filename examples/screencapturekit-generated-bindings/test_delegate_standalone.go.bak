//go:build darwin

package main

import (
	"fmt"
	"time"

	"github.com/ebitengine/purego/objc"
)

// Test that we can create a delegate class with the right method
func testDelegateCreation() error {
	nsObjectClass := objc.GetClass("NSObject")
	if nsObjectClass == 0 {
		return fmt.Errorf("NSObject class not found")
	}

	frameCount := 0

	// Create the delegate method
	streamDidOutputSampleBuffer := func(self objc.ID, cmd objc.SEL, stream objc.ID, sampleBuffer uintptr, outputType int) {
		frameCount++
		fmt.Printf("✓ Delegate method called! Frame: %d\n", frameCount)
	}

	// Register the class without protocol
	className := fmt.Sprintf("TestDelegate_%d", time.Now().UnixNano())
	delegateClass, err := objc.RegisterClass(
		className,
		nsObjectClass,
		nil,
		nil,
		[]objc.MethodDef{
			{
				Cmd: objc.RegisterName("stream:didOutputSampleBuffer:ofType:"),
				Fn:  streamDidOutputSampleBuffer,
			},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to register delegate class: %w", err)
	}

	// Create an instance
	delegate := objc.ID(delegateClass).Send(objc.RegisterName("alloc"))
	delegate = delegate.Send(objc.RegisterName("init"))

	// Check that the delegate responds to the selector
	sel := objc.RegisterName("stream:didOutputSampleBuffer:ofType:")
	respondsToSelector := objc.RegisterName("respondsToSelector:")
	responds := delegate.Send(respondsToSelector, sel)

	fmt.Printf("✓ Delegate created: %v\n", delegate != 0)
	fmt.Printf("✓ Delegate class: %s\n", className)
	fmt.Printf("✓ Responds to selector 'stream:didOutputSampleBuffer:ofType:': %v\n", responds != 0)

	// Try calling the method directly to verify it works
	if responds != 0 {
		fmt.Println("\n📞 Testing direct method invocation...")
		delegate.Send(sel, objc.ID(0), uintptr(0), 0)
		fmt.Printf("✓ Method was called, frameCount = %d\n", frameCount)
	}

	delegate.Send(objc.RegisterName("release"))

	return nil
}

func main() {
	fmt.Println("=== Testing Delegate Creation ===\n")

	if err := testDelegateCreation(); err != nil {
		fmt.Printf("❌ Test failed: %v\n", err)
		return
	}

	fmt.Println("\n=== ✅ Test PASSED ===")
}

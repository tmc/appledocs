package main

import (
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/screencapturekit"
)

func init() {
	runtime.LockOSThread()
}

// Example 1: Using the interface-based delegate
type MyFrameHandler struct {
	frameCount int
}

func (h *MyFrameHandler) StreamDidOutputSampleBuffer(stream screencapturekit.SCStream, sampleBuffer uintptr, outputType int) {
	h.frameCount++
	fmt.Printf("✓ Frame %d received (type=%d, buffer=%#x)\n", h.frameCount, outputType, sampleBuffer)
}

// Example 2: Using the simple function-based delegate
func simpleDelegateExample() {
	frameCount := 0
	delegate := screencapturekit.NewSimpleSCStreamOutputDelegate(func(stream screencapturekit.SCStream, buf uintptr, typ int) {
		frameCount++
		fmt.Printf("✓ Simple delegate - Frame %d (type=%d)\n", frameCount, typ)
	})

	fmt.Printf("Created simple delegate: %v\n", delegate != 0)
}

// Example 3: Using the interface with custom logic
type FrameRecorder struct {
	maxFrames  int
	frameCount int
	frames     []uintptr
}

func (r *FrameRecorder) StreamDidOutputSampleBuffer(stream screencapturekit.SCStream, sampleBuffer uintptr, outputType int) {
	r.frameCount++

	if r.frameCount <= r.maxFrames {
		r.frames = append(r.frames, sampleBuffer)
		fmt.Printf("✓ Recorded frame %d/%d\n", r.frameCount, r.maxFrames)
	}

	if r.frameCount == r.maxFrames {
		fmt.Println("✓ Recording complete!")
	}
}

func main() {
	fmt.Println("=== SCStreamOutput Delegate API Demo ===\n")

	// Example 1: Interface-based delegate
	fmt.Println("--- Example 1: Interface-based Delegate ---")
	handler := &MyFrameHandler{}
	delegate1 := screencapturekit.NewSCStreamOutputDelegate(handler)
	fmt.Printf("Created delegate: %v\n", delegate1 != 0)
	fmt.Println("Handler ready to receive frames via StreamDidOutputSampleBuffer\n")

	// Example 2: Simple function-based delegate
	fmt.Println("--- Example 2: Simple Function-based Delegate ---")
	simpleDelegateExample()
	fmt.Println()

	// Example 3: Frame recorder
	fmt.Println("--- Example 3: Frame Recorder ---")
	recorder := &FrameRecorder{maxFrames: 10}
	delegate3 := screencapturekit.NewSCStreamOutputDelegate(recorder)
	fmt.Printf("Created recorder delegate: %v\n", delegate3 != 0)
	fmt.Println("Recorder ready to collect up to 10 frames\n")

	// Comparison: Old vs New API
	fmt.Println("=== API Comparison ===\n")

	fmt.Println("OLD WAY (manual):")
	fmt.Println(`
  frameCount := 0
  streamDidOutputSampleBuffer := func(self objc.ID, cmd objc.SEL, stream objc.ID, buf uintptr, typ int) {
      frameCount++
      fmt.Printf("Frame %d\n", frameCount)
  }

  class, _ := objc.RegisterClass(
      "MyDelegate",
      objc.GetClass("NSObject"),
      []*objc.Protocol{screencapturekit.SCStreamOutputProtocol},
      nil,
      []objc.MethodDef{{
          Cmd: objc.RegisterName("stream:didOutputSampleBuffer:ofType:"),
          Fn: streamDidOutputSampleBuffer,
      }},
  )
  delegate := objc.ID(class).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
`)

	fmt.Println("\nNEW WAY (type-safe):")
	fmt.Println(`
  frameCount := 0
  delegate := screencapturekit.NewSimpleSCStreamOutputDelegate(
      func(stream screencapturekit.Stream, buf uintptr, typ int) {
          frameCount++
          fmt.Printf("Frame %d\n", frameCount)
      },
  )
`)

	fmt.Println("\nBenefits of New API:")
	fmt.Println("  ✓ Type-safe parameters (Stream instead of objc.ID)")
	fmt.Println("  ✓ No manual selector name construction")
	fmt.Println("  ✓ No manual class registration boilerplate")
	fmt.Println("  ✓ Cleaner, more idiomatic Go code")
	fmt.Println("  ✓ Still flexible - can use objc.RegisterClass for advanced cases")

	fmt.Println("\n=== Demo Complete ===")
}

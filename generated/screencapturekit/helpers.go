// Package-level convenience helpers for ScreenCaptureKit.
//
// This file contains hand-written helpers that complement the generated bindings.
// It is not overwritten during code generation.
package screencapturekit

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// SCStreamOutputHandler defines the interface for handling stream output callbacks.
// Implement this interface to receive frames from an SCStream.
type SCStreamOutputHandler interface {
	// StreamDidOutputSampleBuffer is called when the stream outputs a new sample buffer.
	// The stream parameter is the SCStream that produced the output.
	// The sampleBuffer is a CMSampleBufferRef (as uintptr).
	// The outputType indicates the type of content (screen, audio, etc.).
	StreamDidOutputSampleBuffer(stream SCStream, sampleBuffer uintptr, outputType int)
}

// NewSCStreamOutputDelegate creates an Objective-C delegate that implements the SCStreamOutput protocol.
// The delegate wraps the provided handler and forwards callbacks to it with type-safe parameters.
//
// Example:
//
//	handler := &MyHandler{}
//	delegate := screencapturekit.NewSCStreamOutputDelegate(handler)
//	stream.AddStreamOutput(delegate, ...)
//
// The delegate is automatically registered as a new Objective-C class and initialized.
// You can use it directly with SCStream's addStreamOutput:type:sampleHandlerQueue:error: method.
func NewSCStreamOutputDelegate(handler SCStreamOutputHandler) objc.ID {
	if handler == nil {
		panic("NewSCStreamOutputDelegate: handler cannot be nil")
	}

	// Generate unique class name to avoid conflicts
	className := fmt.Sprintf("SCStreamOutputDelegate_%d", time.Now().UnixNano())

	// Create the delegate method that wraps the handler
	streamDidOutputSampleBuffer := func(self objc.ID, cmd objc.SEL, stream objc.ID, sampleBuffer uintptr, outputType int) {
		// Convert objc.ID to SCStream type for type safety
		streamObj := SCStreamFrom(unsafe.Pointer(stream))
		handler.StreamDidOutputSampleBuffer(streamObj, sampleBuffer, outputType)
	}

	// Register the class with the SCStreamOutput protocol (if available)
	// Note: SCStreamOutput protocol may not be available at runtime via objc.GetProtocol(),
	// but Objective-C uses duck typing - as long as we implement the right methods,
	// the class will work as a delegate even without formal protocol conformance.
	var protocols []*objc.Protocol
	if SCStreamOutputProtocol != nil {
		protocols = []*objc.Protocol{SCStreamOutputProtocol}
	}

	class, err := objc.RegisterClass(
		className,
		objc.GetClass("NSObject"),
		protocols, // May be nil - duck typing will handle it
		nil,
		[]objc.MethodDef{
			{
				Cmd: objc.RegisterName("stream:didOutputSampleBuffer:ofType:"),
				Fn:  streamDidOutputSampleBuffer,
			},
		},
	)
	if err != nil {
		panic(fmt.Sprintf("NewSCStreamOutputDelegate: failed to register class: %v", err))
	}

	// Create and initialize an instance
	delegate := objc.ID(class).Send(objc.RegisterName("alloc"))
	delegate = delegate.Send(objc.RegisterName("init"))

	return delegate
}

// SimpleSCStreamOutputDelegate is a simple function-based implementation of SCStreamOutputHandler.
// It wraps a single callback function for convenience.
type SimpleSCStreamOutputDelegate struct {
	OnSampleBuffer func(stream SCStream, sampleBuffer uintptr, outputType int)
}

// StreamDidOutputSampleBuffer implements the SCStreamOutputHandler interface.
func (d *SimpleSCStreamOutputDelegate) StreamDidOutputSampleBuffer(stream SCStream, sampleBuffer uintptr, outputType int) {
	if d.OnSampleBuffer != nil {
		d.OnSampleBuffer(stream, sampleBuffer, outputType)
	}
}

// NewSimpleSCStreamOutputDelegate creates a delegate from a single callback function.
// This is a convenience wrapper for cases where you only need to handle sample buffers.
//
// Example:
//
//	delegate := screencapturekit.NewSimpleSCStreamOutputDelegate(func(stream screencapturekit.SCStream, buf uintptr, typ int) {
//	    fmt.Printf("Received frame: type=%d\n", typ)
//	})
//	stream.AddStreamOutput(delegate, ...)
func NewSimpleSCStreamOutputDelegate(onSampleBuffer func(SCStream, uintptr, int)) objc.ID {
	handler := &SimpleSCStreamOutputDelegate{
		OnSampleBuffer: onSampleBuffer,
	}
	return NewSCStreamOutputDelegate(handler)
}

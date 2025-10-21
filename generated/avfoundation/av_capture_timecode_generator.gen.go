// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CaptureTimecodeGenerator] class.
var (
	CaptureTimecodeGeneratorClass     _CaptureTimecodeGeneratorClass
	CaptureTimecodeGeneratorClassOnce sync.Once
)

func getCaptureTimecodeGeneratorClass() _CaptureTimecodeGeneratorClass {
	CaptureTimecodeGeneratorClassOnce.Do(func() {
		CaptureTimecodeGeneratorClass = _CaptureTimecodeGeneratorClass{objc.GetClass("AVCaptureTimecodeGenerator")}
	})
	return CaptureTimecodeGeneratorClass
}

type _CaptureTimecodeGeneratorClass struct {
	class objc.Class
}

// An interface definition for the [CaptureTimecodeGenerator] class.
type ICaptureTimecodeGenerator interface {
	objectivec.IObject
	StartSynchronizationWithTimecodeSource(source unsafe.Pointer)
}

// Generates and synchronizes timecode data from various sources for precise video and audio synchronization.
//
// The class supports multiple timecode sources, including frame counting, system clock synchronization, and MIDI timecode input (MTC). Suitable for playback, recording, or other time-sensitive operations where precise timecode metadata is required. Use the method to set up the desired timecode source.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator
type CaptureTimecodeGenerator struct {
	objectivec.Object
}

// CaptureTimecodeGeneratorFrom constructs a [CaptureTimecodeGenerator] from an unsafe.Pointer.
//
// Generates and synchronizes timecode data from various sources for precise video and audio synchronization.
func CaptureTimecodeGeneratorFrom(ptr unsafe.Pointer) CaptureTimecodeGenerator {
	return CaptureTimecodeGenerator{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureTimecodeGeneratorClass) Alloc() CaptureTimecodeGenerator {
	rv := objc.Send[CaptureTimecodeGenerator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureTimecodeGeneratorClass) New() CaptureTimecodeGenerator {
	rv := objc.Send[CaptureTimecodeGenerator](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureTimecodeGenerator) Init() CaptureTimecodeGenerator {
	rv := objc.Send[CaptureTimecodeGenerator](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureTimecodeGenerator) Autorelease() CaptureTimecodeGenerator {
	rv := objc.Send[CaptureTimecodeGenerator](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureTimecodeGenerator creates a new CaptureTimecodeGenerator instance.
func NewCaptureTimecodeGenerator() CaptureTimecodeGenerator {
	return getCaptureTimecodeGeneratorClass().New()
}


// Synchronizes the generator with the specified timecode source.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/startSynchronization(source:)
func (c_ CaptureTimecodeGenerator) StartSynchronizationWithTimecodeSource(source unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("startSynchronizationWithTimecodeSource:"), source)
}




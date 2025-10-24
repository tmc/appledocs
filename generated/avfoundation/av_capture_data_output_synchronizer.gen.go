// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptureDataOutputSynchronizer] class.
var (
	CaptureDataOutputSynchronizerClass     _CaptureDataOutputSynchronizerClass
	CaptureDataOutputSynchronizerClassOnce sync.Once
)

func getCaptureDataOutputSynchronizerClass() _CaptureDataOutputSynchronizerClass {
	CaptureDataOutputSynchronizerClassOnce.Do(func() {
		CaptureDataOutputSynchronizerClass = _CaptureDataOutputSynchronizerClass{objc.GetClass("AVCaptureDataOutputSynchronizer")}
	})
	return CaptureDataOutputSynchronizerClass
}

type _CaptureDataOutputSynchronizerClass struct {
	class objc.Class
}





// An interface definition for the [CaptureDataOutputSynchronizer] class.
type ICaptureDataOutputSynchronizer interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureDataOutputSynchronizerClass) Alloc() CaptureDataOutputSynchronizer {
	rv := objc.Send[CaptureDataOutputSynchronizer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureDataOutputSynchronizerClass) New() CaptureDataOutputSynchronizer {
	rv := objc.Send[CaptureDataOutputSynchronizer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureDataOutputSynchronizer) Init() CaptureDataOutputSynchronizer {
	rv := objc.Send[CaptureDataOutputSynchronizer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureDataOutputSynchronizer) Autorelease() CaptureDataOutputSynchronizer {
	rv := objc.Send[CaptureDataOutputSynchronizer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureDataOutputSynchronizer creates a new CaptureDataOutputSynchronizer instance.
func NewCaptureDataOutputSynchronizer() CaptureDataOutputSynchronizer {
	return getCaptureDataOutputSynchronizerClass().New()
}





// An object that coordinates time-matched delivery of data from multiple capture outputs.
//
// Use this class when you need to capture media from multiple capture outputs and want to receive all data samples from the same timestamp in a single delegate callback. For example, when you use an object to coordinate the output of and objects, you can easily match each captured video frame to depth information captured at the same moment.


// An object that coordinates time-matched delivery of data from multiple capture outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDataOutputSynchronizer
type CaptureDataOutputSynchronizer struct {
	objectivec.Object
}

// CaptureDataOutputSynchronizerFrom constructs a [CaptureDataOutputSynchronizer] from an unsafe.Pointer.
//
// An object that coordinates time-matched delivery of data from multiple capture outputs.
func CaptureDataOutputSynchronizerFrom(ptr unsafe.Pointer) CaptureDataOutputSynchronizer {
	return CaptureDataOutputSynchronizer{objectivec.Object{objc.ID(ptr)}}
}






// Creates a capture output synchronizer for the specified capture outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDataOutputSynchronizer/init(dataOutputs:)
func NewCaptureDataOutputSynchronizerWithDataOutputs(dataOutputs []CaptureOutput) CaptureDataOutputSynchronizer {
	instance := getCaptureDataOutputSynchronizerClass().Alloc()
	rv := objc.Send[CaptureDataOutputSynchronizer](instance.ID, objc.Sel("initWithDataOutputs:"), dataOutputs)
	rv.Autorelease()
	return rv
}




























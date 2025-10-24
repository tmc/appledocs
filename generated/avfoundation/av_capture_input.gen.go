// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptureInput] class.
var (
	CaptureInputClass     _CaptureInputClass
	CaptureInputClassOnce sync.Once
)

func getCaptureInputClass() _CaptureInputClass {
	CaptureInputClassOnce.Do(func() {
		CaptureInputClass = _CaptureInputClass{objc.GetClass("AVCaptureInput")}
	})
	return CaptureInputClass
}

type _CaptureInputClass struct {
	class objc.Class
}





// An interface definition for the [CaptureInput] class.
type ICaptureInput interface {
	objectivec.IObject
	

	// properties:
	Ports() []CaptureInputPort


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureInputClass) Alloc() CaptureInput {
	rv := objc.Send[CaptureInput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureInputClass) New() CaptureInput {
	rv := objc.Send[CaptureInput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureInput) Init() CaptureInput {
	rv := objc.Send[CaptureInput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureInput) Autorelease() CaptureInput {
	rv := objc.Send[CaptureInput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureInput creates a new CaptureInput instance.
func NewCaptureInput() CaptureInput {
	return getCaptureInputClass().New()
}





// An abstract superclass for objects that provide input data to a capture session.
//
// You create concrete instances of this class, such as , to add inputs to a capture session. An input provides one or more streams of media data. For example, input devices can provide both audio and video data. The framework represents each media stream that an input provides as an object. A capture makes connections between capture inputs and capture outputs using a object. The connection defines the mapping between a set of port objects and an .


// An abstract superclass for objects that provide input data to a capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput
type CaptureInput struct {
	objectivec.Object
}

// CaptureInputFrom constructs a [CaptureInput] from an unsafe.Pointer.
//
// An abstract superclass for objects that provide input data to a capture session.
func CaptureInputFrom(ptr unsafe.Pointer) CaptureInput {
	return CaptureInput{objectivec.Object{objc.ID(ptr)}}
}

























// The ports available on a capture input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/ports
func (c_ CaptureInput) Ports() []CaptureInputPort {
	rv := objc.Send[[]CaptureInputPort](c_.ID, objc.Sel("ports"))
	return rv
}









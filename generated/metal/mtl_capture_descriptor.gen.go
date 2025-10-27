// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptureDescriptor] class.
var (
	CaptureDescriptorClass     _CaptureDescriptorClass
	CaptureDescriptorClassOnce sync.Once
)

func getCaptureDescriptorClass() _CaptureDescriptorClass {
	CaptureDescriptorClassOnce.Do(func() {
		CaptureDescriptorClass = _CaptureDescriptorClass{objc.GetClass("MTLCaptureDescriptor")}
	})
	return CaptureDescriptorClass
}

type _CaptureDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [CaptureDescriptor] class.
type ICaptureDescriptor interface {
	objectivec.IObject
	

	// properties:
	CaptureObject() objc.ID
	SetCaptureObject(value objc.ID)
	Destination() CaptureDestination
	SetDestination(value CaptureDestination)
	OutputURL() foundation.foundation.INSURL
	SetOutputURL(value foundation.foundation.INSURL)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureDescriptorClass) Alloc() CaptureDescriptor {
	rv := objc.Send[CaptureDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureDescriptorClass) New() CaptureDescriptor {
	rv := objc.Send[CaptureDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureDescriptor) Init() CaptureDescriptor {
	rv := objc.Send[CaptureDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureDescriptor) Autorelease() CaptureDescriptor {
	rv := objc.Send[CaptureDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureDescriptor creates a new CaptureDescriptor instance.
func NewCaptureDescriptor() CaptureDescriptor {
	return getCaptureDescriptorClass().New()
}





// A configuration for a Metal capture session.


// A configuration for a Metal capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor
type CaptureDescriptor struct {
	objectivec.Object
}

// CaptureDescriptorFrom constructs a [CaptureDescriptor] from an unsafe.Pointer.
//
// A configuration for a Metal capture session.
func CaptureDescriptorFrom(ptr unsafe.Pointer) CaptureDescriptor {
	return CaptureDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// The instance whose contents should be captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/captureObject
func (c_ CaptureDescriptor) CaptureObject() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("captureObject"))
	return rv
}


// The instance whose contents should be captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/captureObject
func (c_ CaptureDescriptor) SetCaptureObject(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCaptureObject:"), value)
}


// The destination for any captured command data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/destination
func (c_ CaptureDescriptor) Destination() CaptureDestination {
	rv := objc.Send[CaptureDestination](c_.ID, objc.Sel("destination"))
	return rv
}


// The destination for any captured command data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/destination
func (c_ CaptureDescriptor) SetDestination(value CaptureDestination) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestination:"), value)
}


// A URL for a file to write the capture data into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/outputURL
func (c_ CaptureDescriptor) OutputURL() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](c_.ID, objc.Sel("outputURL"))
	return rv
}


// A URL for a file to write the capture data into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/outputURL
func (c_ CaptureDescriptor) SetOutputURL(value foundation.foundation.INSURL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputURL:"), value)
}









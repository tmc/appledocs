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
}

// A configuration for a Metal capture session.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CaptureDescriptorClass) Alloc() CaptureDescriptor {
	rv := objc.Send[CaptureDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The instance whose contents should be captured.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/captureObject
func (c_ CaptureDescriptor) CaptureObject() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("captureObject"))
	return rv
}


// SetCaptureObject sets the value of the captureObject property.
// The instance whose contents should be captured.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/captureObject
func (c_ CaptureDescriptor) SetCaptureObject(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCaptureObject:"), value)
}



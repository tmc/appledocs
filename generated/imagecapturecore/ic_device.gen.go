// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ICDevice] class.
var (
	ICDeviceClass     _ICDeviceClass
	ICDeviceClassOnce sync.Once
)

func getICDeviceClass() _ICDeviceClass {
	ICDeviceClassOnce.Do(func() {
		ICDeviceClass = _ICDeviceClass{objc.GetClass("ICDevice")}
	})
	return ICDeviceClass
}

type _ICDeviceClass struct {
	class objc.Class
}

// An interface definition for the [ICDevice] class.
type IICDevice interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other ImageCaptureCore classes.


// A parent class referenced by other ImageCaptureCore classes. [Full Topic]
type ICDevice struct {
	objectivec.Object
}

// ICDeviceFrom constructs a [ICDevice] from an unsafe.Pointer.
//
// A parent class referenced by other ImageCaptureCore classes.
func ICDeviceFrom(ptr unsafe.Pointer) ICDevice {
	return ICDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _ICDeviceClass) Alloc() ICDevice {
	rv := objc.Send[ICDevice](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ICDeviceClass) New() ICDevice {
	rv := objc.Send[ICDevice](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICDevice) Init() ICDevice {
	rv := objc.Send[ICDevice](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICDevice) Autorelease() ICDevice {
	rv := objc.Send[ICDevice](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICDevice creates a new ICDevice instance.
func NewICDevice() ICDevice {
	return getICDeviceClass().New()
}





// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZGraphicsDisplay] class.
var (
	VZGraphicsDisplayClass     _VZGraphicsDisplayClass
	VZGraphicsDisplayClassOnce sync.Once
)

func getVZGraphicsDisplayClass() _VZGraphicsDisplayClass {
	VZGraphicsDisplayClassOnce.Do(func() {
		VZGraphicsDisplayClass = _VZGraphicsDisplayClass{objc.GetClass("VZGraphicsDisplay")}
	})
	return VZGraphicsDisplayClass
}

type _VZGraphicsDisplayClass struct {
	class objc.Class
}

// An interface definition for the [VZGraphicsDisplay] class.
type IVZGraphicsDisplay interface {
	objectivec.IObject
	AddObserver(observer objc.ID)
	ReconfigureWithConfigurationError(configuration unsafe.Pointer, error_ unsafe.Pointer) bool
	ReconfigureWithSizeInPixelsError(sizeInPixels coregraphics.CGSize, error_ unsafe.Pointer) bool
	RemoveObserver(observer objc.ID)
}

// A class that represents a graphics display in a VM.
//
// Don’t instantiate a directly. Graphics displays are first configured on a subclass. When you create a from the configuration, the displays are available through the property of the configuration’s .
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay
type VZGraphicsDisplay struct {
	objectivec.Object
}

// VZGraphicsDisplayFrom constructs a [VZGraphicsDisplay] from an unsafe.Pointer.
//
// A class that represents a graphics display in a VM.
func VZGraphicsDisplayFrom(ptr unsafe.Pointer) VZGraphicsDisplay {
	return VZGraphicsDisplay{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZGraphicsDisplayClass) Alloc() VZGraphicsDisplay {
	rv := objc.Send[VZGraphicsDisplay](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZGraphicsDisplayClass) New() VZGraphicsDisplay {
	rv := objc.Send[VZGraphicsDisplay](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZGraphicsDisplay) Init() VZGraphicsDisplay {
	rv := objc.Send[VZGraphicsDisplay](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZGraphicsDisplay) Autorelease() VZGraphicsDisplay {
	rv := objc.Send[VZGraphicsDisplay](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGraphicsDisplay creates a new VZGraphicsDisplay instance.
func NewVZGraphicsDisplay() VZGraphicsDisplay {
	return getVZGraphicsDisplayClass().New()
}


// Adds an observer to notify about display configuration changes.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/addObserver(_:)
func (v_ VZGraphicsDisplay) AddObserver(observer objc.ID) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addObserver:"), observer)
}

// Reconfigure this display with the new display configuration you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/reconfigure(configuration:)
func (v_ VZGraphicsDisplay) ReconfigureWithConfigurationError(configuration unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("reconfigureWithConfiguration:error:"), configuration, error_)
	return rv
}

// Resize this display with the new dimensions you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/reconfigure(sizeInPixels:)
func (v_ VZGraphicsDisplay) ReconfigureWithSizeInPixelsError(sizeInPixels coregraphics.CGSize, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("reconfigureWithSizeInPixels:error:"), sizeInPixels, error_)
	return rv
}

// Removes a display configuration change observer.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/removeObserver(_:)
func (v_ VZGraphicsDisplay) RemoveObserver(observer objc.ID) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeObserver:"), observer)
}

// Returns the size of the display, in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/sizeInPixels
func (v_ VZGraphicsDisplay) SizeInPixels() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("sizeInPixels"))
	return rv
}




// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZMacGraphicsDisplay] class.
var (
	VZMacGraphicsDisplayClass     _VZMacGraphicsDisplayClass
	VZMacGraphicsDisplayClassOnce sync.Once
)

func getVZMacGraphicsDisplayClass() _VZMacGraphicsDisplayClass {
	VZMacGraphicsDisplayClassOnce.Do(func() {
		VZMacGraphicsDisplayClass = _VZMacGraphicsDisplayClass{objc.GetClass("VZMacGraphicsDisplay")}
	})
	return VZMacGraphicsDisplayClass
}

type _VZMacGraphicsDisplayClass struct {
	class objc.Class
}

// An interface definition for the [VZMacGraphicsDisplay] class.
type IVZMacGraphicsDisplay interface {
	objectivec.IObject
}

// An object that represents the graphics display on a Mac.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplay
type VZMacGraphicsDisplay struct {
	objectivec.Object
}

// VZMacGraphicsDisplayFrom constructs a [VZMacGraphicsDisplay] from an unsafe.Pointer.
//
// An object that represents the graphics display on a Mac.
func VZMacGraphicsDisplayFrom(ptr unsafe.Pointer) VZMacGraphicsDisplay {
	return VZMacGraphicsDisplay{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZMacGraphicsDisplayClass) Alloc() VZMacGraphicsDisplay {
	rv := objc.Send[VZMacGraphicsDisplay](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZMacGraphicsDisplayClass) New() VZMacGraphicsDisplay {
	rv := objc.Send[VZMacGraphicsDisplay](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacGraphicsDisplay) Init() VZMacGraphicsDisplay {
	rv := objc.Send[VZMacGraphicsDisplay](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacGraphicsDisplay) Autorelease() VZMacGraphicsDisplay {
	rv := objc.Send[VZMacGraphicsDisplay](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGraphicsDisplay creates a new VZMacGraphicsDisplay instance.
func NewVZMacGraphicsDisplay() VZMacGraphicsDisplay {
	return getVZMacGraphicsDisplayClass().New()
}


// Returns the pixel density of the display in pixels per inch.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplay/pixelsPerInch
func (v_ VZMacGraphicsDisplay) PixelsPerInch() int {
	rv := objc.Send[int](v_.ID, objc.Sel("pixelsPerInch"))
	return rv
}




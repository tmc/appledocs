// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVirtioGraphicsScanout] class.
var (
	VZVirtioGraphicsScanoutClass     _VZVirtioGraphicsScanoutClass
	VZVirtioGraphicsScanoutClassOnce sync.Once
)

func getVZVirtioGraphicsScanoutClass() _VZVirtioGraphicsScanoutClass {
	VZVirtioGraphicsScanoutClassOnce.Do(func() {
		VZVirtioGraphicsScanoutClass = _VZVirtioGraphicsScanoutClass{objc.GetClass("VZVirtioGraphicsScanout")}
	})
	return VZVirtioGraphicsScanoutClass
}

type _VZVirtioGraphicsScanoutClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioGraphicsScanout] class.
type IVZVirtioGraphicsScanout interface {
	IVZGraphicsDisplay
}

// A Virtio graphics scanout that corresponds to a Virtio graphics scanout configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanout
type VZVirtioGraphicsScanout struct {
	VZGraphicsDisplay
}

// VZVirtioGraphicsScanoutFrom constructs a [VZVirtioGraphicsScanout] from an unsafe.Pointer.
//
// A Virtio graphics scanout that corresponds to a Virtio graphics scanout configuration.
func VZVirtioGraphicsScanoutFrom(ptr unsafe.Pointer) VZVirtioGraphicsScanout {
	return VZVirtioGraphicsScanout{
		VZGraphicsDisplay: VZGraphicsDisplayFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioGraphicsScanoutClass) Alloc() VZVirtioGraphicsScanout {
	rv := objc.Send[VZVirtioGraphicsScanout](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioGraphicsScanoutClass) New() VZVirtioGraphicsScanout {
	rv := objc.Send[VZVirtioGraphicsScanout](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioGraphicsScanout) Init() VZVirtioGraphicsScanout {
	rv := objc.Send[VZVirtioGraphicsScanout](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioGraphicsScanout) Autorelease() VZVirtioGraphicsScanout {
	rv := objc.Send[VZVirtioGraphicsScanout](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioGraphicsScanout creates a new VZVirtioGraphicsScanout instance.
func NewVZVirtioGraphicsScanout() VZVirtioGraphicsScanout {
	return getVZVirtioGraphicsScanoutClass().New()
}





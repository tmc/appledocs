// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZDirectoryShare] class.
var (
	VZDirectoryShareClass     _VZDirectoryShareClass
	VZDirectoryShareClassOnce sync.Once
)

func getVZDirectoryShareClass() _VZDirectoryShareClass {
	VZDirectoryShareClassOnce.Do(func() {
		VZDirectoryShareClass = _VZDirectoryShareClass{objc.GetClass("VZDirectoryShare")}
	})
	return VZDirectoryShareClass
}

type _VZDirectoryShareClass struct {
	class objc.Class
}

// An interface definition for the [VZDirectoryShare] class.
type IVZDirectoryShare interface {
	objectivec.IObject
}

// The base class for a directory share.
//
// A directory share defines how the system exposes host directories to a guest VM. Don’t instantiate directly, use one of its subclasses such as or instead.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDirectoryShare
type VZDirectoryShare struct {
	objectivec.Object
}

// VZDirectoryShareFrom constructs a [VZDirectoryShare] from an unsafe.Pointer.
//
// The base class for a directory share.
func VZDirectoryShareFrom(ptr unsafe.Pointer) VZDirectoryShare {
	return VZDirectoryShare{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZDirectoryShareClass) Alloc() VZDirectoryShare {
	rv := objc.Send[VZDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZDirectoryShareClass) New() VZDirectoryShare {
	rv := objc.Send[VZDirectoryShare](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZDirectoryShare) Init() VZDirectoryShare {
	rv := objc.Send[VZDirectoryShare](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZDirectoryShare) Autorelease() VZDirectoryShare {
	rv := objc.Send[VZDirectoryShare](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDirectoryShare creates a new VZDirectoryShare instance.
func NewVZDirectoryShare() VZDirectoryShare {
	return getVZDirectoryShareClass().New()
}





// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZLinuxRosettaUnixSocketCachingOptions] class.
var (
	VZLinuxRosettaUnixSocketCachingOptionsClass     _VZLinuxRosettaUnixSocketCachingOptionsClass
	VZLinuxRosettaUnixSocketCachingOptionsClassOnce sync.Once
)

func getVZLinuxRosettaUnixSocketCachingOptionsClass() _VZLinuxRosettaUnixSocketCachingOptionsClass {
	VZLinuxRosettaUnixSocketCachingOptionsClassOnce.Do(func() {
		VZLinuxRosettaUnixSocketCachingOptionsClass = _VZLinuxRosettaUnixSocketCachingOptionsClass{objc.GetClass("VZLinuxRosettaUnixSocketCachingOptions")}
	})
	return VZLinuxRosettaUnixSocketCachingOptionsClass
}

type _VZLinuxRosettaUnixSocketCachingOptionsClass struct {
	class objc.Class
}

// An interface definition for the [VZLinuxRosettaUnixSocketCachingOptions] class.
type IVZLinuxRosettaUnixSocketCachingOptions interface {
	objectivec.IObject
}

// An object that represents caching options for a UNIX domain socket.
//
// This object configures Rosetta to communicate with the Rosetta daemon using a UNIX domain socket.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions
type VZLinuxRosettaUnixSocketCachingOptions struct {
	objectivec.Object
}

// VZLinuxRosettaUnixSocketCachingOptionsFrom constructs a [VZLinuxRosettaUnixSocketCachingOptions] from an unsafe.Pointer.
//
// An object that represents caching options for a UNIX domain socket.
func VZLinuxRosettaUnixSocketCachingOptionsFrom(ptr unsafe.Pointer) VZLinuxRosettaUnixSocketCachingOptions {
	return VZLinuxRosettaUnixSocketCachingOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZLinuxRosettaUnixSocketCachingOptionsClass) Alloc() VZLinuxRosettaUnixSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaUnixSocketCachingOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZLinuxRosettaUnixSocketCachingOptionsClass) New() VZLinuxRosettaUnixSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaUnixSocketCachingOptions](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZLinuxRosettaUnixSocketCachingOptions) Init() VZLinuxRosettaUnixSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaUnixSocketCachingOptions](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZLinuxRosettaUnixSocketCachingOptions) Autorelease() VZLinuxRosettaUnixSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaUnixSocketCachingOptions](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxRosettaUnixSocketCachingOptions creates a new VZLinuxRosettaUnixSocketCachingOptions instance.
func NewVZLinuxRosettaUnixSocketCachingOptions() VZLinuxRosettaUnixSocketCachingOptions {
	return getVZLinuxRosettaUnixSocketCachingOptionsClass().New()
}





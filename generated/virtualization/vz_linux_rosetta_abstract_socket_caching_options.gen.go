// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZLinuxRosettaAbstractSocketCachingOptions] class.
var (
	VZLinuxRosettaAbstractSocketCachingOptionsClass     _VZLinuxRosettaAbstractSocketCachingOptionsClass
	VZLinuxRosettaAbstractSocketCachingOptionsClassOnce sync.Once
)

func getVZLinuxRosettaAbstractSocketCachingOptionsClass() _VZLinuxRosettaAbstractSocketCachingOptionsClass {
	VZLinuxRosettaAbstractSocketCachingOptionsClassOnce.Do(func() {
		VZLinuxRosettaAbstractSocketCachingOptionsClass = _VZLinuxRosettaAbstractSocketCachingOptionsClass{objc.GetClass("VZLinuxRosettaAbstractSocketCachingOptions")}
	})
	return VZLinuxRosettaAbstractSocketCachingOptionsClass
}

type _VZLinuxRosettaAbstractSocketCachingOptionsClass struct {
	class objc.Class
}

// An interface definition for the [VZLinuxRosettaAbstractSocketCachingOptions] class.
type IVZLinuxRosettaAbstractSocketCachingOptions interface {
	IVZLinuxRosettaCachingOptions
}

// Caching options for an abstract socket.
//
// Use this object to configure Rosetta to communicate with the Rosetta daemon using an abstract socket.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAbstractSocketCachingOptions
type VZLinuxRosettaAbstractSocketCachingOptions struct {
	VZLinuxRosettaCachingOptions
}

// VZLinuxRosettaAbstractSocketCachingOptionsFrom constructs a [VZLinuxRosettaAbstractSocketCachingOptions] from an unsafe.Pointer.
//
// Caching options for an abstract socket.
func VZLinuxRosettaAbstractSocketCachingOptionsFrom(ptr unsafe.Pointer) VZLinuxRosettaAbstractSocketCachingOptions {
	return VZLinuxRosettaAbstractSocketCachingOptions{
		VZLinuxRosettaCachingOptions: VZLinuxRosettaCachingOptionsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZLinuxRosettaAbstractSocketCachingOptionsClass) Alloc() VZLinuxRosettaAbstractSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaAbstractSocketCachingOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZLinuxRosettaAbstractSocketCachingOptionsClass) New() VZLinuxRosettaAbstractSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaAbstractSocketCachingOptions](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZLinuxRosettaAbstractSocketCachingOptions) Init() VZLinuxRosettaAbstractSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaAbstractSocketCachingOptions](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZLinuxRosettaAbstractSocketCachingOptions) Autorelease() VZLinuxRosettaAbstractSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaAbstractSocketCachingOptions](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxRosettaAbstractSocketCachingOptions creates a new VZLinuxRosettaAbstractSocketCachingOptions instance.
func NewVZLinuxRosettaAbstractSocketCachingOptions() VZLinuxRosettaAbstractSocketCachingOptions {
	return getVZLinuxRosettaAbstractSocketCachingOptionsClass().New()
}


// Initialize options to set on a Rosetta directory share.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAbstractSocketCachingOptions/initWithName:error:
func NewVZLinuxRosettaAbstractSocketCachingOptionsWithNameError(name string, error_ unsafe.Pointer) VZLinuxRosettaAbstractSocketCachingOptions {
	instance := getVZLinuxRosettaAbstractSocketCachingOptionsClass().Alloc()
	rv := objc.Send[VZLinuxRosettaAbstractSocketCachingOptions](instance.ID, objc.Sel("initWithName:error:"), objc.String(name), error_)
	rv.Autorelease()
	return rv
}


// The name of the abstract socket that Rosetta uses.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAbstractSocketCachingOptions/name
func (v_ VZLinuxRosettaAbstractSocketCachingOptions) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("name"))
	return rv
}



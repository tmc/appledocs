// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	IVZLinuxRosettaCachingOptions
	Path() string
}

// An object that represents caching options for a UNIX domain socket.
//
// This object configures Rosetta to communicate with the Rosetta daemon using a UNIX domain socket.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions
type VZLinuxRosettaUnixSocketCachingOptions struct {
	VZLinuxRosettaCachingOptions
}

// VZLinuxRosettaUnixSocketCachingOptionsFrom constructs a [VZLinuxRosettaUnixSocketCachingOptions] from an unsafe.Pointer.
//
// An object that represents caching options for a UNIX domain socket.
func VZLinuxRosettaUnixSocketCachingOptionsFrom(ptr unsafe.Pointer) VZLinuxRosettaUnixSocketCachingOptions {
	return VZLinuxRosettaUnixSocketCachingOptions{
		VZLinuxRosettaCachingOptions: VZLinuxRosettaCachingOptionsFrom(ptr),
	}
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




// Creates a new Rosetta caching options object for a UNIX domain socket with the path you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions/initWithPath:error:
func NewVZLinuxRosettaUnixSocketCachingOptionsWithPathError(path string, error_ unsafe.Pointer) VZLinuxRosettaUnixSocketCachingOptions {
	instance := getVZLinuxRosettaUnixSocketCachingOptionsClass().Alloc()
	rv := objc.Send[VZLinuxRosettaUnixSocketCachingOptions](instance.ID, objc.Sel("initWithPath:error:"), objc.String(path), error_)
	rv.Autorelease()
	return rv
}


// The maximum allowed length of the path to the UNIX domain socket.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions/maximumPathLength
func (vc _VZLinuxRosettaUnixSocketCachingOptionsClass) MaximumPathLength() uint {
	rv := objc.Send[uint](objc.ID(vc.class), objc.Sel("maximumPathLength"))
	return rv
}
// The maximum allowed length of the path to the UNIX domain socket.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions/maximumPathLength
func (v_ VZLinuxRosettaUnixSocketCachingOptions) MaximumPathLength() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("maximumPathLength"))
	return rv
}

// The path to the UNIX domain socket that Rosetta uses.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions/path
func (v_ VZLinuxRosettaUnixSocketCachingOptions) Path() string {
	rv := objc.Send[string](v_.ID, objc.Sel("path"))
	return rv
}



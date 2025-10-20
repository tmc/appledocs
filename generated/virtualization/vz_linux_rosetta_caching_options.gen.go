// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZLinuxRosettaCachingOptions] class.
var (
	VZLinuxRosettaCachingOptionsClass     _VZLinuxRosettaCachingOptionsClass
	VZLinuxRosettaCachingOptionsClassOnce sync.Once
)

func getVZLinuxRosettaCachingOptionsClass() _VZLinuxRosettaCachingOptionsClass {
	VZLinuxRosettaCachingOptionsClassOnce.Do(func() {
		VZLinuxRosettaCachingOptionsClass = _VZLinuxRosettaCachingOptionsClass{objc.GetClass("VZLinuxRosettaCachingOptions")}
	})
	return VZLinuxRosettaCachingOptionsClass
}

type _VZLinuxRosettaCachingOptionsClass struct {
	class objc.Class
}

// An interface definition for the [VZLinuxRosettaCachingOptions] class.
type IVZLinuxRosettaCachingOptions interface {
	objectivec.IObject
}

// An abstract class that defines UNIX socket-based caching options for Rosetta.
//
// define the communication mechanism between the Rosetta daemon and the Rosetta runtime. Don’t instantiate directly. Use one of its subclasses, such as or instead.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions
type VZLinuxRosettaCachingOptions struct {
	objectivec.Object
}

// VZLinuxRosettaCachingOptionsFrom constructs a [VZLinuxRosettaCachingOptions] from an unsafe.Pointer.
//
// An abstract class that defines UNIX socket-based caching options for Rosetta.
func VZLinuxRosettaCachingOptionsFrom(ptr unsafe.Pointer) VZLinuxRosettaCachingOptions {
	return VZLinuxRosettaCachingOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZLinuxRosettaCachingOptionsClass) Alloc() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZLinuxRosettaCachingOptionsClass) New() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZLinuxRosettaCachingOptions) Init() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZLinuxRosettaCachingOptions) Autorelease() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxRosettaCachingOptions creates a new VZLinuxRosettaCachingOptions instance.
func NewVZLinuxRosettaCachingOptions() VZLinuxRosettaCachingOptions {
	return getVZLinuxRosettaCachingOptionsClass().New()
}





// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterAccessControl] class.
var (
	MTRClusterAccessControlClass     _MTRClusterAccessControlClass
	MTRClusterAccessControlClassOnce sync.Once
)

func getMTRClusterAccessControlClass() _MTRClusterAccessControlClass {
	MTRClusterAccessControlClassOnce.Do(func() {
		MTRClusterAccessControlClass = _MTRClusterAccessControlClass{objc.GetClass("MTRClusterAccessControl")}
	})
	return MTRClusterAccessControlClass
}

type _MTRClusterAccessControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterAccessControl] class.
type IMTRClusterAccessControl interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterAccessControl
type MTRClusterAccessControl struct {
	MTRGenericCluster
}

// MTRClusterAccessControlFrom constructs a [MTRClusterAccessControl] from an unsafe.Pointer.
func MTRClusterAccessControlFrom(ptr unsafe.Pointer) MTRClusterAccessControl {
	return MTRClusterAccessControl{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterAccessControlClass) Alloc() MTRClusterAccessControl {
	rv := objc.Send[MTRClusterAccessControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterAccessControlClass) New() MTRClusterAccessControl {
	rv := objc.Send[MTRClusterAccessControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterAccessControl) Init() MTRClusterAccessControl {
	rv := objc.Send[MTRClusterAccessControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterAccessControl) Autorelease() MTRClusterAccessControl {
	rv := objc.Send[MTRClusterAccessControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterAccessControl creates a new MTRClusterAccessControl instance.
func NewMTRClusterAccessControl() MTRClusterAccessControl {
	return getMTRClusterAccessControlClass().New()
}





// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterAccessControl] class.
var (
	MTRBaseClusterAccessControlClass     _MTRBaseClusterAccessControlClass
	MTRBaseClusterAccessControlClassOnce sync.Once
)

func getMTRBaseClusterAccessControlClass() _MTRBaseClusterAccessControlClass {
	MTRBaseClusterAccessControlClassOnce.Do(func() {
		MTRBaseClusterAccessControlClass = _MTRBaseClusterAccessControlClass{objc.GetClass("MTRBaseClusterAccessControl")}
	})
	return MTRBaseClusterAccessControlClass
}

type _MTRBaseClusterAccessControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterAccessControl] class.
type IMTRBaseClusterAccessControl interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterAccessControl
type MTRBaseClusterAccessControl struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterAccessControlFrom constructs a [MTRBaseClusterAccessControl] from an unsafe.Pointer.
func MTRBaseClusterAccessControlFrom(ptr unsafe.Pointer) MTRBaseClusterAccessControl {
	return MTRBaseClusterAccessControl{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterAccessControlClass) Alloc() MTRBaseClusterAccessControl {
	rv := objc.Send[MTRBaseClusterAccessControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterAccessControlClass) New() MTRBaseClusterAccessControl {
	rv := objc.Send[MTRBaseClusterAccessControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterAccessControl) Init() MTRBaseClusterAccessControl {
	rv := objc.Send[MTRBaseClusterAccessControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterAccessControl) Autorelease() MTRBaseClusterAccessControl {
	rv := objc.Send[MTRBaseClusterAccessControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterAccessControl creates a new MTRBaseClusterAccessControl instance.
func NewMTRBaseClusterAccessControl() MTRBaseClusterAccessControl {
	return getMTRBaseClusterAccessControlClass().New()
}





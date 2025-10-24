// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterGroupKeyManagement] class.
var (
	MTRClusterGroupKeyManagementClass     _MTRClusterGroupKeyManagementClass
	MTRClusterGroupKeyManagementClassOnce sync.Once
)

func getMTRClusterGroupKeyManagementClass() _MTRClusterGroupKeyManagementClass {
	MTRClusterGroupKeyManagementClassOnce.Do(func() {
		MTRClusterGroupKeyManagementClass = _MTRClusterGroupKeyManagementClass{objc.GetClass("MTRClusterGroupKeyManagement")}
	})
	return MTRClusterGroupKeyManagementClass
}

type _MTRClusterGroupKeyManagementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterGroupKeyManagement] class.
type IMTRClusterGroupKeyManagement interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterGroupKeyManagement
type MTRClusterGroupKeyManagement struct {
	MTRGenericCluster
}

// MTRClusterGroupKeyManagementFrom constructs a [MTRClusterGroupKeyManagement] from an unsafe.Pointer.
func MTRClusterGroupKeyManagementFrom(ptr unsafe.Pointer) MTRClusterGroupKeyManagement {
	return MTRClusterGroupKeyManagement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterGroupKeyManagementClass) Alloc() MTRClusterGroupKeyManagement {
	rv := objc.Send[MTRClusterGroupKeyManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterGroupKeyManagementClass) New() MTRClusterGroupKeyManagement {
	rv := objc.Send[MTRClusterGroupKeyManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterGroupKeyManagement) Init() MTRClusterGroupKeyManagement {
	rv := objc.Send[MTRClusterGroupKeyManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterGroupKeyManagement) Autorelease() MTRClusterGroupKeyManagement {
	rv := objc.Send[MTRClusterGroupKeyManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterGroupKeyManagement creates a new MTRClusterGroupKeyManagement instance.
func NewMTRClusterGroupKeyManagement() MTRClusterGroupKeyManagement {
	return getMTRClusterGroupKeyManagementClass().New()
}





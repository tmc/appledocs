// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterGroupKeyManagement] class.
var (
	MTRBaseClusterGroupKeyManagementClass     _MTRBaseClusterGroupKeyManagementClass
	MTRBaseClusterGroupKeyManagementClassOnce sync.Once
)

func getMTRBaseClusterGroupKeyManagementClass() _MTRBaseClusterGroupKeyManagementClass {
	MTRBaseClusterGroupKeyManagementClassOnce.Do(func() {
		MTRBaseClusterGroupKeyManagementClass = _MTRBaseClusterGroupKeyManagementClass{objc.GetClass("MTRBaseClusterGroupKeyManagement")}
	})
	return MTRBaseClusterGroupKeyManagementClass
}

type _MTRBaseClusterGroupKeyManagementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterGroupKeyManagement] class.
type IMTRBaseClusterGroupKeyManagement interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterGroupKeyManagement
type MTRBaseClusterGroupKeyManagement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterGroupKeyManagementFrom constructs a [MTRBaseClusterGroupKeyManagement] from an unsafe.Pointer.
func MTRBaseClusterGroupKeyManagementFrom(ptr unsafe.Pointer) MTRBaseClusterGroupKeyManagement {
	return MTRBaseClusterGroupKeyManagement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterGroupKeyManagementClass) Alloc() MTRBaseClusterGroupKeyManagement {
	rv := objc.Send[MTRBaseClusterGroupKeyManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterGroupKeyManagementClass) New() MTRBaseClusterGroupKeyManagement {
	rv := objc.Send[MTRBaseClusterGroupKeyManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterGroupKeyManagement) Init() MTRBaseClusterGroupKeyManagement {
	rv := objc.Send[MTRBaseClusterGroupKeyManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterGroupKeyManagement) Autorelease() MTRBaseClusterGroupKeyManagement {
	rv := objc.Send[MTRBaseClusterGroupKeyManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterGroupKeyManagement creates a new MTRBaseClusterGroupKeyManagement instance.
func NewMTRBaseClusterGroupKeyManagement() MTRBaseClusterGroupKeyManagement {
	return getMTRBaseClusterGroupKeyManagementClass().New()
}





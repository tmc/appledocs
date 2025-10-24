// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterModeSelect] class.
var (
	MTRClusterModeSelectClass     _MTRClusterModeSelectClass
	MTRClusterModeSelectClassOnce sync.Once
)

func getMTRClusterModeSelectClass() _MTRClusterModeSelectClass {
	MTRClusterModeSelectClassOnce.Do(func() {
		MTRClusterModeSelectClass = _MTRClusterModeSelectClass{objc.GetClass("MTRClusterModeSelect")}
	})
	return MTRClusterModeSelectClass
}

type _MTRClusterModeSelectClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterModeSelect] class.
type IMTRClusterModeSelect interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterModeSelect
type MTRClusterModeSelect struct {
	MTRGenericCluster
}

// MTRClusterModeSelectFrom constructs a [MTRClusterModeSelect] from an unsafe.Pointer.
func MTRClusterModeSelectFrom(ptr unsafe.Pointer) MTRClusterModeSelect {
	return MTRClusterModeSelect{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterModeSelectClass) Alloc() MTRClusterModeSelect {
	rv := objc.Send[MTRClusterModeSelect](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterModeSelectClass) New() MTRClusterModeSelect {
	rv := objc.Send[MTRClusterModeSelect](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterModeSelect) Init() MTRClusterModeSelect {
	rv := objc.Send[MTRClusterModeSelect](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterModeSelect) Autorelease() MTRClusterModeSelect {
	rv := objc.Send[MTRClusterModeSelect](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterModeSelect creates a new MTRClusterModeSelect instance.
func NewMTRClusterModeSelect() MTRClusterModeSelect {
	return getMTRClusterModeSelectClass().New()
}





// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterModeSelect] class.
var (
	MTRBaseClusterModeSelectClass     _MTRBaseClusterModeSelectClass
	MTRBaseClusterModeSelectClassOnce sync.Once
)

func getMTRBaseClusterModeSelectClass() _MTRBaseClusterModeSelectClass {
	MTRBaseClusterModeSelectClassOnce.Do(func() {
		MTRBaseClusterModeSelectClass = _MTRBaseClusterModeSelectClass{objc.GetClass("MTRBaseClusterModeSelect")}
	})
	return MTRBaseClusterModeSelectClass
}

type _MTRBaseClusterModeSelectClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterModeSelect] class.
type IMTRBaseClusterModeSelect interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterModeSelect
type MTRBaseClusterModeSelect struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterModeSelectFrom constructs a [MTRBaseClusterModeSelect] from an unsafe.Pointer.
func MTRBaseClusterModeSelectFrom(ptr unsafe.Pointer) MTRBaseClusterModeSelect {
	return MTRBaseClusterModeSelect{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterModeSelectClass) Alloc() MTRBaseClusterModeSelect {
	rv := objc.Send[MTRBaseClusterModeSelect](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterModeSelectClass) New() MTRBaseClusterModeSelect {
	rv := objc.Send[MTRBaseClusterModeSelect](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterModeSelect) Init() MTRBaseClusterModeSelect {
	rv := objc.Send[MTRBaseClusterModeSelect](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterModeSelect) Autorelease() MTRBaseClusterModeSelect {
	rv := objc.Send[MTRBaseClusterModeSelect](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterModeSelect creates a new MTRBaseClusterModeSelect instance.
func NewMTRBaseClusterModeSelect() MTRBaseClusterModeSelect {
	return getMTRBaseClusterModeSelectClass().New()
}





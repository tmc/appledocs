// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterColorControl] class.
var (
	MTRBaseClusterColorControlClass     _MTRBaseClusterColorControlClass
	MTRBaseClusterColorControlClassOnce sync.Once
)

func getMTRBaseClusterColorControlClass() _MTRBaseClusterColorControlClass {
	MTRBaseClusterColorControlClassOnce.Do(func() {
		MTRBaseClusterColorControlClass = _MTRBaseClusterColorControlClass{objc.GetClass("MTRBaseClusterColorControl")}
	})
	return MTRBaseClusterColorControlClass
}

type _MTRBaseClusterColorControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterColorControl] class.
type IMTRBaseClusterColorControl interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl
type MTRBaseClusterColorControl struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterColorControlFrom constructs a [MTRBaseClusterColorControl] from an unsafe.Pointer.
func MTRBaseClusterColorControlFrom(ptr unsafe.Pointer) MTRBaseClusterColorControl {
	return MTRBaseClusterColorControl{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterColorControlClass) Alloc() MTRBaseClusterColorControl {
	rv := objc.Send[MTRBaseClusterColorControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterColorControlClass) New() MTRBaseClusterColorControl {
	rv := objc.Send[MTRBaseClusterColorControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterColorControl) Init() MTRBaseClusterColorControl {
	rv := objc.Send[MTRBaseClusterColorControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterColorControl) Autorelease() MTRBaseClusterColorControl {
	rv := objc.Send[MTRBaseClusterColorControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterColorControl creates a new MTRBaseClusterColorControl instance.
func NewMTRBaseClusterColorControl() MTRBaseClusterColorControl {
	return getMTRBaseClusterColorControlClass().New()
}





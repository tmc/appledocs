// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterColorControl] class.
var (
	MTRClusterColorControlClass     _MTRClusterColorControlClass
	MTRClusterColorControlClassOnce sync.Once
)

func getMTRClusterColorControlClass() _MTRClusterColorControlClass {
	MTRClusterColorControlClassOnce.Do(func() {
		MTRClusterColorControlClass = _MTRClusterColorControlClass{objc.GetClass("MTRClusterColorControl")}
	})
	return MTRClusterColorControlClass
}

type _MTRClusterColorControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterColorControl] class.
type IMTRClusterColorControl interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterColorControl
type MTRClusterColorControl struct {
	MTRGenericCluster
}

// MTRClusterColorControlFrom constructs a [MTRClusterColorControl] from an unsafe.Pointer.
func MTRClusterColorControlFrom(ptr unsafe.Pointer) MTRClusterColorControl {
	return MTRClusterColorControl{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterColorControlClass) Alloc() MTRClusterColorControl {
	rv := objc.Send[MTRClusterColorControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterColorControlClass) New() MTRClusterColorControl {
	rv := objc.Send[MTRClusterColorControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterColorControl) Init() MTRClusterColorControl {
	rv := objc.Send[MTRClusterColorControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterColorControl) Autorelease() MTRClusterColorControl {
	rv := objc.Send[MTRClusterColorControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterColorControl creates a new MTRClusterColorControl instance.
func NewMTRClusterColorControl() MTRClusterColorControl {
	return getMTRClusterColorControlClass().New()
}

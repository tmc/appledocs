// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRGenericCluster] class.
var (
	MTRGenericClusterClass     _MTRGenericClusterClass
	MTRGenericClusterClassOnce sync.Once
)

func getMTRGenericClusterClass() _MTRGenericClusterClass {
	MTRGenericClusterClassOnce.Do(func() {
		MTRGenericClusterClass = _MTRGenericClusterClass{objc.GetClass("MTRGenericCluster")}
	})
	return MTRGenericClusterClass
}

type _MTRGenericClusterClass struct {
	class objc.Class
}

// An interface definition for the [MTRGenericCluster] class.
type IMTRGenericCluster interface {
	IMTRCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGenericCluster
type MTRGenericCluster struct {
	MTRCluster
}

// MTRGenericClusterFrom constructs a [MTRGenericCluster] from an unsafe.Pointer.
func MTRGenericClusterFrom(ptr unsafe.Pointer) MTRGenericCluster {
	return MTRGenericCluster{
		MTRCluster: MTRClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGenericClusterClass) Alloc() MTRGenericCluster {
	rv := objc.Send[MTRGenericCluster](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGenericClusterClass) New() MTRGenericCluster {
	rv := objc.Send[MTRGenericCluster](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGenericCluster) Init() MTRGenericCluster {
	rv := objc.Send[MTRGenericCluster](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGenericCluster) Autorelease() MTRGenericCluster {
	rv := objc.Send[MTRGenericCluster](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGenericCluster creates a new MTRGenericCluster instance.
func NewMTRGenericCluster() MTRGenericCluster {
	return getMTRGenericClusterClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgenericcluster/device
func (m_ MTRGenericCluster) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("device"))
	return rv
}


// SetDevice sets the value of the device property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgenericcluster/device
func (m_ MTRGenericCluster) SetDevice(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDevice:"), value)
}




// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterUserLabel] class.
var (
	MTRClusterUserLabelClass     _MTRClusterUserLabelClass
	MTRClusterUserLabelClassOnce sync.Once
)

func getMTRClusterUserLabelClass() _MTRClusterUserLabelClass {
	MTRClusterUserLabelClassOnce.Do(func() {
		MTRClusterUserLabelClass = _MTRClusterUserLabelClass{objc.GetClass("MTRClusterUserLabel")}
	})
	return MTRClusterUserLabelClass
}

type _MTRClusterUserLabelClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterUserLabel] class.
type IMTRClusterUserLabel interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterUserLabel
type MTRClusterUserLabel struct {
	MTRGenericCluster
}

// MTRClusterUserLabelFrom constructs a [MTRClusterUserLabel] from an unsafe.Pointer.
func MTRClusterUserLabelFrom(ptr unsafe.Pointer) MTRClusterUserLabel {
	return MTRClusterUserLabel{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterUserLabelClass) Alloc() MTRClusterUserLabel {
	rv := objc.Send[MTRClusterUserLabel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterUserLabelClass) New() MTRClusterUserLabel {
	rv := objc.Send[MTRClusterUserLabel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterUserLabel) Init() MTRClusterUserLabel {
	rv := objc.Send[MTRClusterUserLabel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterUserLabel) Autorelease() MTRClusterUserLabel {
	rv := objc.Send[MTRClusterUserLabel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterUserLabel creates a new MTRClusterUserLabel instance.
func NewMTRClusterUserLabel() MTRClusterUserLabel {
	return getMTRClusterUserLabelClass().New()
}





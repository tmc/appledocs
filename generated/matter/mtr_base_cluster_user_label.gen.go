// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterUserLabel] class.
var (
	MTRBaseClusterUserLabelClass     _MTRBaseClusterUserLabelClass
	MTRBaseClusterUserLabelClassOnce sync.Once
)

func getMTRBaseClusterUserLabelClass() _MTRBaseClusterUserLabelClass {
	MTRBaseClusterUserLabelClassOnce.Do(func() {
		MTRBaseClusterUserLabelClass = _MTRBaseClusterUserLabelClass{objc.GetClass("MTRBaseClusterUserLabel")}
	})
	return MTRBaseClusterUserLabelClass
}

type _MTRBaseClusterUserLabelClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterUserLabel] class.
type IMTRBaseClusterUserLabel interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterUserLabel
type MTRBaseClusterUserLabel struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterUserLabelFrom constructs a [MTRBaseClusterUserLabel] from an unsafe.Pointer.
func MTRBaseClusterUserLabelFrom(ptr unsafe.Pointer) MTRBaseClusterUserLabel {
	return MTRBaseClusterUserLabel{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterUserLabelClass) Alloc() MTRBaseClusterUserLabel {
	rv := objc.Send[MTRBaseClusterUserLabel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterUserLabelClass) New() MTRBaseClusterUserLabel {
	rv := objc.Send[MTRBaseClusterUserLabel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterUserLabel) Init() MTRBaseClusterUserLabel {
	rv := objc.Send[MTRBaseClusterUserLabel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterUserLabel) Autorelease() MTRBaseClusterUserLabel {
	rv := objc.Send[MTRBaseClusterUserLabel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterUserLabel creates a new MTRBaseClusterUserLabel instance.
func NewMTRBaseClusterUserLabel() MTRBaseClusterUserLabel {
	return getMTRBaseClusterUserLabelClass().New()
}





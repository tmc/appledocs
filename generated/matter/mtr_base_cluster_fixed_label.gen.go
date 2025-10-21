// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterFixedLabel] class.
var (
	MTRBaseClusterFixedLabelClass     _MTRBaseClusterFixedLabelClass
	MTRBaseClusterFixedLabelClassOnce sync.Once
)

func getMTRBaseClusterFixedLabelClass() _MTRBaseClusterFixedLabelClass {
	MTRBaseClusterFixedLabelClassOnce.Do(func() {
		MTRBaseClusterFixedLabelClass = _MTRBaseClusterFixedLabelClass{objc.GetClass("MTRBaseClusterFixedLabel")}
	})
	return MTRBaseClusterFixedLabelClass
}

type _MTRBaseClusterFixedLabelClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterFixedLabel] class.
type IMTRBaseClusterFixedLabel interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterFixedLabel
type MTRBaseClusterFixedLabel struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterFixedLabelFrom constructs a [MTRBaseClusterFixedLabel] from an unsafe.Pointer.
func MTRBaseClusterFixedLabelFrom(ptr unsafe.Pointer) MTRBaseClusterFixedLabel {
	return MTRBaseClusterFixedLabel{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterFixedLabelClass) Alloc() MTRBaseClusterFixedLabel {
	rv := objc.Send[MTRBaseClusterFixedLabel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterFixedLabelClass) New() MTRBaseClusterFixedLabel {
	rv := objc.Send[MTRBaseClusterFixedLabel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterFixedLabel) Init() MTRBaseClusterFixedLabel {
	rv := objc.Send[MTRBaseClusterFixedLabel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterFixedLabel) Autorelease() MTRBaseClusterFixedLabel {
	rv := objc.Send[MTRBaseClusterFixedLabel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterFixedLabel creates a new MTRBaseClusterFixedLabel instance.
func NewMTRBaseClusterFixedLabel() MTRBaseClusterFixedLabel {
	return getMTRBaseClusterFixedLabelClass().New()
}





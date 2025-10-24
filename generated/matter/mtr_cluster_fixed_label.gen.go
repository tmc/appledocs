// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterFixedLabel] class.
var (
	MTRClusterFixedLabelClass     _MTRClusterFixedLabelClass
	MTRClusterFixedLabelClassOnce sync.Once
)

func getMTRClusterFixedLabelClass() _MTRClusterFixedLabelClass {
	MTRClusterFixedLabelClassOnce.Do(func() {
		MTRClusterFixedLabelClass = _MTRClusterFixedLabelClass{objc.GetClass("MTRClusterFixedLabel")}
	})
	return MTRClusterFixedLabelClass
}

type _MTRClusterFixedLabelClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterFixedLabel] class.
type IMTRClusterFixedLabel interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterFixedLabel
type MTRClusterFixedLabel struct {
	MTRGenericCluster
}

// MTRClusterFixedLabelFrom constructs a [MTRClusterFixedLabel] from an unsafe.Pointer.
func MTRClusterFixedLabelFrom(ptr unsafe.Pointer) MTRClusterFixedLabel {
	return MTRClusterFixedLabel{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterFixedLabelClass) Alloc() MTRClusterFixedLabel {
	rv := objc.Send[MTRClusterFixedLabel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterFixedLabelClass) New() MTRClusterFixedLabel {
	rv := objc.Send[MTRClusterFixedLabel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterFixedLabel) Init() MTRClusterFixedLabel {
	rv := objc.Send[MTRClusterFixedLabel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterFixedLabel) Autorelease() MTRClusterFixedLabel {
	rv := objc.Send[MTRClusterFixedLabel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterFixedLabel creates a new MTRClusterFixedLabel instance.
func NewMTRClusterFixedLabel() MTRClusterFixedLabel {
	return getMTRClusterFixedLabelClass().New()
}

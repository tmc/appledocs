// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRMetrics] class.
var (
	MTRMetricsClass     _MTRMetricsClass
	MTRMetricsClassOnce sync.Once
)

func getMTRMetricsClass() _MTRMetricsClass {
	MTRMetricsClassOnce.Do(func() {
		MTRMetricsClass = _MTRMetricsClass{objc.GetClass("MTRMetrics")}
	})
	return MTRMetricsClass
}

type _MTRMetricsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMetrics] class.
type IMTRMetrics interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMetrics
type MTRMetrics struct {
	objectivec.Object
}

// MTRMetricsFrom constructs a [MTRMetrics] from an unsafe.Pointer.
func MTRMetricsFrom(ptr unsafe.Pointer) MTRMetrics {
	return MTRMetrics{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMetricsClass) Alloc() MTRMetrics {
	rv := objc.Send[MTRMetrics](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMetricsClass) New() MTRMetrics {
	rv := objc.Send[MTRMetrics](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMetrics) Init() MTRMetrics {
	rv := objc.Send[MTRMetrics](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMetrics) Autorelease() MTRMetrics {
	rv := objc.Send[MTRMetrics](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMetrics creates a new MTRMetrics instance.
func NewMTRMetrics() MTRMetrics {
	return getMTRMetricsClass().New()
}





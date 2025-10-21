// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRCommissioningParameters] class.
var (
	MTRCommissioningParametersClass     _MTRCommissioningParametersClass
	MTRCommissioningParametersClassOnce sync.Once
)

func getMTRCommissioningParametersClass() _MTRCommissioningParametersClass {
	MTRCommissioningParametersClassOnce.Do(func() {
		MTRCommissioningParametersClass = _MTRCommissioningParametersClass{objc.GetClass("MTRCommissioningParameters")}
	})
	return MTRCommissioningParametersClass
}

type _MTRCommissioningParametersClass struct {
	class objc.Class
}

// An interface definition for the [MTRCommissioningParameters] class.
type IMTRCommissioningParameters interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters
type MTRCommissioningParameters struct {
	objectivec.Object
}

// MTRCommissioningParametersFrom constructs a [MTRCommissioningParameters] from an unsafe.Pointer.
func MTRCommissioningParametersFrom(ptr unsafe.Pointer) MTRCommissioningParameters {
	return MTRCommissioningParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRCommissioningParametersClass) Alloc() MTRCommissioningParameters {
	rv := objc.Send[MTRCommissioningParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRCommissioningParametersClass) New() MTRCommissioningParameters {
	rv := objc.Send[MTRCommissioningParameters](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCommissioningParameters) Init() MTRCommissioningParameters {
	rv := objc.Send[MTRCommissioningParameters](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCommissioningParameters) Autorelease() MTRCommissioningParameters {
	rv := objc.Send[MTRCommissioningParameters](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCommissioningParameters creates a new MTRCommissioningParameters instance.
func NewMTRCommissioningParameters() MTRCommissioningParameters {
	return getMTRCommissioningParametersClass().New()
}





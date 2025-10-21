// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRCommissionableBrowserResult] class.
var (
	MTRCommissionableBrowserResultClass     _MTRCommissionableBrowserResultClass
	MTRCommissionableBrowserResultClassOnce sync.Once
)

func getMTRCommissionableBrowserResultClass() _MTRCommissionableBrowserResultClass {
	MTRCommissionableBrowserResultClassOnce.Do(func() {
		MTRCommissionableBrowserResultClass = _MTRCommissionableBrowserResultClass{objc.GetClass("MTRCommissionableBrowserResult")}
	})
	return MTRCommissionableBrowserResultClass
}

type _MTRCommissionableBrowserResultClass struct {
	class objc.Class
}

// An interface definition for the [MTRCommissionableBrowserResult] class.
type IMTRCommissionableBrowserResult interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionableBrowserResult
type MTRCommissionableBrowserResult struct {
	objectivec.Object
}

// MTRCommissionableBrowserResultFrom constructs a [MTRCommissionableBrowserResult] from an unsafe.Pointer.
func MTRCommissionableBrowserResultFrom(ptr unsafe.Pointer) MTRCommissionableBrowserResult {
	return MTRCommissionableBrowserResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRCommissionableBrowserResultClass) Alloc() MTRCommissionableBrowserResult {
	rv := objc.Send[MTRCommissionableBrowserResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRCommissionableBrowserResultClass) New() MTRCommissionableBrowserResult {
	rv := objc.Send[MTRCommissionableBrowserResult](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCommissionableBrowserResult) Init() MTRCommissionableBrowserResult {
	rv := objc.Send[MTRCommissionableBrowserResult](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCommissionableBrowserResult) Autorelease() MTRCommissionableBrowserResult {
	rv := objc.Send[MTRCommissionableBrowserResult](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCommissionableBrowserResult creates a new MTRCommissionableBrowserResult instance.
func NewMTRCommissionableBrowserResult() MTRCommissionableBrowserResult {
	return getMTRCommissionableBrowserResultClass().New()
}





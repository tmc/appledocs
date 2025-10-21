// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestStructArrayArgumentResponseParams] class.
var (
	MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass     _MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass
	MTRTestClusterClusterTestStructArrayArgumentResponseParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestStructArrayArgumentResponseParamsClass() _MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass {
	MTRTestClusterClusterTestStructArrayArgumentResponseParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass = _MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass{objc.GetClass("MTRTestClusterClusterTestStructArrayArgumentResponseParams")}
	})
	return MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass
}

type _MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestStructArrayArgumentResponseParams] class.
type IMTRTestClusterClusterTestStructArrayArgumentResponseParams interface {
	IMTRUnitTestingClusterTestStructArrayArgumentResponseParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestStructArrayArgumentResponseParams
type MTRTestClusterClusterTestStructArrayArgumentResponseParams struct {
	MTRUnitTestingClusterTestStructArrayArgumentResponseParams
}

// MTRTestClusterClusterTestStructArrayArgumentResponseParamsFrom constructs a [MTRTestClusterClusterTestStructArrayArgumentResponseParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestStructArrayArgumentResponseParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestStructArrayArgumentResponseParams {
	return MTRTestClusterClusterTestStructArrayArgumentResponseParams{
		MTRUnitTestingClusterTestStructArrayArgumentResponseParams: MTRUnitTestingClusterTestStructArrayArgumentResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass) Alloc() MTRTestClusterClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArrayArgumentResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass) New() MTRTestClusterClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArrayArgumentResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) Init() MTRTestClusterClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArrayArgumentResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) Autorelease() MTRTestClusterClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArrayArgumentResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestStructArrayArgumentResponseParams creates a new MTRTestClusterClusterTestStructArrayArgumentResponseParams instance.
func NewMTRTestClusterClusterTestStructArrayArgumentResponseParams() MTRTestClusterClusterTestStructArrayArgumentResponseParams {
	return getMTRTestClusterClusterTestStructArrayArgumentResponseParamsClass().New()
}





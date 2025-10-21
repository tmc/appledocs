// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestSimpleArgumentResponseParams] class.
var (
	MTRTestClusterClusterTestSimpleArgumentResponseParamsClass     _MTRTestClusterClusterTestSimpleArgumentResponseParamsClass
	MTRTestClusterClusterTestSimpleArgumentResponseParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestSimpleArgumentResponseParamsClass() _MTRTestClusterClusterTestSimpleArgumentResponseParamsClass {
	MTRTestClusterClusterTestSimpleArgumentResponseParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestSimpleArgumentResponseParamsClass = _MTRTestClusterClusterTestSimpleArgumentResponseParamsClass{objc.GetClass("MTRTestClusterClusterTestSimpleArgumentResponseParams")}
	})
	return MTRTestClusterClusterTestSimpleArgumentResponseParamsClass
}

type _MTRTestClusterClusterTestSimpleArgumentResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestSimpleArgumentResponseParams] class.
type IMTRTestClusterClusterTestSimpleArgumentResponseParams interface {
	IMTRUnitTestingClusterTestSimpleArgumentResponseParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestSimpleArgumentResponseParams
type MTRTestClusterClusterTestSimpleArgumentResponseParams struct {
	MTRUnitTestingClusterTestSimpleArgumentResponseParams
}

// MTRTestClusterClusterTestSimpleArgumentResponseParamsFrom constructs a [MTRTestClusterClusterTestSimpleArgumentResponseParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestSimpleArgumentResponseParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestSimpleArgumentResponseParams {
	return MTRTestClusterClusterTestSimpleArgumentResponseParams{
		MTRUnitTestingClusterTestSimpleArgumentResponseParams: MTRUnitTestingClusterTestSimpleArgumentResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestSimpleArgumentResponseParamsClass) Alloc() MTRTestClusterClusterTestSimpleArgumentResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestSimpleArgumentResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestSimpleArgumentResponseParamsClass) New() MTRTestClusterClusterTestSimpleArgumentResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestSimpleArgumentResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestSimpleArgumentResponseParams) Init() MTRTestClusterClusterTestSimpleArgumentResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestSimpleArgumentResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestSimpleArgumentResponseParams) Autorelease() MTRTestClusterClusterTestSimpleArgumentResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestSimpleArgumentResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestSimpleArgumentResponseParams creates a new MTRTestClusterClusterTestSimpleArgumentResponseParams instance.
func NewMTRTestClusterClusterTestSimpleArgumentResponseParams() MTRTestClusterClusterTestSimpleArgumentResponseParams {
	return getMTRTestClusterClusterTestSimpleArgumentResponseParamsClass().New()
}





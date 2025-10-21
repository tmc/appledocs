// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams] class.
var (
	MTRTestClusterClusterTestSimpleOptionalArgumentRequestParamsClass     _MTRTestClusterClusterTestSimpleOptionalArgumentRequestParamsClass
	MTRTestClusterClusterTestSimpleOptionalArgumentRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestSimpleOptionalArgumentRequestParamsClass() _MTRTestClusterClusterTestSimpleOptionalArgumentRequestParamsClass {
	MTRTestClusterClusterTestSimpleOptionalArgumentRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestSimpleOptionalArgumentRequestParamsClass = _MTRTestClusterClusterTestSimpleOptionalArgumentRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams")}
	})
	return MTRTestClusterClusterTestSimpleOptionalArgumentRequestParamsClass
}

type _MTRTestClusterClusterTestSimpleOptionalArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams] class.
type IMTRTestClusterClusterTestSimpleOptionalArgumentRequestParams interface {
	IMTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams
type MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams struct {
	MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams
}

// MTRTestClusterClusterTestSimpleOptionalArgumentRequestParamsFrom constructs a [MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestSimpleOptionalArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams {
	return MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams{
		MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams: MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestSimpleOptionalArgumentRequestParamsClass) Alloc() MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestSimpleOptionalArgumentRequestParamsClass) New() MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams) Init() MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams) Autorelease() MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestSimpleOptionalArgumentRequestParams creates a new MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams instance.
func NewMTRTestClusterClusterTestSimpleOptionalArgumentRequestParams() MTRTestClusterClusterTestSimpleOptionalArgumentRequestParams {
	return getMTRTestClusterClusterTestSimpleOptionalArgumentRequestParamsClass().New()
}





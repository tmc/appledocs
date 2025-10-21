// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestComplexNullableOptionalRequestParams] class.
var (
	MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass     _MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass
	MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass() _MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass {
	MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass = _MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestComplexNullableOptionalRequestParams")}
	})
	return MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass
}

type _MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestComplexNullableOptionalRequestParams] class.
type IMTRTestClusterClusterTestComplexNullableOptionalRequestParams interface {
	IMTRUnitTestingClusterTestComplexNullableOptionalRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestComplexNullableOptionalRequestParams
type MTRTestClusterClusterTestComplexNullableOptionalRequestParams struct {
	MTRUnitTestingClusterTestComplexNullableOptionalRequestParams
}

// MTRTestClusterClusterTestComplexNullableOptionalRequestParamsFrom constructs a [MTRTestClusterClusterTestComplexNullableOptionalRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestComplexNullableOptionalRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestComplexNullableOptionalRequestParams {
	return MTRTestClusterClusterTestComplexNullableOptionalRequestParams{
		MTRUnitTestingClusterTestComplexNullableOptionalRequestParams: MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass) Alloc() MTRTestClusterClusterTestComplexNullableOptionalRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestComplexNullableOptionalRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass) New() MTRTestClusterClusterTestComplexNullableOptionalRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestComplexNullableOptionalRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) Init() MTRTestClusterClusterTestComplexNullableOptionalRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestComplexNullableOptionalRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) Autorelease() MTRTestClusterClusterTestComplexNullableOptionalRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestComplexNullableOptionalRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestComplexNullableOptionalRequestParams creates a new MTRTestClusterClusterTestComplexNullableOptionalRequestParams instance.
func NewMTRTestClusterClusterTestComplexNullableOptionalRequestParams() MTRTestClusterClusterTestComplexNullableOptionalRequestParams {
	return getMTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass().New()
}





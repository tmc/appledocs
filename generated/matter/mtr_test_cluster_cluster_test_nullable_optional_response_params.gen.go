// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestNullableOptionalResponseParams] class.
var (
	MTRTestClusterClusterTestNullableOptionalResponseParamsClass     _MTRTestClusterClusterTestNullableOptionalResponseParamsClass
	MTRTestClusterClusterTestNullableOptionalResponseParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestNullableOptionalResponseParamsClass() _MTRTestClusterClusterTestNullableOptionalResponseParamsClass {
	MTRTestClusterClusterTestNullableOptionalResponseParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestNullableOptionalResponseParamsClass = _MTRTestClusterClusterTestNullableOptionalResponseParamsClass{objc.GetClass("MTRTestClusterClusterTestNullableOptionalResponseParams")}
	})
	return MTRTestClusterClusterTestNullableOptionalResponseParamsClass
}

type _MTRTestClusterClusterTestNullableOptionalResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestNullableOptionalResponseParams] class.
type IMTRTestClusterClusterTestNullableOptionalResponseParams interface {
	IMTRUnitTestingClusterTestNullableOptionalResponseParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestNullableOptionalResponseParams
type MTRTestClusterClusterTestNullableOptionalResponseParams struct {
	MTRUnitTestingClusterTestNullableOptionalResponseParams
}

// MTRTestClusterClusterTestNullableOptionalResponseParamsFrom constructs a [MTRTestClusterClusterTestNullableOptionalResponseParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestNullableOptionalResponseParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestNullableOptionalResponseParams {
	return MTRTestClusterClusterTestNullableOptionalResponseParams{
		MTRUnitTestingClusterTestNullableOptionalResponseParams: MTRUnitTestingClusterTestNullableOptionalResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestNullableOptionalResponseParamsClass) Alloc() MTRTestClusterClusterTestNullableOptionalResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestNullableOptionalResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestNullableOptionalResponseParamsClass) New() MTRTestClusterClusterTestNullableOptionalResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestNullableOptionalResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestNullableOptionalResponseParams) Init() MTRTestClusterClusterTestNullableOptionalResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestNullableOptionalResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestNullableOptionalResponseParams) Autorelease() MTRTestClusterClusterTestNullableOptionalResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestNullableOptionalResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestNullableOptionalResponseParams creates a new MTRTestClusterClusterTestNullableOptionalResponseParams instance.
func NewMTRTestClusterClusterTestNullableOptionalResponseParams() MTRTestClusterClusterTestNullableOptionalResponseParams {
	return getMTRTestClusterClusterTestNullableOptionalResponseParamsClass().New()
}





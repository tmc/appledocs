// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRUnitTestingClusterTestNullableOptionalRequestParams] class.
var (
	MTRUnitTestingClusterTestNullableOptionalRequestParamsClass     _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass
	MTRUnitTestingClusterTestNullableOptionalRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestNullableOptionalRequestParamsClass() _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass {
	MTRUnitTestingClusterTestNullableOptionalRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestNullableOptionalRequestParamsClass = _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestNullableOptionalRequestParams")}
	})
	return MTRUnitTestingClusterTestNullableOptionalRequestParamsClass
}

type _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestNullableOptionalRequestParams] class.
type IMTRUnitTestingClusterTestNullableOptionalRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalRequestParams
type MTRUnitTestingClusterTestNullableOptionalRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestNullableOptionalRequestParamsFrom constructs a [MTRUnitTestingClusterTestNullableOptionalRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestNullableOptionalRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestNullableOptionalRequestParams {
	return MTRUnitTestingClusterTestNullableOptionalRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass) Alloc() MTRUnitTestingClusterTestNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass) New() MTRUnitTestingClusterTestNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) Init() MTRUnitTestingClusterTestNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) Autorelease() MTRUnitTestingClusterTestNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestNullableOptionalRequestParams creates a new MTRUnitTestingClusterTestNullableOptionalRequestParams instance.
func NewMTRUnitTestingClusterTestNullableOptionalRequestParams() MTRUnitTestingClusterTestNullableOptionalRequestParams {
	return getMTRUnitTestingClusterTestNullableOptionalRequestParamsClass().New()
}





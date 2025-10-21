// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestComplexNullableOptionalRequestParams] class.
var (
	MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsClass     _MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsClass
	MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestComplexNullableOptionalRequestParamsClass() _MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsClass {
	MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsClass = _MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestComplexNullableOptionalRequestParams")}
	})
	return MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsClass
}

type _MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestComplexNullableOptionalRequestParams] class.
type IMTRUnitTestingClusterTestComplexNullableOptionalRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams
type MTRUnitTestingClusterTestComplexNullableOptionalRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsFrom constructs a [MTRUnitTestingClusterTestComplexNullableOptionalRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestComplexNullableOptionalRequestParams {
	return MTRUnitTestingClusterTestComplexNullableOptionalRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsClass) Alloc() MTRUnitTestingClusterTestComplexNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestComplexNullableOptionalRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsClass) New() MTRUnitTestingClusterTestComplexNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestComplexNullableOptionalRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) Init() MTRUnitTestingClusterTestComplexNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestComplexNullableOptionalRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) Autorelease() MTRUnitTestingClusterTestComplexNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestComplexNullableOptionalRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestComplexNullableOptionalRequestParams creates a new MTRUnitTestingClusterTestComplexNullableOptionalRequestParams instance.
func NewMTRUnitTestingClusterTestComplexNullableOptionalRequestParams() MTRUnitTestingClusterTestComplexNullableOptionalRequestParams {
	return getMTRUnitTestingClusterTestComplexNullableOptionalRequestParamsClass().New()
}





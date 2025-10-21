// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRUnitTestingClusterTestComplexNullableOptionalResponseParams] class.
var (
	MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsClass     _MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsClass
	MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestComplexNullableOptionalResponseParamsClass() _MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsClass {
	MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsClass = _MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestComplexNullableOptionalResponseParams")}
	})
	return MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsClass
}

type _MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestComplexNullableOptionalResponseParams] class.
type IMTRUnitTestingClusterTestComplexNullableOptionalResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams
type MTRUnitTestingClusterTestComplexNullableOptionalResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsFrom constructs a [MTRUnitTestingClusterTestComplexNullableOptionalResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestComplexNullableOptionalResponseParams {
	return MTRUnitTestingClusterTestComplexNullableOptionalResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsClass) Alloc() MTRUnitTestingClusterTestComplexNullableOptionalResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestComplexNullableOptionalResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsClass) New() MTRUnitTestingClusterTestComplexNullableOptionalResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestComplexNullableOptionalResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) Init() MTRUnitTestingClusterTestComplexNullableOptionalResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestComplexNullableOptionalResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) Autorelease() MTRUnitTestingClusterTestComplexNullableOptionalResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestComplexNullableOptionalResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestComplexNullableOptionalResponseParams creates a new MTRUnitTestingClusterTestComplexNullableOptionalResponseParams instance.
func NewMTRUnitTestingClusterTestComplexNullableOptionalResponseParams() MTRUnitTestingClusterTestComplexNullableOptionalResponseParams {
	return getMTRUnitTestingClusterTestComplexNullableOptionalResponseParamsClass().New()
}





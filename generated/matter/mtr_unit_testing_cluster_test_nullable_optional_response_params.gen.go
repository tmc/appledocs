// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestNullableOptionalResponseParams] class.
var (
	MTRUnitTestingClusterTestNullableOptionalResponseParamsClass     _MTRUnitTestingClusterTestNullableOptionalResponseParamsClass
	MTRUnitTestingClusterTestNullableOptionalResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestNullableOptionalResponseParamsClass() _MTRUnitTestingClusterTestNullableOptionalResponseParamsClass {
	MTRUnitTestingClusterTestNullableOptionalResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestNullableOptionalResponseParamsClass = _MTRUnitTestingClusterTestNullableOptionalResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestNullableOptionalResponseParams")}
	})
	return MTRUnitTestingClusterTestNullableOptionalResponseParamsClass
}

type _MTRUnitTestingClusterTestNullableOptionalResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestNullableOptionalResponseParams] class.
type IMTRUnitTestingClusterTestNullableOptionalResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalResponseParams
type MTRUnitTestingClusterTestNullableOptionalResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestNullableOptionalResponseParamsFrom constructs a [MTRUnitTestingClusterTestNullableOptionalResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestNullableOptionalResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestNullableOptionalResponseParams {
	return MTRUnitTestingClusterTestNullableOptionalResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestNullableOptionalResponseParamsClass) Alloc() MTRUnitTestingClusterTestNullableOptionalResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestNullableOptionalResponseParamsClass) New() MTRUnitTestingClusterTestNullableOptionalResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) Init() MTRUnitTestingClusterTestNullableOptionalResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) Autorelease() MTRUnitTestingClusterTestNullableOptionalResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestNullableOptionalResponseParams creates a new MTRUnitTestingClusterTestNullableOptionalResponseParams instance.
func NewMTRUnitTestingClusterTestNullableOptionalResponseParams() MTRUnitTestingClusterTestNullableOptionalResponseParams {
	return getMTRUnitTestingClusterTestNullableOptionalResponseParamsClass().New()
}





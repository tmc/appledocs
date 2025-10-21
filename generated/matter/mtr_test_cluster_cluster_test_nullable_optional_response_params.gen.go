// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnullableoptionalresponseparams/originalvalue
func (m_ MTRTestClusterClusterTestNullableOptionalResponseParams) OriginalValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("originalValue"))
	return rv
}


// SetOriginalValue sets the value of the originalValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnullableoptionalresponseparams/originalvalue
func (m_ MTRTestClusterClusterTestNullableOptionalResponseParams) SetOriginalValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOriginalValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnullableoptionalresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestNullableOptionalResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnullableoptionalresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestNullableOptionalResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnullableoptionalresponseparams/value
func (m_ MTRTestClusterClusterTestNullableOptionalResponseParams) Value() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnullableoptionalresponseparams/value
func (m_ MTRTestClusterClusterTestNullableOptionalResponseParams) SetValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnullableoptionalresponseparams/wasnull
func (m_ MTRTestClusterClusterTestNullableOptionalResponseParams) WasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("wasNull"))
	return rv
}


// SetWasNull sets the value of the wasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnullableoptionalresponseparams/wasnull
func (m_ MTRTestClusterClusterTestNullableOptionalResponseParams) SetWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnullableoptionalresponseparams/waspresent
func (m_ MTRTestClusterClusterTestNullableOptionalResponseParams) WasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("wasPresent"))
	return rv
}


// SetWasPresent sets the value of the wasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnullableoptionalresponseparams/waspresent
func (m_ MTRTestClusterClusterTestNullableOptionalResponseParams) SetWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWasPresent:"), value)
}




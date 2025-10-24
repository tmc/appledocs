// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	OriginalValue() objc.IObject /* cross-framework: NSNumber */
	SetOriginalValue(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
	WasNull() objc.IObject /* cross-framework: NSNumber */
	SetWasNull(value objc.IObject /* cross-framework: NSNumber */)
	WasPresent() objc.IObject /* cross-framework: NSNumber */
	SetWasPresent(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnullableoptionalresponseparams/originalvalue
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) OriginalValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("originalValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnullableoptionalresponseparams/originalvalue
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) SetOriginalValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOriginalValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnullableoptionalresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnullableoptionalresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnullableoptionalresponseparams/value
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnullableoptionalresponseparams/value
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnullableoptionalresponseparams/wasnull
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) WasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("wasNull"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnullableoptionalresponseparams/wasnull
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) SetWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWasNull:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnullableoptionalresponseparams/waspresent
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) WasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("wasPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnullableoptionalresponseparams/waspresent
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) SetWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWasPresent:"), value)
}




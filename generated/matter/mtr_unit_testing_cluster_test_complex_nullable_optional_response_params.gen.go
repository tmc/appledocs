// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	NullableIntValue() objc.IObject /* cross-framework: NSNumber */
	SetNullableIntValue(value objc.IObject /* cross-framework: NSNumber */)
	NullableIntWasNull() objc.IObject /* cross-framework: NSNumber */
	SetNullableIntWasNull(value objc.IObject /* cross-framework: NSNumber */)
	NullableListValue() unsafe.Pointer
	SetNullableListValue(value unsafe.Pointer)
	NullableListWasNull() objc.IObject /* cross-framework: NSNumber */
	SetNullableListWasNull(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalIntValue() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalIntValue(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalIntWasNull() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalIntWasNull(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalIntWasPresent() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalIntWasPresent(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalListValue() unsafe.Pointer
	SetNullableOptionalListValue(value unsafe.Pointer)
	NullableOptionalListWasNull() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalListWasNull(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalListWasPresent() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalListWasPresent(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalStringValue() objc.IObject /* cross-framework: NSString */
	SetNullableOptionalStringValue(value objc.IObject /* cross-framework: NSString */)
	NullableOptionalStringWasNull() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalStringWasNull(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalStringWasPresent() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalStringWasPresent(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalStructValue() IMTRUnitTestingClusterSimpleStruct
	SetNullableOptionalStructValue(value IMTRUnitTestingClusterSimpleStruct)
	NullableOptionalStructWasNull() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalStructWasNull(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalStructWasPresent() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalStructWasPresent(value objc.IObject /* cross-framework: NSNumber */)
	NullableStringValue() objc.IObject /* cross-framework: NSString */
	SetNullableStringValue(value objc.IObject /* cross-framework: NSString */)
	NullableStringWasNull() objc.IObject /* cross-framework: NSNumber */
	SetNullableStringWasNull(value objc.IObject /* cross-framework: NSNumber */)
	NullableStructValue() IMTRUnitTestingClusterSimpleStruct
	SetNullableStructValue(value IMTRUnitTestingClusterSimpleStruct)
	NullableStructWasNull() objc.IObject /* cross-framework: NSNumber */
	SetNullableStructWasNull(value objc.IObject /* cross-framework: NSNumber */)
	OptionalIntValue() objc.IObject /* cross-framework: NSNumber */
	SetOptionalIntValue(value objc.IObject /* cross-framework: NSNumber */)
	OptionalIntWasPresent() objc.IObject /* cross-framework: NSNumber */
	SetOptionalIntWasPresent(value objc.IObject /* cross-framework: NSNumber */)
	OptionalListValue() unsafe.Pointer
	SetOptionalListValue(value unsafe.Pointer)
	OptionalListWasPresent() objc.IObject /* cross-framework: NSNumber */
	SetOptionalListWasPresent(value objc.IObject /* cross-framework: NSNumber */)
	OptionalStringValue() objc.IObject /* cross-framework: NSString */
	SetOptionalStringValue(value objc.IObject /* cross-framework: NSString */)
	OptionalStringWasPresent() objc.IObject /* cross-framework: NSNumber */
	SetOptionalStringWasPresent(value objc.IObject /* cross-framework: NSNumber */)
	OptionalStructValue() IMTRUnitTestingClusterSimpleStruct
	SetOptionalStructValue(value IMTRUnitTestingClusterSimpleStruct)
	OptionalStructWasPresent() objc.IObject /* cross-framework: NSNumber */
	SetOptionalStructWasPresent(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableintvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableIntValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableIntValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableintvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableIntValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableIntValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableintwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableIntWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableIntWasNull"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableintwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableIntWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableIntWasNull:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablelistvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableListValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableListValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablelistvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableListValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableListValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablelistwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableListWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableListWasNull"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablelistwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableListWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableListWasNull:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalintvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalIntValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalIntValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalintvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalIntValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalIntValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalintwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalIntWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalIntWasNull"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalintwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalIntWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalIntWasNull:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalintwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalIntWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalIntWasPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalintwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalIntWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalIntWasPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionallistvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalListValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableOptionalListValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionallistvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalListValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalListValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionallistwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalListWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalListWasNull"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionallistwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalListWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalListWasNull:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionallistwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalListWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalListWasPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionallistwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalListWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalListWasPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableOptionalStringValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStringValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStringWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalStringWasNull"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStringWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStringWasNull:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStringWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalStringWasPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStringWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStringWasPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStructValue() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableOptionalStructValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStructValue(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStructValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStructWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalStructWasNull"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStructWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStructWasNull:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStructWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalStructWasPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStructWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStructWasPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablestringvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableStringValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablestringvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStringValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablestringwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableStringWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableStringWasNull"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablestringwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableStringWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStringWasNull:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablestructvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableStructValue() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableStructValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablestructvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableStructValue(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStructValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablestructwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableStructWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableStructWasNull"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablestructwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableStructWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStructWasNull:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalintvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalIntValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalIntValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalintvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalIntValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalIntValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalintwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalIntWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalIntWasPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalintwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalIntWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalIntWasPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionallistvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalListValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("optionalListValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionallistvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalListValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalListValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionallistwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalListWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalListWasPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionallistwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalListWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalListWasPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalstringvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("optionalStringValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalstringvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStringValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalstringwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalStringWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalStringWasPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalstringwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalStringWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStringWasPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalstructvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalStructValue() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("optionalStructValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalstructvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalStructValue(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStructValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalstructwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalStructWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalStructWasPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalstructwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalStructWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStructWasPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}




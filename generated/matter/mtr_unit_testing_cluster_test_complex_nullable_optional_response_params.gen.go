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
	NullableIntValue() foundation.Number
	SetNullableIntValue(value foundation.INumber)
	NullableIntWasNull() foundation.Number
	SetNullableIntWasNull(value foundation.INumber)
	NullableListValue() unsafe.Pointer
	SetNullableListValue(value unsafe.Pointer)
	NullableListWasNull() foundation.Number
	SetNullableListWasNull(value foundation.INumber)
	NullableOptionalIntValue() foundation.Number
	SetNullableOptionalIntValue(value foundation.INumber)
	NullableOptionalIntWasNull() foundation.Number
	SetNullableOptionalIntWasNull(value foundation.INumber)
	NullableOptionalIntWasPresent() foundation.Number
	SetNullableOptionalIntWasPresent(value foundation.INumber)
	NullableOptionalListValue() unsafe.Pointer
	SetNullableOptionalListValue(value unsafe.Pointer)
	NullableOptionalListWasNull() foundation.Number
	SetNullableOptionalListWasNull(value foundation.INumber)
	NullableOptionalListWasPresent() foundation.Number
	SetNullableOptionalListWasPresent(value foundation.INumber)
	NullableOptionalStringValue() string
	SetNullableOptionalStringValue(value string)
	NullableOptionalStringWasNull() foundation.Number
	SetNullableOptionalStringWasNull(value foundation.INumber)
	NullableOptionalStringWasPresent() foundation.Number
	SetNullableOptionalStringWasPresent(value foundation.INumber)
	NullableOptionalStructValue() MTRUnitTestingClusterSimpleStruct
	SetNullableOptionalStructValue(value IMTRUnitTestingClusterSimpleStruct)
	NullableOptionalStructWasNull() foundation.Number
	SetNullableOptionalStructWasNull(value foundation.INumber)
	NullableOptionalStructWasPresent() foundation.Number
	SetNullableOptionalStructWasPresent(value foundation.INumber)
	NullableStringValue() string
	SetNullableStringValue(value string)
	NullableStringWasNull() foundation.Number
	SetNullableStringWasNull(value foundation.INumber)
	NullableStructValue() MTRUnitTestingClusterSimpleStruct
	SetNullableStructValue(value IMTRUnitTestingClusterSimpleStruct)
	NullableStructWasNull() foundation.Number
	SetNullableStructWasNull(value foundation.INumber)
	OptionalIntValue() foundation.Number
	SetOptionalIntValue(value foundation.INumber)
	OptionalIntWasPresent() foundation.Number
	SetOptionalIntWasPresent(value foundation.INumber)
	OptionalListValue() unsafe.Pointer
	SetOptionalListValue(value unsafe.Pointer)
	OptionalListWasPresent() foundation.Number
	SetOptionalListWasPresent(value foundation.INumber)
	OptionalStringValue() string
	SetOptionalStringValue(value string)
	OptionalStringWasPresent() foundation.Number
	SetOptionalStringWasPresent(value foundation.INumber)
	OptionalStructValue() MTRUnitTestingClusterSimpleStruct
	SetOptionalStructValue(value IMTRUnitTestingClusterSimpleStruct)
	OptionalStructWasPresent() foundation.Number
	SetOptionalStructWasPresent(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableintvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableIntValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableIntValue"))
	return rv
}


// SetNullableIntValue sets the value of the nullableIntValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableintvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableIntValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableIntValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableintwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableIntWasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableIntWasNull"))
	return rv
}


// SetNullableIntWasNull sets the value of the nullableIntWasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableintwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableIntWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableIntWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablelistvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableListValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableListValue"))
	return rv
}


// SetNullableListValue sets the value of the nullableListValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablelistvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableListValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableListValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablelistwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableListWasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableListWasNull"))
	return rv
}


// SetNullableListWasNull sets the value of the nullableListWasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablelistwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableListWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableListWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalintvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalIntValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalIntValue"))
	return rv
}


// SetNullableOptionalIntValue sets the value of the nullableOptionalIntValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalintvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalIntValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalIntValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalintwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalIntWasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalIntWasNull"))
	return rv
}


// SetNullableOptionalIntWasNull sets the value of the nullableOptionalIntWasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalintwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalIntWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalIntWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalintwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalIntWasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalIntWasPresent"))
	return rv
}


// SetNullableOptionalIntWasPresent sets the value of the nullableOptionalIntWasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalintwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalIntWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalIntWasPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionallistvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalListValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableOptionalListValue"))
	return rv
}


// SetNullableOptionalListValue sets the value of the nullableOptionalListValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionallistvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalListValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalListValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionallistwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalListWasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalListWasNull"))
	return rv
}


// SetNullableOptionalListWasNull sets the value of the nullableOptionalListWasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionallistwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalListWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalListWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionallistwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalListWasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalListWasPresent"))
	return rv
}


// SetNullableOptionalListWasPresent sets the value of the nullableOptionalListWasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionallistwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalListWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalListWasPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStringValue() string {
	rv := objc.Send[string](m_.ID, objc.Sel("nullableOptionalStringValue"))
	return rv
}


// SetNullableOptionalStringValue sets the value of the nullableOptionalStringValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStringValue(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStringValue:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStringWasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalStringWasNull"))
	return rv
}


// SetNullableOptionalStringWasNull sets the value of the nullableOptionalStringWasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStringWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStringWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStringWasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalStringWasPresent"))
	return rv
}


// SetNullableOptionalStringWasPresent sets the value of the nullableOptionalStringWasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStringWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStringWasPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStructValue() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableOptionalStructValue"))
	return rv
}


// SetNullableOptionalStructValue sets the value of the nullableOptionalStructValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStructValue(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStructValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStructWasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalStructWasNull"))
	return rv
}


// SetNullableOptionalStructWasNull sets the value of the nullableOptionalStructWasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStructWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStructWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStructWasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalStructWasPresent"))
	return rv
}


// SetNullableOptionalStructWasPresent sets the value of the nullableOptionalStructWasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStructWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStructWasPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablestringvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableStringValue() string {
	rv := objc.Send[string](m_.ID, objc.Sel("nullableStringValue"))
	return rv
}


// SetNullableStringValue sets the value of the nullableStringValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablestringvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableStringValue(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStringValue:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablestringwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableStringWasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableStringWasNull"))
	return rv
}


// SetNullableStringWasNull sets the value of the nullableStringWasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablestringwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableStringWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStringWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablestructvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableStructValue() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableStructValue"))
	return rv
}


// SetNullableStructValue sets the value of the nullableStructValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablestructvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableStructValue(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStructValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablestructwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableStructWasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableStructWasNull"))
	return rv
}


// SetNullableStructWasNull sets the value of the nullableStructWasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/nullablestructwasnull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableStructWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStructWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalintvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalIntValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionalIntValue"))
	return rv
}


// SetOptionalIntValue sets the value of the optionalIntValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalintvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalIntValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalIntValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalintwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalIntWasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionalIntWasPresent"))
	return rv
}


// SetOptionalIntWasPresent sets the value of the optionalIntWasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalintwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalIntWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalIntWasPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionallistvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalListValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("optionalListValue"))
	return rv
}


// SetOptionalListValue sets the value of the optionalListValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionallistvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalListValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalListValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionallistwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalListWasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionalListWasPresent"))
	return rv
}


// SetOptionalListWasPresent sets the value of the optionalListWasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionallistwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalListWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalListWasPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalstringvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalStringValue() string {
	rv := objc.Send[string](m_.ID, objc.Sel("optionalStringValue"))
	return rv
}


// SetOptionalStringValue sets the value of the optionalStringValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalstringvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalStringValue(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStringValue:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalstringwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalStringWasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionalStringWasPresent"))
	return rv
}


// SetOptionalStringWasPresent sets the value of the optionalStringWasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalstringwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalStringWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStringWasPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalstructvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalStructValue() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("optionalStructValue"))
	return rv
}


// SetOptionalStructValue sets the value of the optionalStructValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalstructvalue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalStructValue(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStructValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalstructwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalStructWasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionalStructWasPresent"))
	return rv
}


// SetOptionalStructWasPresent sets the value of the optionalStructWasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/optionalstructwaspresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalStructWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStructWasPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}




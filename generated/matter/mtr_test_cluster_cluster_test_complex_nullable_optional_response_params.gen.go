// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestComplexNullableOptionalResponseParams] class.
var (
	MTRTestClusterClusterTestComplexNullableOptionalResponseParamsClass     _MTRTestClusterClusterTestComplexNullableOptionalResponseParamsClass
	MTRTestClusterClusterTestComplexNullableOptionalResponseParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestComplexNullableOptionalResponseParamsClass() _MTRTestClusterClusterTestComplexNullableOptionalResponseParamsClass {
	MTRTestClusterClusterTestComplexNullableOptionalResponseParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestComplexNullableOptionalResponseParamsClass = _MTRTestClusterClusterTestComplexNullableOptionalResponseParamsClass{objc.GetClass("MTRTestClusterClusterTestComplexNullableOptionalResponseParams")}
	})
	return MTRTestClusterClusterTestComplexNullableOptionalResponseParamsClass
}

type _MTRTestClusterClusterTestComplexNullableOptionalResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestComplexNullableOptionalResponseParams] class.
type IMTRTestClusterClusterTestComplexNullableOptionalResponseParams interface {
	IMTRUnitTestingClusterTestComplexNullableOptionalResponseParams
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestComplexNullableOptionalResponseParams
type MTRTestClusterClusterTestComplexNullableOptionalResponseParams struct {
	MTRUnitTestingClusterTestComplexNullableOptionalResponseParams
}

// MTRTestClusterClusterTestComplexNullableOptionalResponseParamsFrom constructs a [MTRTestClusterClusterTestComplexNullableOptionalResponseParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestComplexNullableOptionalResponseParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestComplexNullableOptionalResponseParams {
	return MTRTestClusterClusterTestComplexNullableOptionalResponseParams{
		MTRUnitTestingClusterTestComplexNullableOptionalResponseParams: MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestComplexNullableOptionalResponseParamsClass) Alloc() MTRTestClusterClusterTestComplexNullableOptionalResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestComplexNullableOptionalResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestComplexNullableOptionalResponseParamsClass) New() MTRTestClusterClusterTestComplexNullableOptionalResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestComplexNullableOptionalResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) Init() MTRTestClusterClusterTestComplexNullableOptionalResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestComplexNullableOptionalResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) Autorelease() MTRTestClusterClusterTestComplexNullableOptionalResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestComplexNullableOptionalResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestComplexNullableOptionalResponseParams creates a new MTRTestClusterClusterTestComplexNullableOptionalResponseParams instance.
func NewMTRTestClusterClusterTestComplexNullableOptionalResponseParams() MTRTestClusterClusterTestComplexNullableOptionalResponseParams {
	return getMTRTestClusterClusterTestComplexNullableOptionalResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableintvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableIntValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableIntValue"))
	return rv
}


// SetNullableIntValue sets the value of the nullableIntValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableintvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableIntValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableIntValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableintwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableIntWasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableIntWasNull"))
	return rv
}


// SetNullableIntWasNull sets the value of the nullableIntWasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableintwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableIntWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableIntWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablelistvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableListValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableListValue"))
	return rv
}


// SetNullableListValue sets the value of the nullableListValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablelistvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableListValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableListValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablelistwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableListWasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableListWasNull"))
	return rv
}


// SetNullableListWasNull sets the value of the nullableListWasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablelistwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableListWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableListWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalintvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalIntValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalIntValue"))
	return rv
}


// SetNullableOptionalIntValue sets the value of the nullableOptionalIntValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalintvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalIntValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalIntValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalintwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalIntWasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalIntWasNull"))
	return rv
}


// SetNullableOptionalIntWasNull sets the value of the nullableOptionalIntWasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalintwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalIntWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalIntWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalintwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalIntWasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalIntWasPresent"))
	return rv
}


// SetNullableOptionalIntWasPresent sets the value of the nullableOptionalIntWasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalintwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalIntWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalIntWasPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionallistvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalListValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableOptionalListValue"))
	return rv
}


// SetNullableOptionalListValue sets the value of the nullableOptionalListValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionallistvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalListValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalListValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionallistwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalListWasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalListWasNull"))
	return rv
}


// SetNullableOptionalListWasNull sets the value of the nullableOptionalListWasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionallistwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalListWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalListWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionallistwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalListWasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalListWasPresent"))
	return rv
}


// SetNullableOptionalListWasPresent sets the value of the nullableOptionalListWasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionallistwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalListWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalListWasPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalStringValue() string {
	rv := objc.Send[string](m_.ID, objc.Sel("nullableOptionalStringValue"))
	return rv
}


// SetNullableOptionalStringValue sets the value of the nullableOptionalStringValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStringValue(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStringValue:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalStringWasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalStringWasNull"))
	return rv
}


// SetNullableOptionalStringWasNull sets the value of the nullableOptionalStringWasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStringWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStringWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalStringWasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalStringWasPresent"))
	return rv
}


// SetNullableOptionalStringWasPresent sets the value of the nullableOptionalStringWasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStringWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStringWasPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalStructValue() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableOptionalStructValue"))
	return rv
}


// SetNullableOptionalStructValue sets the value of the nullableOptionalStructValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStructValue(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStructValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalStructWasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalStructWasNull"))
	return rv
}


// SetNullableOptionalStructWasNull sets the value of the nullableOptionalStructWasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStructWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStructWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalStructWasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalStructWasPresent"))
	return rv
}


// SetNullableOptionalStructWasPresent sets the value of the nullableOptionalStructWasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStructWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStructWasPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablestringvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableStringValue() string {
	rv := objc.Send[string](m_.ID, objc.Sel("nullableStringValue"))
	return rv
}


// SetNullableStringValue sets the value of the nullableStringValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablestringvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableStringValue(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStringValue:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablestringwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableStringWasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableStringWasNull"))
	return rv
}


// SetNullableStringWasNull sets the value of the nullableStringWasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablestringwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableStringWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStringWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablestructvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableStructValue() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableStructValue"))
	return rv
}


// SetNullableStructValue sets the value of the nullableStructValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablestructvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableStructValue(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStructValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablestructwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableStructWasNull() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableStructWasNull"))
	return rv
}


// SetNullableStructWasNull sets the value of the nullableStructWasNull property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablestructwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableStructWasNull(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStructWasNull:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalintvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) OptionalIntValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionalIntValue"))
	return rv
}


// SetOptionalIntValue sets the value of the optionalIntValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalintvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetOptionalIntValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalIntValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalintwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) OptionalIntWasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionalIntWasPresent"))
	return rv
}


// SetOptionalIntWasPresent sets the value of the optionalIntWasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalintwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetOptionalIntWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalIntWasPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionallistvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) OptionalListValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("optionalListValue"))
	return rv
}


// SetOptionalListValue sets the value of the optionalListValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionallistvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetOptionalListValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalListValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionallistwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) OptionalListWasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionalListWasPresent"))
	return rv
}


// SetOptionalListWasPresent sets the value of the optionalListWasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionallistwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetOptionalListWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalListWasPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalstringvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) OptionalStringValue() string {
	rv := objc.Send[string](m_.ID, objc.Sel("optionalStringValue"))
	return rv
}


// SetOptionalStringValue sets the value of the optionalStringValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalstringvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetOptionalStringValue(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStringValue:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalstringwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) OptionalStringWasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionalStringWasPresent"))
	return rv
}


// SetOptionalStringWasPresent sets the value of the optionalStringWasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalstringwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetOptionalStringWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStringWasPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalstructvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) OptionalStructValue() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("optionalStructValue"))
	return rv
}


// SetOptionalStructValue sets the value of the optionalStructValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalstructvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetOptionalStructValue(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStructValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalstructwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) OptionalStructWasPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionalStructWasPresent"))
	return rv
}


// SetOptionalStructWasPresent sets the value of the optionalStructWasPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalstructwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetOptionalStructWasPresent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStructWasPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}




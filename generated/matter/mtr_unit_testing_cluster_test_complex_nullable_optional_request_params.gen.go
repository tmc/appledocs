// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableint
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableInt() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableInt"))
	return rv
}


// SetNullableInt sets the value of the nullableInt property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableint
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableInt(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableInt:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullablelist
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableList"))
	return rv
}


// SetNullableList sets the value of the nullableList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullablelist
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableoptionalint
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableOptionalInt() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalInt"))
	return rv
}


// SetNullableOptionalInt sets the value of the nullableOptionalInt property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableoptionalint
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalInt(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalInt:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableoptionallist
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableOptionalList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableOptionalList"))
	return rv
}


// SetNullableOptionalList sets the value of the nullableOptionalList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableoptionallist
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableoptionalstring
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableOptionalString() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("nullableOptionalString"))
	return rv
}


// SetNullableOptionalString sets the value of the nullableOptionalString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableoptionalstring
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalString(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalString:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableoptionalstruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableOptionalStruct() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableOptionalStruct"))
	return rv
}


// SetNullableOptionalStruct sets the value of the nullableOptionalStruct property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableoptionalstruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStruct:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullablestring
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableString() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("nullableString"))
	return rv
}


// SetNullableString sets the value of the nullableString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullablestring
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableString(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableString:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullablestruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableStruct() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableStruct"))
	return rv
}


// SetNullableStruct sets the value of the nullableStruct property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullablestruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStruct:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/optionalint
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) OptionalInt() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionalInt"))
	return rv
}


// SetOptionalInt sets the value of the optionalInt property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/optionalint
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetOptionalInt(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalInt:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/optionallist
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) OptionalList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("optionalList"))
	return rv
}


// SetOptionalList sets the value of the optionalList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/optionallist
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetOptionalList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/optionalstring
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) OptionalString() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("optionalString"))
	return rv
}


// SetOptionalString sets the value of the optionalString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/optionalstring
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetOptionalString(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalString:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/optionalstruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) OptionalStruct() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("optionalStruct"))
	return rv
}


// SetOptionalStruct sets the value of the optionalStruct property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/optionalstruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetOptionalStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStruct:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}




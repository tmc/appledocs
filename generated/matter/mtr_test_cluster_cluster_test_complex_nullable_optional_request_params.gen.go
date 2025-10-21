// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestComplexNullableOptionalRequestParams] class.
var (
	MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass     _MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass
	MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass() _MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass {
	MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass = _MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestComplexNullableOptionalRequestParams")}
	})
	return MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass
}

type _MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestComplexNullableOptionalRequestParams] class.
type IMTRTestClusterClusterTestComplexNullableOptionalRequestParams interface {
	IMTRUnitTestingClusterTestComplexNullableOptionalRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestComplexNullableOptionalRequestParams
type MTRTestClusterClusterTestComplexNullableOptionalRequestParams struct {
	MTRUnitTestingClusterTestComplexNullableOptionalRequestParams
}

// MTRTestClusterClusterTestComplexNullableOptionalRequestParamsFrom constructs a [MTRTestClusterClusterTestComplexNullableOptionalRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestComplexNullableOptionalRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestComplexNullableOptionalRequestParams {
	return MTRTestClusterClusterTestComplexNullableOptionalRequestParams{
		MTRUnitTestingClusterTestComplexNullableOptionalRequestParams: MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass) Alloc() MTRTestClusterClusterTestComplexNullableOptionalRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestComplexNullableOptionalRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass) New() MTRTestClusterClusterTestComplexNullableOptionalRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestComplexNullableOptionalRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) Init() MTRTestClusterClusterTestComplexNullableOptionalRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestComplexNullableOptionalRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) Autorelease() MTRTestClusterClusterTestComplexNullableOptionalRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestComplexNullableOptionalRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestComplexNullableOptionalRequestParams creates a new MTRTestClusterClusterTestComplexNullableOptionalRequestParams instance.
func NewMTRTestClusterClusterTestComplexNullableOptionalRequestParams() MTRTestClusterClusterTestComplexNullableOptionalRequestParams {
	return getMTRTestClusterClusterTestComplexNullableOptionalRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableint
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) NullableInt() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableInt"))
	return rv
}


// SetNullableInt sets the value of the nullableInt property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableint
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetNullableInt(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableInt:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullablelist
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) NullableList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableList"))
	return rv
}


// SetNullableList sets the value of the nullableList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullablelist
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetNullableList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableoptionalint
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) NullableOptionalInt() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalInt"))
	return rv
}


// SetNullableOptionalInt sets the value of the nullableOptionalInt property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableoptionalint
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalInt(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalInt:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableoptionallist
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) NullableOptionalList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableOptionalList"))
	return rv
}


// SetNullableOptionalList sets the value of the nullableOptionalList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableoptionallist
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableoptionalstring
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) NullableOptionalString() string {
	rv := objc.Send[string](m_.ID, objc.Sel("nullableOptionalString"))
	return rv
}


// SetNullableOptionalString sets the value of the nullableOptionalString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableoptionalstring
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalString(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalString:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableoptionalstruct
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) NullableOptionalStruct() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableOptionalStruct"))
	return rv
}


// SetNullableOptionalStruct sets the value of the nullableOptionalStruct property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableoptionalstruct
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalStruct(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStruct:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullablestring
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) NullableString() string {
	rv := objc.Send[string](m_.ID, objc.Sel("nullableString"))
	return rv
}


// SetNullableString sets the value of the nullableString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullablestring
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetNullableString(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableString:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullablestruct
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) NullableStruct() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableStruct"))
	return rv
}


// SetNullableStruct sets the value of the nullableStruct property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullablestruct
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetNullableStruct(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStruct:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/optionalint
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) OptionalInt() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionalInt"))
	return rv
}


// SetOptionalInt sets the value of the optionalInt property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/optionalint
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetOptionalInt(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalInt:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/optionallist
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) OptionalList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("optionalList"))
	return rv
}


// SetOptionalList sets the value of the optionalList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/optionallist
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetOptionalList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/optionalstring
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) OptionalString() string {
	rv := objc.Send[string](m_.ID, objc.Sel("optionalString"))
	return rv
}


// SetOptionalString sets the value of the optionalString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/optionalstring
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetOptionalString(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalString:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/optionalstruct
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) OptionalStruct() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("optionalStruct"))
	return rv
}


// SetOptionalStruct sets the value of the optionalStruct property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/optionalstruct
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetOptionalStruct(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStruct:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}




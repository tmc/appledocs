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
	// properties:
	NullableInt() objc.IObject /* cross-framework: NSNumber */
	SetNullableInt(value objc.IObject /* cross-framework: NSNumber */)
	NullableList() unsafe.Pointer
	SetNullableList(value unsafe.Pointer)
	NullableOptionalInt() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalInt(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalList() unsafe.Pointer
	SetNullableOptionalList(value unsafe.Pointer)
	NullableOptionalString() objc.IObject /* cross-framework: NSString */
	SetNullableOptionalString(value objc.IObject /* cross-framework: NSString */)
	NullableOptionalStruct() IMTRUnitTestingClusterSimpleStruct
	SetNullableOptionalStruct(value IMTRUnitTestingClusterSimpleStruct)
	NullableString() objc.IObject /* cross-framework: NSString */
	SetNullableString(value objc.IObject /* cross-framework: NSString */)
	NullableStruct() IMTRUnitTestingClusterSimpleStruct
	SetNullableStruct(value IMTRUnitTestingClusterSimpleStruct)
	OptionalInt() objc.IObject /* cross-framework: NSNumber */
	SetOptionalInt(value objc.IObject /* cross-framework: NSNumber */)
	OptionalList() unsafe.Pointer
	SetOptionalList(value unsafe.Pointer)
	OptionalString() objc.IObject /* cross-framework: NSString */
	SetOptionalString(value objc.IObject /* cross-framework: NSString */)
	OptionalStruct() IMTRUnitTestingClusterSimpleStruct
	SetOptionalStruct(value IMTRUnitTestingClusterSimpleStruct)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableint
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableInt"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableint
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableInt:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullablelist
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableList"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullablelist
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableList:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableoptionalint
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableOptionalInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalInt"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableoptionalint
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalInt:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableoptionallist
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableOptionalList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableOptionalList"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableoptionallist
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalList:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableoptionalstring
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableOptionalString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableOptionalString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableoptionalstring
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableoptionalstruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableOptionalStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableOptionalStruct"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullableoptionalstruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStruct:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullablestring
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullablestring
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullablestruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableStruct"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/nullablestruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStruct:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/optionalint
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) OptionalInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalInt"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/optionalint
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetOptionalInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalInt:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/optionallist
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) OptionalList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("optionalList"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/optionallist
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetOptionalList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalList:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/optionalstring
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) OptionalString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("optionalString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/optionalstring
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetOptionalString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/optionalstruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) OptionalStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("optionalStruct"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/optionalstruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetOptionalStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStruct:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestcomplexnullableoptionalrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}




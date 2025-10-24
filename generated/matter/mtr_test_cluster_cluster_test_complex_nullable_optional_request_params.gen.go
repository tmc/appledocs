// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableint
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) NullableInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableInt"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableint
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetNullableInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableInt:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullablelist
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) NullableList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableList"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullablelist
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetNullableList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableList:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableoptionalint
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) NullableOptionalInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalInt"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableoptionalint
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalInt:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableoptionallist
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) NullableOptionalList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableOptionalList"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableoptionallist
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalList:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableoptionalstring
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) NullableOptionalString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableOptionalString"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableoptionalstring
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalString:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableoptionalstruct
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) NullableOptionalStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableOptionalStruct"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullableoptionalstruct
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStruct:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullablestring
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) NullableString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableString"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullablestring
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetNullableString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableString:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullablestruct
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) NullableStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableStruct"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/nullablestruct
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetNullableStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStruct:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/optionalint
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) OptionalInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalInt"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/optionalint
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetOptionalInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalInt:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/optionallist
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) OptionalList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("optionalList"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/optionallist
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetOptionalList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalList:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/optionalstring
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) OptionalString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("optionalString"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/optionalstring
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetOptionalString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalString:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/optionalstruct
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) OptionalStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("optionalStruct"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/optionalstruct
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetOptionalStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStruct:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestComplexNullableOptionalRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableintvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableIntValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableIntValue"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableintvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableIntValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableIntValue:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableintwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableIntWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableIntWasNull"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableintwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableIntWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableIntWasNull:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablelistvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableListValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableListValue"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablelistvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableListValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableListValue:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablelistwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableListWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableListWasNull"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablelistwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableListWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableListWasNull:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalintvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalIntValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalIntValue"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalintvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalIntValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalIntValue:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalintwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalIntWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalIntWasNull"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalintwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalIntWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalIntWasNull:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalintwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalIntWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalIntWasPresent"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalintwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalIntWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalIntWasPresent:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionallistvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalListValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableOptionalListValue"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionallistvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalListValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalListValue:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionallistwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalListWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalListWasNull"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionallistwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalListWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalListWasNull:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionallistwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalListWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalListWasPresent"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionallistwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalListWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalListWasPresent:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableOptionalStringValue"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStringValue:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalStringWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalStringWasNull"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStringWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStringWasNull:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalStringWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalStringWasPresent"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstringwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStringWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStringWasPresent:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalStructValue() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableOptionalStructValue"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStructValue(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStructValue:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalStructWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalStructWasNull"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStructWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStructWasNull:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableOptionalStructWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalStructWasPresent"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullableoptionalstructwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStructWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStructWasPresent:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablestringvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableStringValue"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablestringvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStringValue:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablestringwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableStringWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableStringWasNull"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablestringwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableStringWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStringWasNull:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablestructvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableStructValue() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableStructValue"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablestructvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableStructValue(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStructValue:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablestructwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) NullableStructWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableStructWasNull"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/nullablestructwasnull
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetNullableStructWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStructWasNull:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalintvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) OptionalIntValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalIntValue"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalintvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetOptionalIntValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalIntValue:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalintwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) OptionalIntWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalIntWasPresent"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalintwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetOptionalIntWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalIntWasPresent:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionallistvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) OptionalListValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("optionalListValue"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionallistvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetOptionalListValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalListValue:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionallistwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) OptionalListWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalListWasPresent"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionallistwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetOptionalListWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalListWasPresent:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalstringvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) OptionalStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("optionalStringValue"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalstringvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetOptionalStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStringValue:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalstringwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) OptionalStringWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalStringWasPresent"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalstringwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetOptionalStringWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStringWasPresent:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalstructvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) OptionalStructValue() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("optionalStructValue"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalstructvalue
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetOptionalStructValue(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStructValue:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalstructwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) OptionalStructWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalStructWasPresent"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/optionalstructwaspresent
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetOptionalStructWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStructWasPresent:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestcomplexnullableoptionalresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestComplexNullableOptionalResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

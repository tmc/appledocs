// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestComplexNullableOptionalRequestParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestComplexNullableOptionalRequestParams */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestComplexNullableOptionalRequestParams */
// An interface definition for the [MTRUnitTestingClusterTestComplexNullableOptionalRequestParams] class.
type IMTRUnitTestingClusterTestComplexNullableOptionalRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestComplexNullableOptionalRequestParams */
	// properties:
	NullableInt() objc.IObject /* cross-framework: NSNumber */
	SetNullableInt(value objc.IObject /* cross-framework: NSNumber */)
	NullableList() objc.IObject /* cross-framework: NSArray */
	SetNullableList(value objc.IObject /* cross-framework: NSArray */)
	NullableOptionalInt() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalInt(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalList() objc.IObject /* cross-framework: NSArray */
	SetNullableOptionalList(value objc.IObject /* cross-framework: NSArray */)
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
	OptionalList() objc.IObject /* cross-framework: NSArray */
	SetOptionalList(value objc.IObject /* cross-framework: NSArray */)
	OptionalString() objc.IObject /* cross-framework: NSString */
	SetOptionalString(value objc.IObject /* cross-framework: NSString */)
	OptionalStruct() IMTRUnitTestingClusterSimpleStruct
	SetOptionalStruct(value IMTRUnitTestingClusterSimpleStruct)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestComplexNullableOptionalRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestComplexNullableOptionalRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsClass) Alloc() MTRUnitTestingClusterTestComplexNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestComplexNullableOptionalRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestComplexNullableOptionalRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams
type MTRUnitTestingClusterTestComplexNullableOptionalRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsFrom constructs a [MTRUnitTestingClusterTestComplexNullableOptionalRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestComplexNullableOptionalRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestComplexNullableOptionalRequestParams {
	return MTRUnitTestingClusterTestComplexNullableOptionalRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestComplexNullableOptionalRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestComplexNullableOptionalRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestComplexNullableOptionalRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestComplexNullableOptionalRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestComplexNullableOptionalRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/nullableInt
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableInt"))
	return rv
}/* debug [instance_properties/getter]: nullableInt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/nullableInt
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableInt:"), value)
}/* debug [instance_properties/setter]: nullableInt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/nullableList
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableList() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("nullableList"))
	return rv
}/* debug [instance_properties/getter]: nullableList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/nullableList
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableList(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableList:"), value)
}/* debug [instance_properties/setter]: nullableList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/nullableOptionalInt
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableOptionalInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalInt"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalInt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/nullableOptionalInt
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalInt:"), value)
}/* debug [instance_properties/setter]: nullableOptionalInt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/nullableOptionalList
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableOptionalList() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("nullableOptionalList"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/nullableOptionalList
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalList(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalList:"), value)
}/* debug [instance_properties/setter]: nullableOptionalList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/nullableOptionalString
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableOptionalString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableOptionalString"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/nullableOptionalString
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalString:"), value)
}/* debug [instance_properties/setter]: nullableOptionalString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/nullableOptionalStruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableOptionalStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableOptionalStruct"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/nullableOptionalStruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableOptionalStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStruct:"), value)
}/* debug [instance_properties/setter]: nullableOptionalStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/nullableString
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableString"))
	return rv
}/* debug [instance_properties/getter]: nullableString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/nullableString
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableString:"), value)
}/* debug [instance_properties/setter]: nullableString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/nullableStruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) NullableStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableStruct"))
	return rv
}/* debug [instance_properties/getter]: nullableStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/nullableStruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetNullableStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStruct:"), value)
}/* debug [instance_properties/setter]: nullableStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/optionalInt
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) OptionalInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalInt"))
	return rv
}/* debug [instance_properties/getter]: optionalInt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/optionalInt
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetOptionalInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalInt:"), value)
}/* debug [instance_properties/setter]: optionalInt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/optionalList
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) OptionalList() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("optionalList"))
	return rv
}/* debug [instance_properties/getter]: optionalList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/optionalList
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetOptionalList(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalList:"), value)
}/* debug [instance_properties/setter]: optionalList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/optionalString
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) OptionalString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("optionalString"))
	return rv
}/* debug [instance_properties/getter]: optionalString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/optionalString
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetOptionalString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalString:"), value)
}/* debug [instance_properties/setter]: optionalString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/optionalStruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) OptionalStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("optionalStruct"))
	return rv
}/* debug [instance_properties/getter]: optionalStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/optionalStruct
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetOptionalStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStruct:"), value)
}/* debug [instance_properties/setter]: optionalStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestComplexNullableOptionalRequestParams */




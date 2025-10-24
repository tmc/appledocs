// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestComplexNullableOptionalResponseParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestComplexNullableOptionalResponseParams */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestComplexNullableOptionalResponseParams */
// An interface definition for the [MTRUnitTestingClusterTestComplexNullableOptionalResponseParams] class.
type IMTRUnitTestingClusterTestComplexNullableOptionalResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestComplexNullableOptionalResponseParams */
	// properties:
	NullableIntValue() objc.IObject /* cross-framework: NSNumber */
	SetNullableIntValue(value objc.IObject /* cross-framework: NSNumber */)
	NullableIntWasNull() objc.IObject /* cross-framework: NSNumber */
	SetNullableIntWasNull(value objc.IObject /* cross-framework: NSNumber */)
	NullableListValue() objc.IObject /* cross-framework: NSArray */
	SetNullableListValue(value objc.IObject /* cross-framework: NSArray */)
	NullableListWasNull() objc.IObject /* cross-framework: NSNumber */
	SetNullableListWasNull(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalIntValue() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalIntValue(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalIntWasNull() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalIntWasNull(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalIntWasPresent() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalIntWasPresent(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalListValue() objc.IObject /* cross-framework: NSArray */
	SetNullableOptionalListValue(value objc.IObject /* cross-framework: NSArray */)
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
	OptionalListValue() objc.IObject /* cross-framework: NSArray */
	SetOptionalListValue(value objc.IObject /* cross-framework: NSArray */)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestComplexNullableOptionalResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestComplexNullableOptionalResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsClass) Alloc() MTRUnitTestingClusterTestComplexNullableOptionalResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestComplexNullableOptionalResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestComplexNullableOptionalResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams
type MTRUnitTestingClusterTestComplexNullableOptionalResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsFrom constructs a [MTRUnitTestingClusterTestComplexNullableOptionalResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestComplexNullableOptionalResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestComplexNullableOptionalResponseParams {
	return MTRUnitTestingClusterTestComplexNullableOptionalResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestComplexNullableOptionalResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/init(responseValue:)
func NewMTRUnitTestingClusterTestComplexNullableOptionalResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRUnitTestingClusterTestComplexNullableOptionalResponseParams {
	instance := getMTRUnitTestingClusterTestComplexNullableOptionalResponseParamsClass().Alloc()
	rv := objc.Send[MTRUnitTestingClusterTestComplexNullableOptionalResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRUnitTestingClusterTestComplexNullableOptionalResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestComplexNullableOptionalResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestComplexNullableOptionalResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestComplexNullableOptionalResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestComplexNullableOptionalResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableIntValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableIntValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableIntValue"))
	return rv
}/* debug [instance_properties/getter]: nullableIntValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableIntValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableIntValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableIntValue:"), value)
}/* debug [instance_properties/setter]: nullableIntValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableIntWasNull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableIntWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableIntWasNull"))
	return rv
}/* debug [instance_properties/getter]: nullableIntWasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableIntWasNull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableIntWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableIntWasNull:"), value)
}/* debug [instance_properties/setter]: nullableIntWasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableListValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableListValue() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("nullableListValue"))
	return rv
}/* debug [instance_properties/getter]: nullableListValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableListValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableListValue(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableListValue:"), value)
}/* debug [instance_properties/setter]: nullableListValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableListWasNull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableListWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableListWasNull"))
	return rv
}/* debug [instance_properties/getter]: nullableListWasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableListWasNull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableListWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableListWasNull:"), value)
}/* debug [instance_properties/setter]: nullableListWasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalIntValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalIntValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalIntValue"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalIntValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalIntValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalIntValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalIntValue:"), value)
}/* debug [instance_properties/setter]: nullableOptionalIntValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalIntWasNull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalIntWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalIntWasNull"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalIntWasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalIntWasNull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalIntWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalIntWasNull:"), value)
}/* debug [instance_properties/setter]: nullableOptionalIntWasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalIntWasPresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalIntWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalIntWasPresent"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalIntWasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalIntWasPresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalIntWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalIntWasPresent:"), value)
}/* debug [instance_properties/setter]: nullableOptionalIntWasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalListValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalListValue() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("nullableOptionalListValue"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalListValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalListValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalListValue(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalListValue:"), value)
}/* debug [instance_properties/setter]: nullableOptionalListValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalListWasNull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalListWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalListWasNull"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalListWasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalListWasNull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalListWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalListWasNull:"), value)
}/* debug [instance_properties/setter]: nullableOptionalListWasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalListWasPresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalListWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalListWasPresent"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalListWasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalListWasPresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalListWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalListWasPresent:"), value)
}/* debug [instance_properties/setter]: nullableOptionalListWasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalStringValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableOptionalStringValue"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalStringValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalStringValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStringValue:"), value)
}/* debug [instance_properties/setter]: nullableOptionalStringValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalStringWasNull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStringWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalStringWasNull"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalStringWasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalStringWasNull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStringWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStringWasNull:"), value)
}/* debug [instance_properties/setter]: nullableOptionalStringWasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalStringWasPresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStringWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalStringWasPresent"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalStringWasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalStringWasPresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStringWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStringWasPresent:"), value)
}/* debug [instance_properties/setter]: nullableOptionalStringWasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalStructValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStructValue() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableOptionalStructValue"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalStructValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalStructValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStructValue(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStructValue:"), value)
}/* debug [instance_properties/setter]: nullableOptionalStructValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalStructWasNull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStructWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalStructWasNull"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalStructWasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalStructWasNull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStructWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStructWasNull:"), value)
}/* debug [instance_properties/setter]: nullableOptionalStructWasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalStructWasPresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableOptionalStructWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalStructWasPresent"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalStructWasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableOptionalStructWasPresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableOptionalStructWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStructWasPresent:"), value)
}/* debug [instance_properties/setter]: nullableOptionalStructWasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableStringValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableStringValue"))
	return rv
}/* debug [instance_properties/getter]: nullableStringValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableStringValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStringValue:"), value)
}/* debug [instance_properties/setter]: nullableStringValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableStringWasNull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableStringWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableStringWasNull"))
	return rv
}/* debug [instance_properties/getter]: nullableStringWasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableStringWasNull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableStringWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStringWasNull:"), value)
}/* debug [instance_properties/setter]: nullableStringWasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableStructValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableStructValue() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableStructValue"))
	return rv
}/* debug [instance_properties/getter]: nullableStructValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableStructValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableStructValue(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStructValue:"), value)
}/* debug [instance_properties/setter]: nullableStructValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableStructWasNull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) NullableStructWasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableStructWasNull"))
	return rv
}/* debug [instance_properties/getter]: nullableStructWasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/nullableStructWasNull
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetNullableStructWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStructWasNull:"), value)
}/* debug [instance_properties/setter]: nullableStructWasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/optionalIntValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalIntValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalIntValue"))
	return rv
}/* debug [instance_properties/getter]: optionalIntValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/optionalIntValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalIntValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalIntValue:"), value)
}/* debug [instance_properties/setter]: optionalIntValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/optionalIntWasPresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalIntWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalIntWasPresent"))
	return rv
}/* debug [instance_properties/getter]: optionalIntWasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/optionalIntWasPresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalIntWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalIntWasPresent:"), value)
}/* debug [instance_properties/setter]: optionalIntWasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/optionalListValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalListValue() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("optionalListValue"))
	return rv
}/* debug [instance_properties/getter]: optionalListValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/optionalListValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalListValue(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalListValue:"), value)
}/* debug [instance_properties/setter]: optionalListValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/optionalListWasPresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalListWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalListWasPresent"))
	return rv
}/* debug [instance_properties/getter]: optionalListWasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/optionalListWasPresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalListWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalListWasPresent:"), value)
}/* debug [instance_properties/setter]: optionalListWasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/optionalStringValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("optionalStringValue"))
	return rv
}/* debug [instance_properties/getter]: optionalStringValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/optionalStringValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStringValue:"), value)
}/* debug [instance_properties/setter]: optionalStringValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/optionalStringWasPresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalStringWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalStringWasPresent"))
	return rv
}/* debug [instance_properties/getter]: optionalStringWasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/optionalStringWasPresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalStringWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStringWasPresent:"), value)
}/* debug [instance_properties/setter]: optionalStringWasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/optionalStructValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalStructValue() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("optionalStructValue"))
	return rv
}/* debug [instance_properties/getter]: optionalStructValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/optionalStructValue
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalStructValue(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStructValue:"), value)
}/* debug [instance_properties/setter]: optionalStructValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/optionalStructWasPresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) OptionalStructWasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalStructWasPresent"))
	return rv
}/* debug [instance_properties/getter]: optionalStructWasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/optionalStructWasPresent
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetOptionalStructWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStructWasPresent:"), value)
}/* debug [instance_properties/setter]: optionalStructWasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestComplexNullableOptionalResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestComplexNullableOptionalResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestComplexNullableOptionalResponseParams */



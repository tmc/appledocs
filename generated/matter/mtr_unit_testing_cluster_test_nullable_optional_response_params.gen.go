// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestNullableOptionalResponseParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestNullableOptionalResponseParams */
// The class instance for the [MTRUnitTestingClusterTestNullableOptionalResponseParams] class.
var (
	MTRUnitTestingClusterTestNullableOptionalResponseParamsClass     _MTRUnitTestingClusterTestNullableOptionalResponseParamsClass
	MTRUnitTestingClusterTestNullableOptionalResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestNullableOptionalResponseParamsClass() _MTRUnitTestingClusterTestNullableOptionalResponseParamsClass {
	MTRUnitTestingClusterTestNullableOptionalResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestNullableOptionalResponseParamsClass = _MTRUnitTestingClusterTestNullableOptionalResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestNullableOptionalResponseParams")}
	})
	return MTRUnitTestingClusterTestNullableOptionalResponseParamsClass
}

type _MTRUnitTestingClusterTestNullableOptionalResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestNullableOptionalResponseParams */
// An interface definition for the [MTRUnitTestingClusterTestNullableOptionalResponseParams] class.
type IMTRUnitTestingClusterTestNullableOptionalResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestNullableOptionalResponseParams */
	// properties:
	OriginalValue() objc.IObject /* cross-framework: NSNumber */
	SetOriginalValue(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
	WasNull() objc.IObject /* cross-framework: NSNumber */
	SetWasNull(value objc.IObject /* cross-framework: NSNumber */)
	WasPresent() objc.IObject /* cross-framework: NSNumber */
	SetWasPresent(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestNullableOptionalResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestNullableOptionalResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestNullableOptionalResponseParamsClass) Alloc() MTRUnitTestingClusterTestNullableOptionalResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestNullableOptionalResponseParamsClass) New() MTRUnitTestingClusterTestNullableOptionalResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) Init() MTRUnitTestingClusterTestNullableOptionalResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) Autorelease() MTRUnitTestingClusterTestNullableOptionalResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestNullableOptionalResponseParams creates a new MTRUnitTestingClusterTestNullableOptionalResponseParams instance.
func NewMTRUnitTestingClusterTestNullableOptionalResponseParams() MTRUnitTestingClusterTestNullableOptionalResponseParams {
	return getMTRUnitTestingClusterTestNullableOptionalResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestNullableOptionalResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalResponseParams
type MTRUnitTestingClusterTestNullableOptionalResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestNullableOptionalResponseParamsFrom constructs a [MTRUnitTestingClusterTestNullableOptionalResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestNullableOptionalResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestNullableOptionalResponseParams {
	return MTRUnitTestingClusterTestNullableOptionalResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestNullableOptionalResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalResponseParams/init(responseValue:)
func NewMTRUnitTestingClusterTestNullableOptionalResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRUnitTestingClusterTestNullableOptionalResponseParams {
	instance := getMTRUnitTestingClusterTestNullableOptionalResponseParamsClass().Alloc()
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRUnitTestingClusterTestNullableOptionalResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestNullableOptionalResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestNullableOptionalResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestNullableOptionalResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestNullableOptionalResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalResponseParams/originalValue
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) OriginalValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("originalValue"))
	return rv
}/* debug [instance_properties/getter]: originalValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalResponseParams/originalValue
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) SetOriginalValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOriginalValue:"), value)
}/* debug [instance_properties/setter]: originalValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalResponseParams/value
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalResponseParams/value
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalResponseParams/wasNull
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) WasNull() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("wasNull"))
	return rv
}/* debug [instance_properties/getter]: wasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalResponseParams/wasNull
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) SetWasNull(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWasNull:"), value)
}/* debug [instance_properties/setter]: wasNull */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalResponseParams/wasPresent
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) WasPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("wasPresent"))
	return rv
}/* debug [instance_properties/getter]: wasPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalResponseParams/wasPresent
func (m_ MTRUnitTestingClusterTestNullableOptionalResponseParams) SetWasPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWasPresent:"), value)
}/* debug [instance_properties/setter]: wasPresent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestNullableOptionalResponseParams */



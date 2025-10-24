// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterBooleanResponseParams */


/* debug [class_header]: Header for MTRUnitTestingClusterBooleanResponseParams */
// The class instance for the [MTRUnitTestingClusterBooleanResponseParams] class.
var (
	MTRUnitTestingClusterBooleanResponseParamsClass     _MTRUnitTestingClusterBooleanResponseParamsClass
	MTRUnitTestingClusterBooleanResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterBooleanResponseParamsClass() _MTRUnitTestingClusterBooleanResponseParamsClass {
	MTRUnitTestingClusterBooleanResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterBooleanResponseParamsClass = _MTRUnitTestingClusterBooleanResponseParamsClass{objc.GetClass("MTRUnitTestingClusterBooleanResponseParams")}
	})
	return MTRUnitTestingClusterBooleanResponseParamsClass
}

type _MTRUnitTestingClusterBooleanResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterBooleanResponseParams */
// An interface definition for the [MTRUnitTestingClusterBooleanResponseParams] class.
type IMTRUnitTestingClusterBooleanResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterBooleanResponseParams */
	// properties:
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterBooleanResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterBooleanResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterBooleanResponseParamsClass) Alloc() MTRUnitTestingClusterBooleanResponseParams {
	rv := objc.Send[MTRUnitTestingClusterBooleanResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterBooleanResponseParamsClass) New() MTRUnitTestingClusterBooleanResponseParams {
	rv := objc.Send[MTRUnitTestingClusterBooleanResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterBooleanResponseParams) Init() MTRUnitTestingClusterBooleanResponseParams {
	rv := objc.Send[MTRUnitTestingClusterBooleanResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterBooleanResponseParams) Autorelease() MTRUnitTestingClusterBooleanResponseParams {
	rv := objc.Send[MTRUnitTestingClusterBooleanResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterBooleanResponseParams creates a new MTRUnitTestingClusterBooleanResponseParams instance.
func NewMTRUnitTestingClusterBooleanResponseParams() MTRUnitTestingClusterBooleanResponseParams {
	return getMTRUnitTestingClusterBooleanResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterBooleanResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterBooleanResponseParams
type MTRUnitTestingClusterBooleanResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterBooleanResponseParamsFrom constructs a [MTRUnitTestingClusterBooleanResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterBooleanResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterBooleanResponseParams {
	return MTRUnitTestingClusterBooleanResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterBooleanResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterBooleanResponseParams/init(responseValue:)
func NewMTRUnitTestingClusterBooleanResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRUnitTestingClusterBooleanResponseParams {
	instance := getMTRUnitTestingClusterBooleanResponseParamsClass().Alloc()
	rv := objc.Send[MTRUnitTestingClusterBooleanResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRUnitTestingClusterBooleanResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterBooleanResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterBooleanResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterBooleanResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterBooleanResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterBooleanResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterBooleanResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterBooleanResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterBooleanResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterBooleanResponseParams/value
func (m_ MTRUnitTestingClusterBooleanResponseParams) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterBooleanResponseParams/value
func (m_ MTRUnitTestingClusterBooleanResponseParams) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterBooleanResponseParams */



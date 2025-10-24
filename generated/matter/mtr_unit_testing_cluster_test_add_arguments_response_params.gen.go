// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestAddArgumentsResponseParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestAddArgumentsResponseParams */
// The class instance for the [MTRUnitTestingClusterTestAddArgumentsResponseParams] class.
var (
	MTRUnitTestingClusterTestAddArgumentsResponseParamsClass     _MTRUnitTestingClusterTestAddArgumentsResponseParamsClass
	MTRUnitTestingClusterTestAddArgumentsResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestAddArgumentsResponseParamsClass() _MTRUnitTestingClusterTestAddArgumentsResponseParamsClass {
	MTRUnitTestingClusterTestAddArgumentsResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestAddArgumentsResponseParamsClass = _MTRUnitTestingClusterTestAddArgumentsResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestAddArgumentsResponseParams")}
	})
	return MTRUnitTestingClusterTestAddArgumentsResponseParamsClass
}

type _MTRUnitTestingClusterTestAddArgumentsResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestAddArgumentsResponseParams */
// An interface definition for the [MTRUnitTestingClusterTestAddArgumentsResponseParams] class.
type IMTRUnitTestingClusterTestAddArgumentsResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestAddArgumentsResponseParams */
	// properties:
	ReturnValue() objc.IObject /* cross-framework: NSNumber */
	SetReturnValue(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestAddArgumentsResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestAddArgumentsResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestAddArgumentsResponseParamsClass) Alloc() MTRUnitTestingClusterTestAddArgumentsResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestAddArgumentsResponseParamsClass) New() MTRUnitTestingClusterTestAddArgumentsResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestAddArgumentsResponseParams) Init() MTRUnitTestingClusterTestAddArgumentsResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestAddArgumentsResponseParams) Autorelease() MTRUnitTestingClusterTestAddArgumentsResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestAddArgumentsResponseParams creates a new MTRUnitTestingClusterTestAddArgumentsResponseParams instance.
func NewMTRUnitTestingClusterTestAddArgumentsResponseParams() MTRUnitTestingClusterTestAddArgumentsResponseParams {
	return getMTRUnitTestingClusterTestAddArgumentsResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestAddArgumentsResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsResponseParams
type MTRUnitTestingClusterTestAddArgumentsResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestAddArgumentsResponseParamsFrom constructs a [MTRUnitTestingClusterTestAddArgumentsResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestAddArgumentsResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestAddArgumentsResponseParams {
	return MTRUnitTestingClusterTestAddArgumentsResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestAddArgumentsResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsResponseParams/init(responseValue:)
func NewMTRUnitTestingClusterTestAddArgumentsResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRUnitTestingClusterTestAddArgumentsResponseParams {
	instance := getMTRUnitTestingClusterTestAddArgumentsResponseParamsClass().Alloc()
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRUnitTestingClusterTestAddArgumentsResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestAddArgumentsResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestAddArgumentsResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestAddArgumentsResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestAddArgumentsResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsResponseParams/returnValue
func (m_ MTRUnitTestingClusterTestAddArgumentsResponseParams) ReturnValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("returnValue"))
	return rv
}/* debug [instance_properties/getter]: returnValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsResponseParams/returnValue
func (m_ MTRUnitTestingClusterTestAddArgumentsResponseParams) SetReturnValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReturnValue:"), value)
}/* debug [instance_properties/setter]: returnValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestAddArgumentsResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestAddArgumentsResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestAddArgumentsResponseParams */



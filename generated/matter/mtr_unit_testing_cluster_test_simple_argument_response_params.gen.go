// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestSimpleArgumentResponseParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestSimpleArgumentResponseParams */
// The class instance for the [MTRUnitTestingClusterTestSimpleArgumentResponseParams] class.
var (
	MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass     _MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass
	MTRUnitTestingClusterTestSimpleArgumentResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestSimpleArgumentResponseParamsClass() _MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass {
	MTRUnitTestingClusterTestSimpleArgumentResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass = _MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestSimpleArgumentResponseParams")}
	})
	return MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass
}

type _MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestSimpleArgumentResponseParams */
// An interface definition for the [MTRUnitTestingClusterTestSimpleArgumentResponseParams] class.
type IMTRUnitTestingClusterTestSimpleArgumentResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestSimpleArgumentResponseParams */
	// properties:
	ReturnValue() objc.IObject /* cross-framework: NSNumber */
	SetReturnValue(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestSimpleArgumentResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestSimpleArgumentResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass) Alloc() MTRUnitTestingClusterTestSimpleArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass) New() MTRUnitTestingClusterTestSimpleArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestSimpleArgumentResponseParams) Init() MTRUnitTestingClusterTestSimpleArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestSimpleArgumentResponseParams) Autorelease() MTRUnitTestingClusterTestSimpleArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestSimpleArgumentResponseParams creates a new MTRUnitTestingClusterTestSimpleArgumentResponseParams instance.
func NewMTRUnitTestingClusterTestSimpleArgumentResponseParams() MTRUnitTestingClusterTestSimpleArgumentResponseParams {
	return getMTRUnitTestingClusterTestSimpleArgumentResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestSimpleArgumentResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleArgumentResponseParams
type MTRUnitTestingClusterTestSimpleArgumentResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestSimpleArgumentResponseParamsFrom constructs a [MTRUnitTestingClusterTestSimpleArgumentResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestSimpleArgumentResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestSimpleArgumentResponseParams {
	return MTRUnitTestingClusterTestSimpleArgumentResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestSimpleArgumentResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleArgumentResponseParams/init(responseValue:)
func NewMTRUnitTestingClusterTestSimpleArgumentResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRUnitTestingClusterTestSimpleArgumentResponseParams {
	instance := getMTRUnitTestingClusterTestSimpleArgumentResponseParamsClass().Alloc()
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRUnitTestingClusterTestSimpleArgumentResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestSimpleArgumentResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestSimpleArgumentResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestSimpleArgumentResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestSimpleArgumentResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleArgumentResponseParams/returnValue
func (m_ MTRUnitTestingClusterTestSimpleArgumentResponseParams) ReturnValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("returnValue"))
	return rv
}/* debug [instance_properties/getter]: returnValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleArgumentResponseParams/returnValue
func (m_ MTRUnitTestingClusterTestSimpleArgumentResponseParams) SetReturnValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReturnValue:"), value)
}/* debug [instance_properties/setter]: returnValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleArgumentResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestSimpleArgumentResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleArgumentResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestSimpleArgumentResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestSimpleArgumentResponseParams */



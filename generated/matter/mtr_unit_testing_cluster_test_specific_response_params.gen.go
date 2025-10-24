// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestSpecificResponseParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestSpecificResponseParams */
// The class instance for the [MTRUnitTestingClusterTestSpecificResponseParams] class.
var (
	MTRUnitTestingClusterTestSpecificResponseParamsClass     _MTRUnitTestingClusterTestSpecificResponseParamsClass
	MTRUnitTestingClusterTestSpecificResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestSpecificResponseParamsClass() _MTRUnitTestingClusterTestSpecificResponseParamsClass {
	MTRUnitTestingClusterTestSpecificResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestSpecificResponseParamsClass = _MTRUnitTestingClusterTestSpecificResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestSpecificResponseParams")}
	})
	return MTRUnitTestingClusterTestSpecificResponseParamsClass
}

type _MTRUnitTestingClusterTestSpecificResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestSpecificResponseParams */
// An interface definition for the [MTRUnitTestingClusterTestSpecificResponseParams] class.
type IMTRUnitTestingClusterTestSpecificResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestSpecificResponseParams */
	// properties:
	ReturnValue() objc.IObject /* cross-framework: NSNumber */
	SetReturnValue(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestSpecificResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestSpecificResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestSpecificResponseParamsClass) Alloc() MTRUnitTestingClusterTestSpecificResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestSpecificResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestSpecificResponseParamsClass) New() MTRUnitTestingClusterTestSpecificResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestSpecificResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestSpecificResponseParams) Init() MTRUnitTestingClusterTestSpecificResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestSpecificResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestSpecificResponseParams) Autorelease() MTRUnitTestingClusterTestSpecificResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestSpecificResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestSpecificResponseParams creates a new MTRUnitTestingClusterTestSpecificResponseParams instance.
func NewMTRUnitTestingClusterTestSpecificResponseParams() MTRUnitTestingClusterTestSpecificResponseParams {
	return getMTRUnitTestingClusterTestSpecificResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestSpecificResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSpecificResponseParams
type MTRUnitTestingClusterTestSpecificResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestSpecificResponseParamsFrom constructs a [MTRUnitTestingClusterTestSpecificResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestSpecificResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestSpecificResponseParams {
	return MTRUnitTestingClusterTestSpecificResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestSpecificResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSpecificResponseParams/init(responseValue:)
func NewMTRUnitTestingClusterTestSpecificResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRUnitTestingClusterTestSpecificResponseParams {
	instance := getMTRUnitTestingClusterTestSpecificResponseParamsClass().Alloc()
	rv := objc.Send[MTRUnitTestingClusterTestSpecificResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRUnitTestingClusterTestSpecificResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestSpecificResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestSpecificResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestSpecificResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestSpecificResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSpecificResponseParams/returnValue
func (m_ MTRUnitTestingClusterTestSpecificResponseParams) ReturnValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("returnValue"))
	return rv
}/* debug [instance_properties/getter]: returnValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSpecificResponseParams/returnValue
func (m_ MTRUnitTestingClusterTestSpecificResponseParams) SetReturnValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReturnValue:"), value)
}/* debug [instance_properties/setter]: returnValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSpecificResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestSpecificResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSpecificResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestSpecificResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestSpecificResponseParams */



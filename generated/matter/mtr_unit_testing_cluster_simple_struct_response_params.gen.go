// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterSimpleStructResponseParams */


/* debug [class_header]: Header for MTRUnitTestingClusterSimpleStructResponseParams */
// The class instance for the [MTRUnitTestingClusterSimpleStructResponseParams] class.
var (
	MTRUnitTestingClusterSimpleStructResponseParamsClass     _MTRUnitTestingClusterSimpleStructResponseParamsClass
	MTRUnitTestingClusterSimpleStructResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterSimpleStructResponseParamsClass() _MTRUnitTestingClusterSimpleStructResponseParamsClass {
	MTRUnitTestingClusterSimpleStructResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterSimpleStructResponseParamsClass = _MTRUnitTestingClusterSimpleStructResponseParamsClass{objc.GetClass("MTRUnitTestingClusterSimpleStructResponseParams")}
	})
	return MTRUnitTestingClusterSimpleStructResponseParamsClass
}

type _MTRUnitTestingClusterSimpleStructResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterSimpleStructResponseParams */
// An interface definition for the [MTRUnitTestingClusterSimpleStructResponseParams] class.
type IMTRUnitTestingClusterSimpleStructResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterSimpleStructResponseParams */
	// properties:
	Arg1() IMTRUnitTestingClusterSimpleStruct
	SetArg1(value IMTRUnitTestingClusterSimpleStruct)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterSimpleStructResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterSimpleStructResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterSimpleStructResponseParamsClass) Alloc() MTRUnitTestingClusterSimpleStructResponseParams {
	rv := objc.Send[MTRUnitTestingClusterSimpleStructResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterSimpleStructResponseParamsClass) New() MTRUnitTestingClusterSimpleStructResponseParams {
	rv := objc.Send[MTRUnitTestingClusterSimpleStructResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterSimpleStructResponseParams) Init() MTRUnitTestingClusterSimpleStructResponseParams {
	rv := objc.Send[MTRUnitTestingClusterSimpleStructResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterSimpleStructResponseParams) Autorelease() MTRUnitTestingClusterSimpleStructResponseParams {
	rv := objc.Send[MTRUnitTestingClusterSimpleStructResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterSimpleStructResponseParams creates a new MTRUnitTestingClusterSimpleStructResponseParams instance.
func NewMTRUnitTestingClusterSimpleStructResponseParams() MTRUnitTestingClusterSimpleStructResponseParams {
	return getMTRUnitTestingClusterSimpleStructResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterSimpleStructResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStructResponseParams
type MTRUnitTestingClusterSimpleStructResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterSimpleStructResponseParamsFrom constructs a [MTRUnitTestingClusterSimpleStructResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterSimpleStructResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterSimpleStructResponseParams {
	return MTRUnitTestingClusterSimpleStructResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterSimpleStructResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStructResponseParams/init(responseValue:)
func NewMTRUnitTestingClusterSimpleStructResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRUnitTestingClusterSimpleStructResponseParams {
	instance := getMTRUnitTestingClusterSimpleStructResponseParamsClass().Alloc()
	rv := objc.Send[MTRUnitTestingClusterSimpleStructResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRUnitTestingClusterSimpleStructResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterSimpleStructResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterSimpleStructResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterSimpleStructResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterSimpleStructResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStructResponseParams/arg1
func (m_ MTRUnitTestingClusterSimpleStructResponseParams) Arg1() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("arg1"))
	return rv
}/* debug [instance_properties/getter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStructResponseParams/arg1
func (m_ MTRUnitTestingClusterSimpleStructResponseParams) SetArg1(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}/* debug [instance_properties/setter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStructResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterSimpleStructResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStructResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterSimpleStructResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterSimpleStructResponseParams */



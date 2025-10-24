// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestListInt8UReverseResponseParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestListInt8UReverseResponseParams */
// The class instance for the [MTRUnitTestingClusterTestListInt8UReverseResponseParams] class.
var (
	MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass     _MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass
	MTRUnitTestingClusterTestListInt8UReverseResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestListInt8UReverseResponseParamsClass() _MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass {
	MTRUnitTestingClusterTestListInt8UReverseResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass = _MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestListInt8UReverseResponseParams")}
	})
	return MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass
}

type _MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestListInt8UReverseResponseParams */
// An interface definition for the [MTRUnitTestingClusterTestListInt8UReverseResponseParams] class.
type IMTRUnitTestingClusterTestListInt8UReverseResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestListInt8UReverseResponseParams */
	// properties:
	Arg1() objc.IObject /* cross-framework: NSArray */
	SetArg1(value objc.IObject /* cross-framework: NSArray */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestListInt8UReverseResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestListInt8UReverseResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass) Alloc() MTRUnitTestingClusterTestListInt8UReverseResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass) New() MTRUnitTestingClusterTestListInt8UReverseResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestListInt8UReverseResponseParams) Init() MTRUnitTestingClusterTestListInt8UReverseResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestListInt8UReverseResponseParams) Autorelease() MTRUnitTestingClusterTestListInt8UReverseResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestListInt8UReverseResponseParams creates a new MTRUnitTestingClusterTestListInt8UReverseResponseParams instance.
func NewMTRUnitTestingClusterTestListInt8UReverseResponseParams() MTRUnitTestingClusterTestListInt8UReverseResponseParams {
	return getMTRUnitTestingClusterTestListInt8UReverseResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestListInt8UReverseResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UReverseResponseParams
type MTRUnitTestingClusterTestListInt8UReverseResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestListInt8UReverseResponseParamsFrom constructs a [MTRUnitTestingClusterTestListInt8UReverseResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestListInt8UReverseResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestListInt8UReverseResponseParams {
	return MTRUnitTestingClusterTestListInt8UReverseResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestListInt8UReverseResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UReverseResponseParams/init(responseValue:)
func NewMTRUnitTestingClusterTestListInt8UReverseResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRUnitTestingClusterTestListInt8UReverseResponseParams {
	instance := getMTRUnitTestingClusterTestListInt8UReverseResponseParamsClass().Alloc()
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRUnitTestingClusterTestListInt8UReverseResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestListInt8UReverseResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestListInt8UReverseResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestListInt8UReverseResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestListInt8UReverseResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UReverseResponseParams/arg1
func (m_ MTRUnitTestingClusterTestListInt8UReverseResponseParams) Arg1() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("arg1"))
	return rv
}/* debug [instance_properties/getter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UReverseResponseParams/arg1
func (m_ MTRUnitTestingClusterTestListInt8UReverseResponseParams) SetArg1(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}/* debug [instance_properties/setter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UReverseResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestListInt8UReverseResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UReverseResponseParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestListInt8UReverseResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestListInt8UReverseResponseParams */



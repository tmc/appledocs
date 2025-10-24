// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams */
// The class instance for the [MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass     _MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass
	MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass() _MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass {
	MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass = _MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams */
// An interface definition for the [MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams] class.
type IMTRUnitTestingClusterTestListNestedStructListArgumentRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams */
	// properties:
	Arg1() objc.IObject /* cross-framework: NSArray */
	SetArg1(value objc.IObject /* cross-framework: NSArray */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass) New() MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams) Init() MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestListNestedStructListArgumentRequestParams creates a new MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestListNestedStructListArgumentRequestParams() MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams {
	return getMTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams
type MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams {
	return MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams) Arg1() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("arg1"))
	return rv
}/* debug [instance_properties/getter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams) SetArg1(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}/* debug [instance_properties/setter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams */




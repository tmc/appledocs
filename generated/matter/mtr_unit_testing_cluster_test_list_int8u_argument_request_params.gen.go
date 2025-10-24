// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestListInt8UArgumentRequestParams */


/* debug [class_header]: Header for MTRUnitTestingClusterTestListInt8UArgumentRequestParams */
// The class instance for the [MTRUnitTestingClusterTestListInt8UArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass     _MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass
	MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass() _MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass {
	MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass = _MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestListInt8UArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestListInt8UArgumentRequestParams */
// An interface definition for the [MTRUnitTestingClusterTestListInt8UArgumentRequestParams] class.
type IMTRUnitTestingClusterTestListInt8UArgumentRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestListInt8UArgumentRequestParams */
	// properties:
	Arg1() objc.IObject /* cross-framework: NSArray */
	SetArg1(value objc.IObject /* cross-framework: NSArray */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestListInt8UArgumentRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestListInt8UArgumentRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestListInt8UArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass) New() MTRUnitTestingClusterTestListInt8UArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestListInt8UArgumentRequestParams) Init() MTRUnitTestingClusterTestListInt8UArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestListInt8UArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestListInt8UArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestListInt8UArgumentRequestParams creates a new MTRUnitTestingClusterTestListInt8UArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestListInt8UArgumentRequestParams() MTRUnitTestingClusterTestListInt8UArgumentRequestParams {
	return getMTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestListInt8UArgumentRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UArgumentRequestParams
type MTRUnitTestingClusterTestListInt8UArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestListInt8UArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestListInt8UArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestListInt8UArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestListInt8UArgumentRequestParams {
	return MTRUnitTestingClusterTestListInt8UArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestListInt8UArgumentRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestListInt8UArgumentRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestListInt8UArgumentRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestListInt8UArgumentRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestListInt8UArgumentRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestListInt8UArgumentRequestParams) Arg1() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("arg1"))
	return rv
}/* debug [instance_properties/getter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UArgumentRequestParams/arg1
func (m_ MTRUnitTestingClusterTestListInt8UArgumentRequestParams) SetArg1(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}/* debug [instance_properties/setter]: arg1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestListInt8UArgumentRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UArgumentRequestParams/serverSideProcessingTimeout
func (m_ MTRUnitTestingClusterTestListInt8UArgumentRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestListInt8UArgumentRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UArgumentRequestParams/timedInvokeTimeoutMs
func (m_ MTRUnitTestingClusterTestListInt8UArgumentRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestListInt8UArgumentRequestParams */



